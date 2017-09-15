package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"git.kono.sh/bkono/trackem"
)

var (
	port = flag.String("port", "", "port to listen on")
)

func main() {
	flag.Parse()
	parseEnv()

	if len(*port) == 0 {
		log.Fatal("port must be provided")
	}

	addr := ":" + *port
	http.HandleFunc("/em.gif", trackem.GetEMReq)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func parseEnv() {
	if len(*port) == 0 {
		*port = os.Getenv("PORT")
	}
}
