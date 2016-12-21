package adapters

import (
	"fmt"
	"log"
	"strconv"

	"database/sql"

	//Only init function required
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

	db, err = sql.Open("mysql", PrepareConn(database))

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	return
}

func PrepareQuery(table, field string, page, pagesize int) (preparedQuery string) {
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
	return
}

func Process(table, field, value, page, pagesize string) []byte {

	log.Println("Processing", table, field, value, page, pagesize)

	db := Conn(config.Conf.MDBDatabase)

	if db != nil {
		pagei, err := strconv.Atoi(page)
		if err != nil {
			pagesi, err := strconv.Atoi(pagesize)
			if err != nil {
				query := PrepareQuery(table, field, value, pagei, pagesi)
				x, err := db.Query(query, value)
				log.Println(err)
				result, _ := JSONify(x)
				return []byte(result)
			}
		}
	}
	return nil
}
