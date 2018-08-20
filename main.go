package main

import "flag"

func main() {
	port := flag.Int("port", 8080, "port to run the mirror server")
	flag.Parse()

	if err := start(Config{Port: *port}); err != nil {
		panic(err)
	}
}
