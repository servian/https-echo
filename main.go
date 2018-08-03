package main

import (
	"log"
	"net/http"
	"strings"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	hostname := strings.Split(req.Host, ":")
	target := "https://" + hostname[0] + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func main() {
	http.ListenAndServe(":80", http.HandlerFunc(redirect))
}
