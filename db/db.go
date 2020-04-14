package db

import(
    "sync"
    "encoding/json"
)

const(
    SET = iota
    DEL
)

type KVStore struct{
    data map[string]string
    sync.RWMutex
}

type Operation struct{
    Type int
    Key string
    Value string
}

func New()(*KVStore, error){
    res := &KVStore{data:make(map[string]string)}
    return res,nil
}

func (s *KVStore) Apply(data []byte)(error){
    ope := &Operation{}
    err := json.Unmarshal(data,ope)
    if err != nil{
        return err
    }
    //apply the commit
    s.Lock()
    switch ope.Type{
    case SET:
        s.Set(ope.Key,ope.Value)
    case DEL:
        s.Del(ope.Key)
    }
    s.Unlock()
    return nil
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

