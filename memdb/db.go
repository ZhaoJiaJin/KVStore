package memdb

import(
    "sync"
    "encoding/json"
)

const(
    ADD = iota
    RM
)

type MemStore struct{
    data map[int64]string
    sync.RWMutex
}

type Operation struct{
    Type int
    Key int64
    Value string
}

func New()(*MemStore){
    res := &MemStore{data:make(map[int64]string)}
    return res
}

func (s *MemStore) Apply(data []byte)(error){
    ope := &Operation{}
    err := json.Unmarshal(data,ope)
    if err != nil{
        return err
    }
    //apply the commit
    s.Lock()
    switch ope.Type{
    case ADD:
        s.Set(ope.Key,ope.Value)
    case RM:
        s.Del(ope.Key)
    }
    s.Unlock()
    return nil
}

func (s *MemStore) Set(k int64,v string){
    s.data[k] = v
}

func (s *MemStore) Del(k int64){
    delete(s.data,k)
}

func (s *MemStore) Get(k int64)(v string){
    s.RLock()
    defer s.RUnlock()
    v,_ = s.data[k]
    return
}

func (s *MemStore) GetCheckPoint()([]byte,error){
    s.RLock()
    defer s.RUnlock()
    return json.Marshal(s.data)
}

func (s *MemStore) RecoverCheckPoint(data []byte)(error){
    s.Lock()
    defer s.Unlock()
    return json.Unmarshal(data,&(s.data))
}

func (s *MemStore) Range(f func(key int64, value string) bool){
    s.RLock()
    defer s.RUnlock()
    for k,v := range s.data{
        cnt := f(k,v)
        if !cnt{
            return
        }
    }
}
