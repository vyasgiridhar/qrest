package adapters

import (
	"fmt"
	"log"
	"strconv"

	"database/sql"

	"github.com/antonholmquist/jason"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vyasgiridhar/qrest/config"
)

func PrepareConn(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Conf.MDBUser, config.Conf.MDBPass, config.Conf.MDBHost, config.Conf.MDBPort, database)
}

//Query : Returns json data for a Query
func Query(query string) (jsonData []byte, err error) {
	return nil, nil
}

//Conn : Creates a new Database connection
//and assigns it to db
func Conn(database string) (db *sql.DB) {
	var err error
	fmt.Println(PrepareConn(database))
	db, err = sql.Open("mysql", PrepareConn(database))

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	return
}

func PrepareSelectQuery(table, field string, page, pagesize int) (preparedQuery string) {
	preparedQuery = ""
	if table != "" {
		preparedQuery = fmt.Sprintf("%s %s", SelectFrom, table)
	}
	if field != "" {
		preparedQuery = fmt.Sprintf("%s where %s = ?", preparedQuery, field)
	}
	if page != 0 && pagesize != 0 {
		preparedQuery = fmt.Sprintf("%s limit %d offset %d", preparedQuery, pagesize, page*pagesize)
	}
	log.Println(field)
	return
}

func Process(table, field, value, page, pagesize string) []byte {

	log.Println("Processing", table, field, value, page, pagesize)

	db := Conn(config.Conf.MDBDatabase)

	var x *sql.Rows
	var err error
	var pagei int
	var pagesi int

	if db != nil {
		pagei, err = strconv.Atoi(page)
		if err == nil || pagei == 0 {
			pagesi, err = strconv.Atoi(pagesize)
			if err == nil || pagesi == 0 {
				query := PrepareSelectQuery(table, field, pagei, pagesi)
				fmt.Println(query)
				fmt.Println(value)
				if value != "" {
					x, err = db.Query(query, value)
				} else {
					x, err = db.Query(query)
				}
				if err != nil {
					log.Println(err)
					return nil
				}
				result, _ := JSONify(x)
				return []byte(result)
			} else {
				log.Println("error at parsing pagesize")
			}
		} else {
			log.Println("error at parsing page")
		}
	} else {
		log.Println("Error while creating db conn")
	}
	return nil
}

func ProcessPost(j *jason.Object) string {

	for x, value := range j.Map() {
		fmt.Println(CheckField("Player", x), value)
	}
	return ""
}

func PrepareInsertQuery(table, field string, data []byte) (statement string) {
	statement = ""
	return
}
