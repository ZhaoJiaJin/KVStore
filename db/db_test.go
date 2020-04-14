package db

import(
    "testing"
)

func  TestMarshal(t *testing.T){
    s,err := New()
    if err != nil{
        panic(err)
    }
    key := "tt"
    val := "haha"
    s.Apply(&Operation{
        Type:SET,
        Key:key,
        Value:val,
    })
    if s.Get(key) != val{
        t.Fatal("expect "+val)
    }
    d,err := s.GetCheckPoint()
    if err != nil{
        panic(err)
    }
    news, _ := New()
    err = news.RecoverCheckPoint(d)
    if news.Get(key) != val{
        t.Fatal("expect "+val)
    }
}
