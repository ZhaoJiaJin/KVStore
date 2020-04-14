package comm

import(
    "math/rand"
    log "github.com/labstack/gommon/log"
    "time"
    pb "kvstore/proto/commpb"
    "context"
)


const(
    // HBINTV heartbeat interval
    HBINTV = 50 * time.Millisecond
)

//startHBCheck start heartbeat check
func (s *Server)startHBCheck(){
    rand.Seed(time.Now().UnixNano())
    // generate a random number between 150 and 300
    // https://en.wikipedia.org/wiki/Raft_(computer_science)
    s.hbtimeout = int64(HBLOWER + rand.Intn(HBUPPER - HBLOWER))
    log.Infof("heartbeat timeout:%v",s.hbtimeout)
    go s.checkhb(time.Duration(s.hbtimeout)*time.Millisecond)
}

func (s *Server) checkhb(dura time.Duration){
    for{
        time.Sleep(dura)
        if s.role == Follower{
            if ! s.lastack{
                s.applyLeader()
            }
            s.lastack = false
        }
    }
}

func (s *Server) applyLeader(){
    s.role = Candidate
    s.incTerm()
    log.Infof("node %v apply for leader at term %v",s.id, s.term)
    s.nodelock.RLock()
    defer s.nodelock.RUnlock()
    nodesum := len(s.nodes)
    yesvotes := 1 // I give myself a yes vote
    for id, nd := range s.nodes{
        if nd.id == s.id{
            continue // skip myself
        }
        client := pb.NewCommpbClient(nd.conn)
        ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
        defer cancel()

        lastLogTerm,lastLogID := s.dblog.GetLastCommit()
        rsp,err := client.AskForVote(ctx, &pb.VoteReq{
            Term:int64(s.getTerm()), 
            Id:int64(s.id),
            LastLogTerm:lastLogTerm,
            LastLogId:lastLogID,
        })
        if err != nil{
            log.Errorf("fail to AskForVote:%v %v %v", id, nd.addr, err)
        }else{
            if rsp.Vtres == pb.VoteRsp_YES{
                yesvotes ++
                log.Infof("node %v vote yes",nd.id)
            }else{
                log.Infof("node %v vote no",nd.id)
            }
        }
        if s.role != Candidate{
            break
        }
    }
    if yesvotes > nodesum / 2{
        log.Infof("node %v get %v votes our of %v nodes, will become leader at term %v",s.id, yesvotes, nodesum, s.term)
        s.role = Leader
    }else{
        log.Infof("node %v get %v votes our of %v nodes, will become follower",s.id, yesvotes, nodesum)
        s.role = Follower
    }
}

func (s *Server) sendHB(){
    for{
        time.Sleep(HBINTV)
        //TODO: maybe should use different goroutine for different nodes.
        if s.role == Leader{
            s.nodelock.RLock()
            for id,nd := range s.nodes{
                if nd.id == s.id{
                    continue//skip myself
                }
                client := pb.NewCommpbClient(nd.conn)
                ctx,cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
                defer cancel()
                rsp,err := client.HeartBeat(ctx,&pb.HBReq{Term:int64(s.getTerm()),Id:int64(s.id)})
                if err != nil{
                    id = id
                    //log.Errorf("fail to send heartbeat to %v %v %v",id,nd.addr,err)
                }else{
                    s.changeTerm(rsp.Term)
                }
            }
            s.nodelock.RUnlock()
        }
    }
}
