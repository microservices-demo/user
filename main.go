package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/db/mongodb"
	"github.com/microservices-demo/user/login"
)

var dev bool
var verbose bool
var port string
var acc string

func init() {
	flag.StringVar(&port, "port", "8084", "Port on which to run")
	flag.BoolVar(&verbose, "verbose", false, "Verbose logging")
	db.Register("mongodb", &mongodb.Mongo{})
}

func main() {

	flag.Parse()
	// Mechanical stuff.
	errc := make(chan error)
	ctx := context.Background()

	// Log domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
	err := db.Init()
	if err != nil {
		logger.Log(err)
		os.Exit(1)
	}

	// Service domain.
	var service login.Service
	{
		service = login.NewFixedService()
		service = login.LoggingMiddleware(logger)(service)
	}

	// Endpoint domain.
	endpoints := login.MakeEndpoints(service)

	// Create and launch the HTTP server.
	go func() {
		logger.Log("transport", "HTTP", "port", port)
		handler := login.MakeHTTPHandler(ctx, endpoints, logger)
		errc <- http.ListenAndServe(fmt.Sprintf(":%v", port), handler)
	}()

	// Capture interrupts.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
}

/*
	http.HandleFunc("/login", login.Handle)
	http.HandleFunc("/register", register.Handle)
	log.Infof("Login service running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
*/
