package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HTTPMirror() http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("[Handler] Error reading body")
		}
		data := string(b)
		fmt.Printf("Url:\t\t%v\nMethod:\t\t%s\nHeaders:\n%s\nBody:\n%s\n\n", r.URL, r.Method, headerString(r.Header), data)
	}
	return http.HandlerFunc(f)
}

func headerString(hdrs http.Header) string {
	buf := bytes.NewBufferString("")
	for k, _ := range hdrs {
		buf.WriteString(fmt.Sprintf("\t%v: %v\n", k, hdrs.Get(k)))
	}
	return buf.String()
}
