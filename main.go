package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/microservices-demo/user/login"
	"github.com/microservices-demo/user/register"
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
	log.Infof("Login service running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
