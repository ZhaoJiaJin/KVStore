package main

import (
  "fmt"
  "time"
  "log"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
  "flag"
)

var(

  client = &http.Client {}
)

func send(port int,key,value string){
    url := fmt.Sprintf("http://localhost:%v/kvstore/update?key=%v&value=%v",port,key,value)
  method := "GET"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  err := writer.Close()
  if err != nil {
    log.Println(err)
    return
  }

  req, err := http.NewRequest(method, url, payload)
  if err != nil {
    log.Println(err)
    return
  }
  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil{
    log.Println(err)
    return
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  log.Println(string(body))

}

func main() {
    port := flag.Int("port",8080,"service port")
    key := flag.String("k","testkey","key")
    value := flag.String("v","testvalue","value")
    flag.Parse()
    for {
        time.Sleep(time.Millisecond* 10)
        send(*port,*key,*value)
    }
  }
