package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/devdinu/mirror/config"
)

type predicate func(r *http.Request) bool

type proxy struct {
	backend  string
	urlRegex *regexp.Regexp
	client   *http.Client
	next     http.Handler
}

func New(pc config.Proxy, next http.Handler) http.Handler {
	urlRegex := regexp.MustCompile(pc.MatchingUrl)
	p := proxy{backend: pc.Backend, urlRegex: urlRegex, client: http.DefaultClient, next: next}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// could put matching requests, replace with standard predicate
		fmt.Printf("[Proxy] matching url: %s with %s\n", r.URL.Path, p.urlRegex)
		if p.urlRegex.MatchString(r.URL.Path) {
			servReq, _ := http.NewRequest(r.Method, p.backend, r.Body)
			resp, err := p.client.Do(servReq)
			if err != nil {
				fmt.Printf("[Proxy] proxying error: %v\n", err)
			}
			w.WriteHeader(resp.StatusCode)
			data, _ := ioutil.ReadAll(resp.Body)
			w.Write(data)
		} else {
			p.next.ServeHTTP(w, r)
		}
	})
}
