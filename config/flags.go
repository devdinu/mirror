package config

import (
	"flag"
	"fmt"
	"strings"
)

type config struct {
	port    int
	methods []string
}

var cfg config

func Load() error {
	port := flag.Int("port", 8080, "port to run the mirror server")
	methods := flag.String("methods", "", "show only request with given method as comma separated values, GET,POST,PUT...")
	flag.Parse()

	cfg = config{
		port:    *port,
		methods: filterMethods(*methods),
	}
	return nil
}

func Address() string {
	return fmt.Sprintf(":%d", cfg.port)
}

func Methods() []string {
	return cfg.methods
}

func filterMethods(methods string) []string {
	//TODO: validate the methods and return error
	// add cases for matching http.MethodGet
	return strings.Split(methods, ",")
}
