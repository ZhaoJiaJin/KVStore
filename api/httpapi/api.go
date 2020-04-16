package httpapi

import (
	"flag"
	"fmt"
    "strconv"
	"log"
	"net/http"
    "kvstore/db"
    "kvstore/memdb"
    "kvstore/comm"
    "encoding/json"
)

// DisableWriteViaHTTPGet determines whether create, update, and delete
// requests are acceptable through an http GET request.
var DisableWriteViaHTTPGet bool

var(
    kvdb *db.KVStore 
    commsrv *comm.Server
)

func init() {
	flag.BoolVar(&DisableWriteViaHTTPGet,
		"strict-http-get", false, "disable write requests via http GET method")
}

//Serve starts HTTP server.
func Serve(port int,dbins *db.KVStore,srv *comm.Server) {
    kvdb = dbins
    commsrv = srv
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), nil))
}

// handleRequest handles all incoming HTTP requests and dispatches them to one
// of the handler functions below.
func handleRequest(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if req.URL.Path == "/kvstore/read" {
			handleReadEntryRequest(w, req)
			return
		} else if DisableWriteViaHTTPGet {
			http.Error(w, "Not found", 404)
			return
		}
		// treat request as if method was POST
		fallthrough
	case "POST":
		if req.URL.Path == "/kvstore/create" {
			handleCreateEntryRequest(w, req)
		} else if req.URL.Path == "/kvstore/update" {
			handleUpdateEntryRequest(w, req)
		} else if req.URL.Path == "/kvstore/delete" {
			handleDeleteEntryRequest(w, req)
		} else if req.URL.Path == "/nodes/add" {
			handleAddNodeRequest(w, req)
		} else if req.URL.Path == "/nodes/remove" {
			handleRemoveNodeRequest(w, req)
		} else {
			http.Error(w, "Not found", 404)
		}
	default:
		http.Error(w, "HTTP method not allowed", 405)
	}
}

func handleCreateEntryRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	value := q.Get("value")
	if key == "" || value == "" {
		http.Error(w, "key or value missing in request", 400)
		return
	}

    ope := db.Operation{
        Type:db.SET,
        Key:key,
        Value:value,
    }
    datab,err := json.Marshal(ope)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    err = commsrv.Propose(datab,comm.NORMAL)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
	fmt.Fprintf(w, "created\n%s: %s\n", key, value)
}

func handleReadEntryRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	if key == "" {
		http.Error(w, "key missing in request", 400)
		return
	}

    value := kvdb.Get(key)
	fmt.Fprintf(w, "read\n%s: %s\n", key, value)
}

func handleUpdateEntryRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	value := q.Get("value")
	if key == "" || value == "" {
		http.Error(w, "key or value missing in request", 400)
		return
	}

    ope := db.Operation{
        Type:db.SET,
        Key:key,
        Value:value,
    }
    datab,err := json.Marshal(ope)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    err = commsrv.Propose(datab,comm.NORMAL)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "updated\n%s: %s\n", key, value)
}

func handleDeleteEntryRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	if key == "" {
		http.Error(w, "key missing in request", 400)
		return
	}

     ope := db.Operation{
        Type:db.DEL,
        Key:key,
    }
    datab,err := json.Marshal(ope)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    err = commsrv.Propose(datab,comm.NORMAL)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "deleted %s\n", key)
}

func handleAddNodeRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	id := q.Get("id")
	if id == "" {
		http.Error(w, "id missing in request", 400)
		return
	}
    addr := q.Get("addr")
	if addr == "" {
		http.Error(w, "addr missing in request", 400)
		return
	}
    idint, err := strconv.ParseInt(id,10,64)
    if err != nil{
		http.Error(w, "wrong id type in request", 400)
		return
    }

    ope := memdb.Operation{
        Type:memdb.ADD,
        Key:idint,
        Value:addr,
    }
    datab,err := json.Marshal(ope)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    err = commsrv.Propose(datab,comm.CFGCHA)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "added node: %v %v", id,addr)
}

func handleRemoveNodeRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	id := q.Get("id")
	if id == "" {
		http.Error(w, "id missing in request", 400)
		return
	}
    idint, err := strconv.ParseInt(id,10,64)
    if err != nil{
		http.Error(w, "wrong id type in request", 400)
		return
    }

    ope := memdb.Operation{
        Type:memdb.RM,
        Key:idint,
    }
    datab,err := json.Marshal(ope)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    err = commsrv.Propose(datab,comm.CFGCHA)
    if err != nil{
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "removed node: %v", id)
}
