package comm

import(
    pb "kvstore/proto/commpb"
    "google.golang.org/grpc"
    log "github.com/labstack/gommon/log"
    "net"
    "context"
)

// AskForVote implements CommpbServer interface, process vote request
func (s *Server) AskForVote(ctx context.Context,req *pb.VoteReq) (res *pb.VoteRsp,err error){
    res = &pb.VoteRsp{}
    voteYes,_ := s.vote(int(req.Term))
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
    s.lastack = true
    s.leaderID = int(req.Id)
    s.changeTerm(int(req.Term))
    if s.role == Candidate{
        s.role = Follower
    }
    return &pb.HBRsp{
        Term: int64(s.term),
    }, nil
}

// GetCheckPoint handle GetCheckpoint requests from followers
func (s *Server) GetCheckPoint(ctx context.Context, req *pb.Msg)(*pb.CP, error){
    data,err := s.getCheckPoint()
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
}


// Confirm handle a confirm commit message from leader
func (s *Server) Confirm(ctx context.Context, req *pb.Commit)(*pb.Msg, error){
}

// SendToLeader handle a commit send from a follower to leader
func (s *Server) SendToLeader(ctx context.Context, req *pb.Commit)(*pb.Msg, error){

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


