package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// Version number constant.
const Version = "0.0.5"

// Homepage url.
const Homepage = "https://github.com/servian/https-echo"

var (
	httpAddr = flag.String("listen", ":80", "Listen address")
	destPort = flag.Int("port", -1, "Destination port")
	versDisp = flag.Bool("version", false, "Display version")
)

func redirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Server", "Https-echo/"+Version+" (+"+Homepage+")")
	if "/health" == req.URL.Path {
		io.WriteString(w, "Healthy.\n")
		return
	}
	hostname := strings.Split(req.Host, ":")
	dps := ""
	if *destPort > 0 {
		dps = fmt.Sprintf(":%d", *destPort)
	}
	target := "https://" + hostname[0] + dps + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s from: %s", target, req.RemoteAddr)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func main() {
	flag.Parse()

	if *versDisp {
		fmt.Printf("Version: v%s %s\n", Version, runtime.Version())
		fmt.Printf("Home Page: %s\n", Homepage)
		os.Exit(0)
	}

	log.Printf("Listening for incoming requests on TCP port '%s'...", *httpAddr)
	err := http.ListenAndServe(*httpAddr, http.HandlerFunc(redirect))
	if err != nil {
		log.Fatal(err)
	}
}
