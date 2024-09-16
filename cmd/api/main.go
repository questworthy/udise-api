package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// application version number ~ [TODO] automatically generate this at build time.
const version = "1.0.0"

// read in these from command-line flags when the app starts.
type config struct {
	port int
	env  string
}

// app struct to hold the dependencies for our HTTP handlers, helpers & middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	// Read port and env command-line flags into the config struct.
	// defaults to port 4000 & the environment "development"
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Use httprouter instance returned by app.routes() as server handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
