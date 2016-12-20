package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const SelectFrom = `select * from `

func CheckField(table, field string) bool {
	check := "select COLUMN_NAME from columns where TABLE_NAME = ?"
	db := Conn("INFORMATION_SCHEMA")
	defer db.Close()
	rs, err := db.Query(check, table)
	if err != nil {
		log.Println("Error")
	}
	columnName := ""
	for rs.Next() {
		rs.Scan(&columnName)
		if columnName == field {
			return true
		}
	}
	return false
}

func Parse(rw http.ResponseWriter, req *http.Request) {
	args := req.URL.Query()
	vars := mux.Vars(req)
	table := vars["table"]
	var page, pagesize, field, value string
	for k := range args {
		fmt.Println(k, " ", args[k][0])
		if k == "page" {
			page = args[k][0]
		} else if k == "pagesize" {
			pagesize = args[k][0]
		} else if CheckField(vars["table"], k) {
			field = k
			value = args[k][0]
		}
	}

	fmt.Println(args)
	rw.Write(Process(table, field, value, page, pagesize))
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
