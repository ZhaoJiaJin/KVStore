package comm

import(
    "sync"
    "time"
    "google.golang.org/grpc"
    log "github.com/labstack/gommon/log"
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
    id int
    term int
    voted bool
    termLock sync.RWMutex
    role ROLE
    addr string
    leaderID int
    hbtimeout int64
    lastack bool
    nodes map[int]*Node //id:string
    errorC chan error
}

// Node represent a node
type Node struct{
    id int
    addr string
    conn *grpc.ClientConn
}

//NewServer return a server
func NewServer(nid int, addrs []string)*Server{
    if nid > len(addrs){
        log.Fatalf("id should be smaller than cluster addresses")
    }
    s := &Server{
        lastack:false,
        id:nid,
        term:0,
        voted:false,
        addr:addrs[nid],
        role:Follower,
        nodes:make(map[int]*Node),
        errorC: make(chan error,1),
    }
    for i,ad := range addrs{
        s.nodes[i] = &Node{id:i,addr:ad}
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
    for _,nd := range s.nodes{
        nd.conn,err = grpc.Dial(nd.addr,opts...)
        if err != nil{
            log.Fatalf("grpc dial failed:%v",err)
        }
    }
}


func (s *Server) changeTerm(newterm int){
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

func (s *Server)getTerm()int{
    s.termLock.RLock()
    defer s.termLock.RUnlock()
    return s.term
}

func (s *Server)vote(newterm int)(res bool,oldterm int){
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
