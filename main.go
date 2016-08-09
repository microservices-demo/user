package main

import (
	"flag"
	"log"
	"net/http"

	"./login"
	"./register"
)

var dev bool
var verbose bool
var port string
var acc string

func init() {
	flag.StringVar(&port, "port", "8084", "Port on which to run")
	flag.BoolVar(&verbose, "verbose", false, "Verbose logging")
}

func main() {

	flag.Parse()
	http.HandleFunc("/login", login.Handle)
	http.HandleFunc("/register", register.Handle)
	log.Printf("Login service running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
