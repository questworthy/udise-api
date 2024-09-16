package main

import (
	"fmt"
	"net/http"
)

// Add a handler for "GET /v1/schools/:id" endpoint.
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of school %d\n", id)

}
