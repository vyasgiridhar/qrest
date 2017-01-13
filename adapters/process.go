package adapters

import (
	"log"

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

func ProcessPost(j *jason.Object, table string) string {
	for x, value := range j.Map() {
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
