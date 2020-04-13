package db

import(
    "sync"
    log "github.com/labstack/gommon/log"
    "encoding/json"
    "os"
    "strings"
    "path/filepath"
    "io/ioutil"
)

const(
    SET = iota
    DEL
)

type KVStore struct{
    data map[string]string
    count int
    sync.RWMutex
    logf *os.File
}

type Operation struct{
    Type int
    Key string
    Value string
}

func New(dir string)(*KVStore, error){
    walfile := filepath.Join(dir,"wal")
    res := &KVStore{data:make(map[string]string)}
    //try to recover from wal file
    content,err := ioutil.ReadFile(walfile)
    if err == nil{
        // wal file exists
        // try to recover
        lines := strings.Split(string(content),"\n")
        for i,line := range lines{
            line = strings.TrimSpace(line)
            if len(line) == 0{
                continue
            }
            if i == 0{
                // first line is check point
                err = res.RecoverCheckPoint([]byte(line))
                if err != nil{
                    return res, err
                }
            }else{
                // other lines are commits
                o := &Operation{}
                err = json.Unmarshal([]byte(line), o)
                if err != nil{
                    log.Warnf("New:%v",err)
                    continue
                }
                res.apply(o)
            }
        }
    }
    //take a new checkpoint after recovery or create a new server
    f,err := os.Create(walfile)
    if err != nil{
        return res,err
    }
    res.logf = f
    cpdata,err := res.GetCheckPoint()
    if err != nil{
        return res,err
    }
    _,err = res.logf.Write(cpdata)
    if err != nil{
        return res,err
    }
    _,err = res.logf.WriteString("\n")
    if err != nil{
        return res,err
    }
    res.count = 0
    return res,nil
}

func (s *KVStore) Apply(ope *Operation)(error){
    //save the commit to disk
    opeb, _ := json.Marshal(ope)
    _,err := s.logf.Write(opeb)
    if err != nil{
        return err
    }
    _,err = s.logf.WriteString("\n")
    if err != nil{
        return err
    }
    s.apply(ope)
    //take a checkpoint every 100 commits
    //and discard previous commits
    if s.count % 100 == 0{
        log.Info("taking a checkpoint")
        err = s.logf.Truncate(0)
        if err != nil{
            return err
        }
        _,err := s.logf.Seek(0,0)
        if err != nil{
            return err
        }
        cpdata,err := s.GetCheckPoint()
        if err != nil{
            return err
        }
        _,err = s.logf.Write(cpdata)
        if err != nil{
            return err
        }
        _,err = s.logf.WriteString("\n")
        if err != nil{
            return err
        }
    }
    return nil
}



func (s *KVStore) apply(ope *Operation){
    //apply the commit
    s.Lock()
    switch ope.Type{
    case SET:
        s.Set(ope.Key,ope.Value)
    case DEL:
        s.Del(ope.Key)
    }
    s.count ++
    s.Unlock()
}

func (s *KVStore) Set(k,v string){
    s.data[k] = v
}

func (s *KVStore) Del(k string){
    delete(s.data,k)
}

func (s *KVStore) Get(k string)(v string){
    s.RLock()
    defer s.RUnlock()
    v,_ = s.data[k]
    return
}

func (s *KVStore) GetCheckPoint()([]byte,error){
    s.RLock()
    defer s.RUnlock()
    return json.Marshal(s.data)
}

func (s *KVStore) RecoverCheckPoint(data []byte)(error){
    s.Lock()
    defer s.Unlock()
    return json.Unmarshal(data,&(s.data))
}


func (s *KVStore)Close(){
    s.logf.Close()
}
