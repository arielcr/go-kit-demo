package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/arielcr/go-kit-demo/greeter"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()

	ctx := context.Background()

	// the greeter service
	srv := greeter.NewService()

	errChan := make(chan error)

	// go routine to stop server when we hit ctrl+c
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := greeter.Endpoints{
		StatusEndpoint: greeter.MakeStatusEndpoint(srv),
		HelloEndpoint:  greeter.MakeHelloEndpoint(srv),
	}

	// go routine to listen for incoming requests
	go func() {
		log.Println("greeter is listening on port: ", *httpAddr)
		handler := greeter.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)

}
