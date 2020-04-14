package logging

import(
    "testing"
    "kvstore/db"
    log "github.com/labstack/gommon/log"
    "encoding/json"
    "strconv"
)

func TestCreate(t *testing.T){
    t.Log("create db instance")
    dbins,err := db.New()
    if err != nil{
        t.Fatal(err)
    }

    t.Log("create db instance with log")
    dblog,err := NewDBLogStore(dbins, "test")
    if err != nil{
        t.Fatal(err)
    }
    key := "ttt"
    t.Log("apply log")
    for i:=0; i <= 110; i ++{
        op := db.Operation{
            Type:db.SET,
            Key:"ttt",
            Value:strconv.Itoa(i),
        }
        opb,err := json.Marshal(op)
        if err != nil{
            t.Fatal(err)
        }
        err = dblog.Apply(1,i+1, opb)
        if err != nil{
            t.Fatal(err)
        }
    }
    value := dbins.Get(key)
    if value != "110"{
        t.Fatalf("get expect 110, got %v",value)
    }
    lastterm,lastid := dblog.GetLastCommit()
    if lastterm != 1{
        t.Fatalf("lastterm expect 1, got %v",lastterm)
    }
    if lastid != 111{
        t.Fatalf("lastid expect 111, got %v",lastid)
    }
    
}

func TestRecover(t *testing.T){
    t.Log("create db instance")
    dbins,err := db.New()
    if err != nil{
        t.Fatal(err)
    }

    t.Log("create db instance with log")
    dblog,err := NewDBLogStore(dbins, "test")
    if err != nil{
        t.Fatal(err)
    }
    key := "ttt"
    value := dbins.Get(key)
    if value != "110"{
        t.Fatalf("get expect 110, got %v",value)
    }
    lastterm,lastid := dblog.GetLastCommit()
    if lastterm != 1{
        t.Fatalf("lastterm expect 1, got %v",lastterm)
    }
    if lastid != 111{
        t.Fatalf("lastid expect 111, got %v",lastid)
    }
    log.Info(value," ", lastterm," ",lastid)
}

func TestMarshal(t *testing.T){
    t.Log("create db instance")
    dbins,err := db.New()
    if err != nil{
        t.Fatal(err)
    }

    t.Log("create db instance with log")
    dblog,err := NewDBLogStore(dbins, "test")
    if err != nil{
        t.Fatal(err)
    }
    key := "ttt"
    value := dbins.Get(key)
    if value != "110"{
        t.Fatalf("get expect 110, got %v",value)
    }
    lastterm,lastid := dblog.GetLastCommit()
    if lastterm != 1{
        t.Fatalf("lastterm expect 1, got %v",lastterm)
    }
    if lastid != 111{
        t.Fatalf("lastid expect 111, got %v",lastid)
    }

    backup,err := dblog.Marshal()
    if err != nil{
        t.Fatal(err)
    }
    dbins1,err := db.New()
    if err != nil{
        t.Fatal(err)
    }

    t.Log("create db instance with log")
    dblog1,err := NewDBLogStore(dbins1, "test1")
    if err != nil{
        t.Fatal(err)
    }
    err = dblog1.Unmarshal(backup)
    if err != nil{
        t.Fatal(err)
    }
    value1 := dbins1.Get(key)
    if value1 != "110"{
        t.Fatalf("get expect 110, got %v",value1)
    }
    lastterm1,lastid1 := dblog1.GetLastCommit()
    if lastterm1 != 1{
        t.Fatalf("lastterm expect 1, got %v",lastterm)
    }
    if lastid1 != 111{
        t.Fatalf("lastid expect 111, got %v",lastid)
    }

}
