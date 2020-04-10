package httpapi

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// DisableWriteViaHTTPGet determines whether create, update, and delete
// requests are acceptable through an http GET request.
var DisableWriteViaHTTPGet bool

func init() {
	flag.BoolVar(&DisableWriteViaHTTPGet,
		"strict-http-get", false, "disable write requests via http GET method")
}

//Serve starts HTTP server.
func Serve(port int) {
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

	// TODO: create entry in DB
	fmt.Fprintf(w, "created\n%s: %s\n", key, value)
}

func handleReadEntryRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	if key == "" {
		http.Error(w, "key missing in request", 400)
		return
	}

	// TODO: look up entry in DB
	value := "🤷"
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

	// TODO: update entry in DB
	fmt.Fprintf(w, "updated\n%s: %s\n", key, value)
}

func handleDeleteEntryRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	if key == "" {
		http.Error(w, "key missing in request", 400)
		return
	}

	// TODO: delete entry in DB
	value := "🤷"
	fmt.Fprintf(w, "deleted\n%s: %s\n", key, value)
}

func handleAddNodeRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	node := q.Get("node")
	if node == "" {
		http.Error(w, "node missing in request", 400)
		return
	}

	// TODO: add node as new member
	fmt.Fprintf(w, "added node: %s", node)
}

func handleRemoveNodeRequest(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	node := q.Get("node")
	if node == "" {
		http.Error(w, "node missing in request", 400)
		return
	}

	// TODO: remove node from members
	fmt.Fprintf(w, "removed node: %s", node)
}
