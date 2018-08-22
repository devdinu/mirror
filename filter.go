package main

import "net/http"

type predicate func(r *http.Request) bool

func filter(p predicate, next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		if p(r) {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(f)
}

func hasMethod(methods []string) predicate {
	return func(r *http.Request) bool {
		for _, m := range methods {
			if r.Method == m {
				return true
			}
		}
		return false
	}
}

func urlMatching(url string) predicate {
	return func(r *http.Request) bool {
		return r.URL.Path == url
	}
}

func anyMatcher(ps ...predicate) predicate {
	return func(r *http.Request) bool {
		for _, p := range ps {
			if p(r) {
				return true
			}
		}
		return false
	}
}
