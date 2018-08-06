package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const Version = "0.0.1"

var (
	httpAddr = flag.String("listen", ":80", "Listen address")
	versDisp = flag.Bool("version", false, "Display version")
)

func redirect(w http.ResponseWriter, req *http.Request) {
	hostname := strings.Split(req.Host, ":")
	target := "https://" + hostname[0] + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s from: ", target, req.RemoteAddr)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func main() {
	flag.Parse()

	if *versDisp {
		fmt.Printf("Version: v%s\n", Version)
		os.Exit(0)
	}

	http.ListenAndServe(*httpAddr, http.HandlerFunc(redirect))
}
