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
	var predicates []predicate
	if len(config.Methods()) != 0 {
		predicates = append(predicates, hasMethod(config.Methods()))
	}
	if len(predicates) != 0 {
		handler = filter(anyMatcher(predicates...), handler)
	}

	http.Handle("/", handler)
	return http.ListenAndServe(addr, nil)
}
