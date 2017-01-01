package qrest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/qrest/adapters"
	"github.com/vyasgiridhar/qrest/config"
)

func ParseGet(rw http.ResponseWriter, req *http.Request) {

	args := req.URL.Query()
	vars := mux.Vars(req)

	table := vars["table"]

	var page, pagesize, field, value string

	for k := range args {

		if k == "page" {
			page = args[k][0]
		} else if k == "pagesize" {
			pagesize = args[k][0]
		} else if adapters.CheckField(table, k) {
			field = k
			value = args[k][0]
		}
	}

	fmt.Println(args)
	rw.Write(adapters.Process(table, field, value, page, pagesize))
}

func ParsePost(rw http.ResponseWriter, req *http.Request) {
	//	vars := mux.Vars(req)

	//table := vars["table"]
	v, _ := jason.NewObjectFromReader(req.Body)
	adapters.ProcessPost(v)
}

func CreateMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{table}", ParseGet)
	r.HandleFunc("/{table}/insert", ParsePost)
	return r
}

func Start(c config.Config) {
	config.Conf = c
	srv := &http.Server{
		Handler:      CreateMux(),
		Addr:         "127.0.0.1:" + strconv.Itoa(c.HTTPPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
