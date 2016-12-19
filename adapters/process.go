package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/maria-rest/config"
)

const (
	SelectFrom  = `select * from ?`
	SelectWhere = `where ? = ?`
)

func PrepareConn(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Conf.MDBUser, config.Conf.MDBPass, config.Conf.MDBHost, config.Conf.MDBPort, database)
}
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

func Conn(database string) (db *sql.DB) {
	var err error

	db, err = sql.Open("mysql", PrepareConn(database))

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	return
}
func PrepareQuery(table,field,value,page,pagesize string) (preparedQuery string) {

}

func Process(table, field, value, page, pagesize string) []byte {
	fmt.Println(table, field, value, page, pagesize)
	db := Conn("Football")
	query := PrepareQuery(table,field,value,page,pagesize string)
	x, err := db.Query(query)
	fmt.Println(err)
	result, _ := JSONify(x)
	return []byte(result)
}

func JSONify(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
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
