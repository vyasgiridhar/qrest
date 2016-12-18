package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	SelectFrom  = `select * from ?`
	SelectWhere = `where ? = ?`
)

func Process(field string, page, pagesize string) {
}
func Parse(rw http.ResponseWriter, req *http.Request) {
	args := req.URL.Query()
	vars := mux.Vars(req)
	fmt.Println(vars["table"])
	page := args.Get("page")
	pagesize := args.Get("pagesize")
	field := args.Get("field")
	if field != "" {
		Process(field, page, pagesize)
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{table}", Parse)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
