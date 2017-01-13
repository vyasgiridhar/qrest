package adapters

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"database/sql"

	"github.com/antonholmquist/jason"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vyasgiridhar/qrest/config"
)

func CheckDatabase(name string) {
	db := Conn(name)
	if db.Ping() != nil {
		panic("SQL database not present or initialized")
	}
}

func PrepareConn(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Conf.MDBUser, config.Conf.MDBPass, config.Conf.MDBHost, config.Conf.MDBPort, database)
}

//Conn : Creates a new Database connection
//and assigns it to db
func Conn(database string) (db *sql.DB) {
	var err error
	db, err = sql.Open("mysql", PrepareConn(database))

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	fmt.Println("Connected to " + config.Conf.MDBDatabase)
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
			}
			log.Println("error at parsing pagesize")
		} else {
			log.Println("error at parsing page")
		}
	} else {
		log.Println("Error while creating db conn")
	}
	return nil
}

func PrepareInsertQuery(query string, j map[string]*jason.Value) (statement string) {

	query += " values ('"
	values := make([]string, 0)
	var val string
	for _, value := range j {
		val, _ = value.String()
		if val == "" {
			integer, _ := value.Int64()
			val = strconv.Itoa(int(integer))
		}
		values = append(values, val)
	}
	query += strings.Join(values, "', '")
	query += "')"
	log.Println(query)
	statement = query
	return
}

func Insertinto(table string, j *jason.Object) (sucess bool) {
	x := j.Map()
	query := "insert into " + table + "("
	i := 0
	for key := range x {
		if i == 0 {
			query += key
			i++
		} else {
			query += "," + key
		}
	}
	query += ")"
	query = PrepareInsertQuery(query, x)
	log.Println(query)
	return
}
