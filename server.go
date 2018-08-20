package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Port int
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func start(cfg Config) error {
	addr := cfg.Address()
	fmt.Println("Listening for requests", addr)
	http.HandleFunc("/", HTTPMirror)
	return http.ListenAndServe(addr, nil)
}
