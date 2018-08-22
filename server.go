package main

import (
	"fmt"
	"net/http"

	"github.com/devdinu/mirror/config"
)

type Config struct {
	Port int
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func start(addr string) error {
	fmt.Println("Listening for requests", addr)
	handler := HTTPMirror()
	handler = filter(
		anyMatcher(hasMethod(config.Methods())),
		handler)

	http.Handle("/", handler)
	return http.ListenAndServe(addr, nil)
}
