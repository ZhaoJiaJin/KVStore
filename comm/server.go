package comm

import(
    "sync"
    "time"
    "google.golang.org/grpc"
    "context"
    "errors"
    log "github.com/labstack/gommon/log"
    pb "kvstore/proto/commpb"
    "kvstore/logging"
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
    nodelock sync.RWMutex
    errorC chan error
    dblog *logging.LogStore
    applylock sync.Mutex
    curcommit *pb.Commit
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
    dblog,err := logging.NewDBLogStore(db, datadir)
    if err != nil{
        log.Fatalf("NewServer, create NewDBLogStore failed:%v",err)
    }
    s := &Server{
        lastack:false,
        id:nid,
        term:0,
        voted:false,
        addr:addrs[nid],
        role:Follower,
        nodes:make(map[int64]*Node),
        errorC: make(chan error,1),
        dblog:dblog,
    }
    for i,ad := range addrs{
        s.nodes[int64(i)] = &Node{id:int64(i),addr:ad}
    }
    s.serve(s.addr)
    s.initConn()
    time.Sleep(time.Second * 1)
    s.startHBCheck()
    go s.sendHB()
    return s
}

func (s *Server)initConn(){
    var err error
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())
    s.nodelock.RLock()
    defer s.nodelock.RUnlock()
    for _,nd := range s.nodes{
        nd.conn,err = grpc.Dial(nd.addr,opts...)
        if err != nil{
            log.Fatalf("grpc dial failed:%v",err)
        }
    }
}


func (s *Server) changeTerm(newterm int64){
    s.termLock.Lock()
    defer s.termLock.Unlock()
    if newterm > s.term{
        s.term = newterm
        s.voted = false
        //if s.role == Leader{}
        s.role = Follower
        log.Infof("node %v got bigger term, stepdown from leader to follower",s.id)
    }
}

func (s *Server)getTerm()int64{
    s.termLock.RLock()
    defer s.termLock.RUnlock()
    return s.term
}

func (s *Server)vote(newterm int64)(res bool,oldterm int64){
    s.termLock.Lock()
    defer s.termLock.Unlock()

    oldterm = s.term
    if newterm > oldterm{
        res = true
        s.term = newterm
        s.voted = true
        s.role = Follower
        log.Infof("node %v got bigger term, stepdown from leader to follower",s.id)
    }else{
        if !s.voted{
            res = true
            s.voted = true
            s.role = Follower
            log.Infof("node %v got bigger term, stepdown from leader to follower",s.id)
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


// AddNode add new node to cluster
func (s *Server) AddNode(id int64, addr string)(error){
    s.nodelock.Lock()
    defer s.nodelock.Unlock()

    if _,ok := s.nodes[id];ok{
        return errors.New("node id already exists, please delete the old node before adding it")
    }
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())
    conn,err := grpc.Dial(addr,opts...)
    if err != nil{
        return err
    }
    s.nodes[id] = &Node{
        id:id,
        addr:addr,
        conn:conn,
    }
    return nil
}

// RmNode remove node from cluster
func (s *Server) RmNode(id int64)(error){
    s.nodelock.Lock()
    defer s.nodelock.Unlock()
    n,ok := s.nodes[id]
    if !ok{
        return errors.New("this node does not exist")
    }
    n.conn.Close()
    delete(s.nodes,id)
    return nil
}

func (s *Server)getCheckPoint()([]byte, error){
    return s.dblog.Marshal()
}


//Propose propose a commit to the cluster
//TODO: propose node add/del
func (s *Server)Propose(data []byte)(err error){
    if s.role != Leader{
        log.Infof("node %v is not the leader, pass the commit the leader node %v",s.id,s.leaderID)
        //not leader, need to propose to leader
        commit := &pb.Commit{
            Data:data,
        }
        client := pb.NewCommpbClient(s.nodes[s.leaderID].conn)
        ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
        defer cancel()
        _,err = client.SendToLeader(ctx,commit)
        if err != nil{
            return err
        }
    }else{
        //leader
        log.Infof("node %v is the leader, process the commit",s.id)
        s.applylock.Lock()
        defer s.applylock.Unlock()
        //Get last log id and term from logging module
        lastterm, lastid := s.dblog.GetLastCommit()
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
        }
        //Send commit to all the nodes
        //prepare the commit
        success := 0
        for _,nd := range s.nodes{
            log.Infof("send the commit to node %v",nd.id)
            client := pb.NewCommpbClient(nd.conn)
            ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
            _,err = client.Prepare(ctx,commit)
            if err != nil{
                log.Infof("node %v is not ready to commit, begin to cancel this commit",nd.id)
                cancel()
                break
            }
            log.Infof("node %v is ready to commit",nd.id)
            success ++
            cancel()
        }

        for _,nd := range s.nodes{
            client := pb.NewCommpbClient(nd.conn)
            ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
            if success == len(s.nodes){
                //confirm   
                log.Infof("sending confirm message to node %v",nd.id)
                _,err = client.Confirm(ctx,commit)
            }else{
                //cancel
                log.Infof("sending cancel message to node %v",nd.id)
                _,err = client.Cancel(ctx,commit)
            }
            cancel()
        }
        if success != len(s.nodes){
            return errors.New("some node is unreachable, this commit is discarded")
        }
        // nodes should compare its last term and log id before commit
        return nil
    }
    return nil
}

func (s *Server)prepare(commit *pb.Commit)(error){
    lastterm, lastid := s.dblog.GetLastCommit()
    if lastterm != commit.LastTerm || lastid != commit.LastId{
        // get checkpoint from leader
        req := &pb.Msg{}

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
    err := s.dblog.Apply(s.curcommit.Term, s.curcommit.Id,s.curcommit.Data)
    if err != nil{
        log.Errorf("confirm -- commit failed:%v",err)
        return err
    }
    s.curcommit = nil
    return nil
}
func (s *Server)cancel(commit *pb.Commit)(error){
    s.curcommit = nil
    return nil
}
