package adapters

import (
	"fmt"
	"strconv"

	"database/sql"

	"github.com/caarlos0/env"
	//Only init function required
	_ "github.com/go-sql-driver/mysql"
	"github.com/vyasgiridhar/maria-rest/config"
)

//Conn : Creates a new Database connection
//and assigns it to db
func Conn() (db *sql.DB) {
	cfg := config.Config{}
	env.Parse(&cfg)
	var err error
	host := cfg.MDBHost + ":" + strconv.Itoa(cfg.MDBPort)
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.MDBUser, cfg.MDBPass, host, cfg.MDBDatabase))

	if err != nil {
		panic(fmt.Sprintf("Unable to connection to database: %v\n", err))
	}
	return
}
