package main

import (
	"flag"
	"kvstore/api/httpapi"
    "strings"
    "kvstore/comm"
	log "github.com/labstack/gommon/log"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	dbport := flag.Int("port", 9121, "database server port")
	//join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	log.Infof("local port:%v", *dbport)
	comm.NewServer(*id, strings.Split(*cluster, ","))
	httpapi.Serve(*dbport)
	select {}
	/*proposeC := make(chan string)
		defer close(proposeC)
		confChangeC := make(chan raftpb.ConfChange)
		defer close(confChangeC)

		// raft provides a commit stream for the proposals from the http api
	    //db.Init(*id)
		var docdb *db.DB
		getSnapshot := func() ([]byte, error) { return docdb.GetSnapshot() }
		commitC, errorC, snapshotterReady := raft.NewRaftNode(*id, strings.Split(*cluster, ","), *join, getSnapshot, proposeC, confChangeC)

		docdb = db.NewDBWithRaft(<-snapshotterReady, proposeC, commitC, errorC)

		// the key-value http handler will propose updates to raft
		httpapi.ServeHttpAPI(docdb, *dbport, confChangeC, errorC,"")*/
}
