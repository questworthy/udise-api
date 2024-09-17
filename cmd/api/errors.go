package main

import (
	"fmt"
	"net/http"
)

// Generic helper for logging an error message.
func (app *application) logError(r *http.Request, err error) {

	// Use the PrintError() method to log the error message, and include the current
	// request method and URL as properties in the log entry.
	app.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

// Generic helper for sending JSON-formatted error messages to the client with a
// given status code. Using an interface{} type for the message parameter, rather
// than just a string type for more flexibility over the values that we can
// include in the response.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)

	if err != nil {

		// Log errors (if any), and fall back to sending the client an
		// empty response with a 500 Internal Server Error status code.
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// Logs the detailed error message, then uses the errorResponse() helper to
// send a 500 Internal Server Error status code and JSON response
// (containing a generic error message) to the client.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Sends a 404 Not Found status code and JSON response to the client.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Sends a 405 Method Not Allowed status code and JSON response to the client.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
