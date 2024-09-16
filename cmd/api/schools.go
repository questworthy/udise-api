package main

import (
	"net/http"

	"github.com/questworthy/udise-api/internal/data"
)

// Add a handler for "GET /v1/schools/:id" endpoint.
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the School struct, containing the ID (UDISECode)
	// we extracted from the URL and some dummy data.

	school := data.School{
		UDISECode:  id,
		SchoolName: "Dummy School",
	}

	// Encode the enveloped struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
