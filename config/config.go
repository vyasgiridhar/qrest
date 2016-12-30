package config

type Config struct {
	HTTPPort    int
	MDBHost     string
	MDBPort     int
	MDBUser     string
	MDBPass     string
	MDBDatabase string
}

var Conf Config
