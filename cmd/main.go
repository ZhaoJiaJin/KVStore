package main

import (
	"flag"
	"kvstore/api/httpapi"
    "strings"
    "kvstore/comm"
    "kvstore/db"
	log "github.com/labstack/gommon/log"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	datadir := flag.String("dir", "data", "data directory")
	id := flag.Int("id", 1, "node ID")
	dbport := flag.Int("port", 9121, "database server port")
	//join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

    kvdb,err := db.New()
    if err != nil{
        panic(err)   
    }
	log.Infof("local port:%v", *dbport)
    s := comm.NewServer(int64(*id), strings.Split(*cluster, ","), kvdb, *datadir)
	httpapi.Serve(*dbport,kvdb,s)
}
