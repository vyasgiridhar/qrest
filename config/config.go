package config

type Config struct {
	HTTPPort    int    `env:"REST_HTTP_PORT" envDefault:"3000"`
	MDBHost     string `env:"REST_MARIA_HOST" envDefault:"127.0.0.1"`
	MDBPort     int    `env:"REST_MARIA_PORT" envDefault:3306`
	MDBUser     string `env:"REST_MARIA_USER"`
	MDBPass     string `env:"REST_MARIA_PASS"`
	MDBDatabase string `env:"REST_MARIA_DATABASE"`
}
