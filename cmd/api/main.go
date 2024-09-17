package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/questworthy/udise-api/internal/jsonlog"
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
	logger *jsonlog.Logger
}

func main() {

	var cfg config

	// Read port and env command-line flags into the config struct.
	// defaults to port 4000 & the environment "development"
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new jsonlog.Logger which writes any messages *at or above* the INFO
	// severity level to the standard out stream.
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// Declare an instance of the app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Use httprouter instance returned by app.routes() as server handler.
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		// Create a new Go log.Logger instance with the log.New() function, passing in
		// our custom Logger as the first parameter. The "" and 0 indicate that the
		// log.Logger instance should not use a prefix or any flags.
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Using the PrintInfo() method to write a "starting server" message at the
	// INFO level. But this time we pass a map containing additional properties (the
	// operating environment and server address) as the final parameter.
	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  cfg.env,
	})

	err := srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}
