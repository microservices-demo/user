package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	corelog "log"

	"github.com/go-kit/kit/log"
	"github.com/microservices-demo/user/api"
	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/db/mongodb"
	"golang.org/x/net/context"
)

var dev bool
var port string
var acc string

func init() {
	flag.StringVar(&port, "port", "8084", "Port on which to run")
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
		corelog.Fatal(err)
	}

	// Service domain.
	var service api.Service
	{
		service = api.NewFixedService()
		service = api.LoggingMiddleware(logger)(service)
	}

	// Endpoint domain.
	endpoints := api.MakeEndpoints(service)

	// Create and launch the HTTP server.
	go func() {
		logger.Log("transport", "HTTP", "port", port)
		handler := api.MakeHTTPHandler(ctx, endpoints, logger)
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
