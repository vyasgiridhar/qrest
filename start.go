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

		rw.Write(adapters.ProcessGet(table, field, value, page, pagesize))

	} else {
		rw.Write([]byte("Table not present in Database" + config.Conf.MDBDatabase))
	}
}

func ParsePut(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	table := vars["table"]
	v, err := jason.NewObjectFromReader(req.Body)
	if err != nil {
		log.Println("Malformed JSON")
		rw.Write([]byte("Malformed JSON"))
	}
	adapters.ProcessPut(v, table)
}

func ParsePost(rw http.ResponseWriter, req *http.Request) {
	args := req.URL.Query()
	vars := mux.Vars(req)

	table := vars["table"]
	v, err := jason.NewObjectFromReader(req.Body)
	if err != nil {
		log.Println("Malformed JSON")
		rw.Write([]byte("Malformed JSON"))
	}
	adapters.PorcessPost(v, table)
}

func CreateMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{table}", ParseGet).Methods("GET")
	r.HandleFunc("/{table}", ParsePut).Methods("PUT")
	r.HandleFunc("/{table}")
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
