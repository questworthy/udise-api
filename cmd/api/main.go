package main

import (
	"context"
	"flag"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/questworthy/udise-api/internal/jsonlog"
)

// application version number ~ [TODO] automatically generate this at build time.
const version = "1.0.0"

// read in these from command-line flags when the app starts.
type config struct {
	port int
	env  string
	// Adds a new limiter struct containing fields for the requests-per-second and burst
	// values, and a boolean field which we can use to enable/disable rate limiting
	// altogether.
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

// app struct to hold the dependencies for our HTTP handlers, helpers & middleware.
type application struct {
	config config
	logger *jsonlog.Logger
	client *bigquery.Client
	ctx    context.Context
}

func main() {

	var cfg config

	// Read port and env command-line flags into the config struct.
	// defaults to port 8080 & the environment "development"
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Create command line flags to read the setting values into the config struct.
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	// Initialize a new jsonlog.Logger which writes any messages *at or above* the INFO
	// severity level to the standard out stream.
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// Initialize BigQuery client
	ctx := context.Background()
	bqClient, err := bigquery.NewClient(ctx, "afe-bot")
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	// Declare an instance of the app struct
	app := &application{
		config: cfg,
		logger: logger,
		client: bqClient,
		ctx:    ctx,
	}

	// Call app.serve() to start the server.
	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
