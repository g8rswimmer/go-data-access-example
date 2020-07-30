package env

import (
	"os"
	"strconv"
	"time"
)

// HTTP contains all of the http configuration
type HTTP struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Config contains all of the configuration
type Config struct {
	HTTP *HTTP
}

const (
	httpPort    = "HTTP_PORT"
	httpReadTO  = "HTTP_READ_TO"
	httpWriteTO = "HTTP_WRITE_TO"
)

// Load will read the environmental variables with defaults
func Load() *Config {
	return &Config{
		HTTP: &HTTP{
			Port:         port(),
			ReadTimeout:  readTO(),
			WriteTimeout: writeTO(),
		},
	}
}

func port() string {
	p := os.Getenv(httpPort)
	if len(p) == 0 {
		p = "8080"
	}
	return p
}

func readTO() time.Duration {
	rto := os.Getenv(httpReadTO)
	if len(rto) == 0 {
		rto = "10"
	}
	return timeout(rto)
}

func writeTO() time.Duration {
	wto := os.Getenv(httpReadTO)
	if len(wto) == 0 {
		wto = "10"
	}
	return timeout(wto)
}

func timeout(to string) time.Duration {
	t, err := strconv.Atoi(to)
	if err != nil {
		panic(err)
	}
	return time.Duration(t) * time.Second
}
