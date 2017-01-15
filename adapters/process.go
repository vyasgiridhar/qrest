package adapters

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/antonholmquist/jason"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vyasgiridhar/qrest/config"
)

const SelectFrom = `select * from `

func CheckField(table, field string) bool {
	check := "select COLUMN_NAME from columns where TABLE_NAME = ?"
	db := Conn("INFORMATION_SCHEMA")
	defer db.Close()
	rs, err := db.Query(check, table)
	if err != nil {
		log.Println(err)
		return false
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

func CheckTable(table string) bool {
	query := "show tables"
	db := Conn(config.Conf.MDBDatabase)
	defer db.Close()
	rs, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return false
	}
	tableName := ""
	for rs.Next() {
		rs.Scan(&tableName)
		if tableName == table {
			return true
		}
	}
	return false
}

func ProcessPut(j *jason.Object, table string) string {
	for x := range j.Map() {
		if CheckField("Player", x) {
		} else {
			return "json invalid"
		}
	}

	if Insertinto(table, j) {
		return "inserted"
	}
	return "json invalid"
}

func ProcessGet(table, field, value, page, pagesize string) []byte {

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

func ProcessPost(j *jason.Object, table string) string {

}
