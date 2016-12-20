package config

type Config struct {
	HTTPPort    int
	MDBHost     string
	MDBPort     int
	MDBUser     string
	MDBPass     string
	MDBDatabase string
}

//Conf : Global configuration
//For debugging
var Conf = Config{
	8000,
	"127.0.0.1",
	3306,
	"root",
	"spd",
	"Football",
}
