package qrest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/qrest/adapters"
	"github.com/vyasgiridhar/qrest/config"
)

func Parse(rw http.ResponseWriter, req *http.Request) {

	args := req.URL.Query()
	vars := mux.Vars(req)

	table := vars["table"]

	var page, pagesize, field, value string

	for k := range args {

		if k == "page" {
			page = args[k][0]
		} else if k == "pagesize" {
			pagesize = args[k][0]
		} else if adapters.CheckField(vars["table"], k) {
			field = k
			value = args[k][0]
		}
	}

	fmt.Println(args)
	rw.Write(adapters.Process(table, field, value, page, pagesize))
}

func CreateMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{table}", Parse)
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
