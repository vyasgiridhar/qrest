package adapters

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
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
