package comm

import(
    "sync"
    "encoding/json"
    "time"
    "google.golang.org/grpc"
    "context"
    "errors"
    log "github.com/labstack/gommon/log"
    pb "kvstore/proto/commpb"
    "kvstore/logging"
    "kvstore/memdb"
    "path/filepath"
)

//ROLE role of a node
type ROLE int

const(
    // Follower follower role
    Follower ROLE = iota
    // Candidate candidate role
    Candidate
    // Leader leader role
    Leader
)

const(
    NORMAL int64 = iota
    EMPTY
    CFGCHA
)

const(
    // HBLOWER heartbeat timeout lower bound
    HBLOWER = 200
    // HBUPPER heartbeat timeout upper bound
    HBUPPER = 400
)

// Server Core data structure for a node
type Server struct{
    id int64
    term int64
    voted bool
    termLock sync.RWMutex
    role ROLE
    addr string
    leaderID int64
    hbtimeout int64
    lastack bool
    nodes map[int64]*Node //id:string
    nodesdb *memdb.MemStore
    nodeslog *logging.LogStore
    nodelock sync.Mutex
    //nodelock sync.RWMutex
    errorC chan error
    dblog *logging.LogStore
    applylock sync.Mutex
    curcommit *pb.Commit
    grpcdisabled bool
}

// Node represent a node
type Node struct{
    id int64
    addr string
    conn *grpc.ClientConn
}

//NewServer return a server
func NewServer(nid int64, addrs []string, db logging.DB, datadir string)*Server{
    if nid > int64(len(addrs)){
        log.Fatalf("id should be smaller than cluster addresses")
    }
    dblog,err := logging.NewDBLogStore(db, filepath.Join(datadir,"db"))
    if err != nil{
        log.Fatalf("NewServer, create NewDBLogStore failed:%v",err)
    }
    s := &Server{
        grpcdisabled:false,
        lastack:true,
        id:nid,
        term:0,
        voted:false,
        addr:addrs[nid],
        role:Follower,
        nodes:make(map[int64]*Node),
        errorC: make(chan error,1),
        dblog:dblog,
        leaderID:-1,
    }

    s.nodesdb = memdb.New()
    s.nodeslog,err = logging.NewDBLogStore(s.nodesdb, filepath.Join(datadir,"member"))
    s.initmembers(addrs)
    s.rebuildConns()
    s.serve(s.addr)
    //s.initConn()
    s.startHBCheck()
    go s.sendHB()
    return s
}


func (s *Server)initmembers(addrs []string){
    _,lastid := s.nodeslog.GetLastCommit()
    for i,ad := range addrs{
        lastid ++
        //s.nodes[int64(i)] = &Node{id:int64(i),addr:ad}
        ope := memdb.Operation{
            Type:memdb.ADD,
            Key:int64(i),
            Value:ad,
        }
        opeb,_ := json.Marshal(ope)
        err := s.nodeslog.Apply(s.term, lastid,opeb)
        if err != nil{
            panic(err)
        }
    }
}
func (s *Server)rebuildConns(){
    s.nodelock.Lock()
    defer s.nodelock.Unlock()
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())
    //TODO: close old connections
    s.nodes = make(map[int64]*Node)
    s.nodesdb.Range(func(key int64,value string)bool{
        conn,err := grpc.Dial(value,opts...)
        if err != nil{
            log.Errorf("grpc dial failed:%v",err)
        }
        s.nodes[key] = &Node{
            id:key,
            addr: value,
            conn:conn,
        }
        return true
    })
}

/*func (s *Server)initConn(){
    s.nodelock.RLock()
    defer s.nodelock.RUnlock()
    for _,nd := range s.nodes{
    }
}*/


func (s *Server) changeTerm(newterm int64){
    s.termLock.Lock()
    defer s.termLock.Unlock()
    if newterm > s.term{
        s.term = newterm
        s.voted = false
        if s.role == Leader{
            s.role = Follower
            log.Infof("another node got bigger term,node %v stepdown from leader to follower",s.id)
        }
    }
}

func (s *Server)getTerm()int64{
    s.termLock.RLock()
    defer s.termLock.RUnlock()
    return s.term
}

func (s *Server)vote(req *pb.VoteReq)(res bool,oldterm int64){
    s.termLock.Lock()
    defer s.termLock.Unlock()
    lastterm,lastid := s.dblog.GetLastCommit()
    oldterm = s.term
    if req.Term > oldterm{
        s.term = req.Term
        s.voted = false
        s.role = Follower
        if (req.LastLogTerm > lastterm) || (req.LastLogTerm == lastterm && req.LastLogId >= lastid){
            res = true
            s.voted = true
            //log.Infof("node %v got bigger term, stepdown from leader to follower",s.id)
        }else{
            res = false
        }
    }else{
        if !s.voted{
            if (req.LastLogTerm > lastterm) || (req.LastLogTerm == lastterm && req.LastLogId >= lastid){
                res = true
                s.voted = true
                s.role = Follower
                log.Infof("node %v got bigger term, stepdown from leader to follower",s.id)
            }else{
                res = false
            }
        }else{
            res = false
        }
    }
    return
}

func (s *Server)incTerm(){
    s.termLock.Lock()
    defer s.termLock.Unlock()
    s.term ++
    s.voted = true
}

func (s *Server)getCheckPoint(tp int64)([]byte, error){
    if tp == NORMAL{
        return s.dblog.Marshal()
    }else{
        return s.nodeslog.Marshal()
    }
}


//ProposeEmpty propose a empty commit to sync data with leader node when a node restarts
// this functon should be invoked whenever these is a leader change
func (s *Server)ProposeEmpty()(err error){
    if s.role != Leader{
        //log.Infof("node %v is not the leader, pass the commit the leader node %v",s.id,s.leaderID)
        //not leader, need to propose to leader
        commit := &pb.Commit{
            Type:EMPTY,
        }
        s.nodelock.Lock()
        client := pb.NewCommpbClient(s.nodes[s.leaderID].conn)
        s.nodelock.Unlock()
        ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
        defer cancel()
        _,err = client.SendToLeader(ctx,commit)
        if err != nil{
            return err
        }
    }
    return nil
}




//Propose propose a commit to the cluster
//TODO: propose node add/del
func (s *Server)Propose(data []byte, ctype int64)(err error){
    if s.role != Leader{
        if ctype != EMPTY{
            log.Infof("node %v is not the leader, pass the commit the leader node %v",s.id,s.leaderID)
        }
        //not leader, need to propose to leader
        commit := &pb.Commit{
            Data:data,
            Type:ctype,
        }
        s.nodelock.Lock()
        client := pb.NewCommpbClient(s.nodes[s.leaderID].conn)
        s.nodelock.Unlock()
        ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
        defer cancel()
        _,err = client.SendToLeader(ctx,commit)
        if err != nil{
            return err
        }
    }else{
        //leader
        if ctype != EMPTY{
            log.Infof("node %v is the leader, process the commit",s.id)
        }
        s.applylock.Lock()
        defer s.applylock.Unlock()
        //Get last log id and term from logging module
        lastterm, lastid := s.dblog.GetLastCommit()
        lastmemterm, lastmemid := s.nodeslog.GetLastCommit()
        id := lastid + 1
        if s.term != lastterm{
            id = 0
        }
        commit := &pb.Commit{
            Data:data,
            Term:s.term,
            Id:id,
            LastTerm:lastterm,
            LastId: lastid,
            SrcId: s.id,
            Type:ctype,
            LastMemTerm: lastmemterm,
            LastMemId: lastmemid,
        }
        //Send commit to all the nodes
        //prepare the commit
        successcnt := 0
        for _,nd := range s.nodes{
            if ctype != EMPTY{
                log.Infof("send the commit to node %v",nd.id)
            }
            client := pb.NewCommpbClient(nd.conn)
            ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
            _,err = client.Prepare(ctx,commit)
            if err != nil{
                if ctype != EMPTY{
                    log.Infof("node %v is not ready to commit",nd.id)
                }
            }else{
                if ctype != EMPTY{
                    log.Infof("node %v is ready to commit",nd.id)
                }
                successcnt ++
            }
            cancel()
        }

        success := false
        if successcnt > len(s.nodes)/2{
            success = true
            if ctype != EMPTY{
                log.Info("commit reaches majority nodes, will confirm the commit")
            }
        }else{
            if ctype != EMPTY{
                log.Info("commit did not reach majority nodes, will cancel the commit")
            }
        }
        for _,nd := range s.nodes{
            client := pb.NewCommpbClient(nd.conn)
            ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
            if success{
                //confirm   
                if ctype != EMPTY{
                    log.Infof("sending confirm message to node %v",nd.id)
                }
                _,err = client.Confirm(ctx,commit)
            }else{
                //cancel
                if ctype != EMPTY{
                    log.Infof("sending cancel message to node %v",nd.id)
                }
                _,err = client.Cancel(ctx,commit)
            }
            cancel()
        }
        // nodes should compare its last term and log id before commit
        if success{
            return nil
        }else{
            return errors.New("fail to commit the request")
        }
    }
    return nil
}

func (s *Server)prepare(commit *pb.Commit)(error){
    lastterm, lastid := s.dblog.GetLastCommit()
    if lastterm != commit.LastTerm || lastid != commit.LastId{
        // get checkpoint from leader
        req := &pb.Msg{
            Type:NORMAL,
        }

        client := pb.NewCommpbClient(s.nodes[commit.SrcId].conn)
        ctx,cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
        defer cancel()
        rsp,err := client.GetCheckPoint(ctx,req)
        if err != nil{
            return err
        }
        if rsp.Status != 0{
            log.Errorf("get checkpoint from master failed")
            return errors.New("get checkpoint failed")
        }
        // recover from checkpoint
        err = s.dblog.Unmarshal(rsp.Data)
        if err != nil{
            log.Errorf("prepare recover Unmarshal failed:%v",err)
            return err
        }
        log.Info("data recovered from leader node!!!")
    }

    lastmemterm, lastmemid := s.nodeslog.GetLastCommit()
    if lastmemterm != commit.LastMemTerm || lastmemid != commit.LastMemId{
        // get checkpoint from leader
        //TODO
        req := &pb.Msg{
            Type:CFGCHA,
        }

        client := pb.NewCommpbClient(s.nodes[commit.SrcId].conn)
        ctx,cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
        defer cancel()
        rsp,err := client.GetCheckPoint(ctx,req)
        if err != nil{
            return err
        }
        if rsp.Status != 0{
            log.Errorf("get checkpoint from master failed")
            return errors.New("get checkpoint failed")
        }
        // recover from checkpoint
        err = s.nodeslog.Unmarshal(rsp.Data)
        if err != nil{
            log.Errorf("prepare recover Unmarshal failed:%v",err)
            return err
        }
        s.rebuildConns()
        log.Info("membership data recovered from leader node!!!")
    }

    s.curcommit = commit
    return nil
}
func (s *Server)confirm(commit *pb.Commit)(error){
    if s.curcommit == nil{
        log.Warnf("confirm:nothing to confirm")
        return nil
    }
    if commit.Term != s.curcommit.Term || commit.Id != s.curcommit.Id{
        log.Errorf("prepared and confirmed commits do not match!")
        return errors.New("Prepared and confirmed commits do not match!")
    }
    //apply the commit
    if commit.Type == NORMAL{
        err := s.dblog.Apply(s.curcommit.Term, s.curcommit.Id,s.curcommit.Data)
        if err != nil{
            log.Errorf("confirm -- normal commit failed:%v",err)
            return err
        }
    }else if commit.Type == CFGCHA{
        log.Infof("prepare to commit a message %s",s.curcommit.Data)
        err := s.nodeslog.Apply(s.curcommit.Term, s.curcommit.Id,s.curcommit.Data)
        if err != nil{
            log.Errorf("confirm -- cfg change commit failed:%v",err)
            return err
        }
        s.rebuildConns()
    }
    s.curcommit = nil
    return nil
}
func (s *Server)cancel(commit *pb.Commit)(error){
    s.curcommit = nil
    return nil
}

func (s *Server)changeLeaderID(newlid int64){
    if newlid != s.leaderID{
        s.leaderID = newlid
        err := s.ProposeEmpty()
        if err != nil{
            log.Errorf("changeLeaderID:%v",err)
        }
    }
}
