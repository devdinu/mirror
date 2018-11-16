package main

import (
	"github.com/devdinu/mirror/config"
)

//TODO: move it to cmd/mirror package
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := start(config.Address()); err != nil {
		panic(err)
	}
}
