package qrest

import (
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
	if adapters.CheckTable(table) {
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

		rw.Write(adapters.Process(table, field, value, page, pagesize))

	} else {
		rw.Write([]byte("Table not present in Database" + config.Conf.MDBDatabase))
	}
}

func ParsePost(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	table := vars["table"]
	v, _ := jason.NewObjectFromReader(req.Body)
	adapters.ProcessPost(v, table)
}

func CreateMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{table}", ParseGet).Methods("GET")
	r.HandleFunc("/{table}", ParsePost).Methods("PUT")
	return r
}

func Start(c config.Config) {
	config.Conf = c

	adapters.CheckDatabase(c.MDBDatabase)

	srv := &http.Server{
		Handler:      CreateMux(),
		Addr:         "127.0.0.1:" + strconv.Itoa(c.HTTPPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
