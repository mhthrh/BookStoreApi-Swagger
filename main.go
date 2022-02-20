package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/View"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "the TCP address for the server to listen on, in the form 'host:port'")

	l := log.New(os.Stdout, "BookStore-api ", log.LstdFlags)
	sm := mux.NewRouter()
	View.RunApiOnRouter(sm, l)
	// create a new server
	s := http.Server{
		Addr:         *addr,             // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 20 * time.Second,  // max time to write response to the client
		IdleTimeout:  180 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on  %s", *addr)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
