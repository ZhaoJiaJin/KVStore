package comm

import(
    pb "kvstore/proto/commpb"
    "google.golang.org/grpc"
    log "github.com/labstack/gommon/log"
    "net"
    "errors"
    "context"
)

var(
    ErrCommDisabled = errors.New("Communication is disabled")
)

func (s *Server)DisableGRPC(){
    log.Warn("Disable All Communications !!!")
    s.grpcdisabled = true
    s.nodelock.Lock()
    defer s.nodelock.Unlock()
    for _,nd := range s.nodes{
        nd.conn.Close()
    }
}

func (s *Server)EnableGRPC(){
    log.Warn("Enable All Communications !!!")
    s.grpcdisabled = false
    s.rebuildConns()
}

// AskForVote implements CommpbServer interface, process vote request
func (s *Server) AskForVote(ctx context.Context,req *pb.VoteReq) (res *pb.VoteRsp,err error){
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    res = &pb.VoteRsp{}
    voteYes,_ := s.vote(req)
    if voteYes{
       res.Vtres = pb.VoteRsp_YES
    }else{
       res.Vtres = pb.VoteRsp_NO
    }
    return
}

// HeartBeat implements CommpbServer interface, process heartbeat request
func (s *Server) HeartBeat(ctx context.Context,req *pb.HBReq) (*pb.HBRsp, error){
    //log.Infof("node %v receive heartbeat",s.id)
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    s.lastack = true
    s.changeLeaderID(req.Id)
    s.changeTerm(req.Term)
    if s.role == Candidate{
        s.role = Follower
    }
    return &pb.HBRsp{
        Term: int64(s.term),
    }, nil
}

// GetCheckPoint handle GetCheckpoint requests from followers
func (s *Server) GetCheckPoint(ctx context.Context, req *pb.Msg)(*pb.CP, error){
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    data,err := s.getCheckPoint(req.Type)
    if err != nil{
        log.Warnf("GetCheckPoint %v",err)
        return &pb.CP{
            Status:1,
        },nil
    }
    return  &pb.CP{
        Data:data,
        Status:0,
    }, nil
}


// Prepare handle a prepare commit message from leader
func (s *Server) Prepare(ctx context.Context, req *pb.Commit)(*pb.Msg, error){
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    res := &pb.Msg{}
    return res,s.prepare(req)
}


// Confirm handle a confirm commit message from leader
func (s *Server) Confirm(ctx context.Context, req *pb.Commit)(*pb.Msg, error){
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    res := &pb.Msg{}
    return res,s.confirm(req)
}

// Cancel handle a cancel commit message from leader
func (s *Server) Cancel(ctx context.Context, req *pb.Commit)(*pb.Msg, error){
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    res := &pb.Msg{}
    return res,s.cancel(req)
}

// SendToLeader handle a commit send from a follower to leader
func (s *Server) SendToLeader(ctx context.Context, req *pb.Commit)(*pb.Msg, error){
    if s.grpcdisabled{
        return nil,ErrCommDisabled
    }
    res := &pb.Msg{}
    err := s.Propose(req.Data,req.Type)
    return res,err
}

// serve start grpc server
func (s *Server)serve(addr string){
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    go func(){
	    var opts []grpc.ServerOption
        grpcServer := grpc.NewServer(opts...)
	    pb.RegisterCommpbServer(grpcServer, s)
        if err := grpcServer.Serve(lis); err != nil{
            log.Fatalf("fail to start grpc server:%v",err)
        }
    }()
}


