package main

import (
	"errors"
	"net/http"

	"github.com/questworthy/udise-api/internal/data"
)

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r)
		return
	}

	// Call the Get() method to fetch the data for a specific school
	school, err := data.Get(id, app.client, app.ctx)
	if err != nil {
		switch {

		case errors.Is(err, data.ErrInvalidUDISE):
			app.badRequestResponse(w, r)

		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Encode the enveloped struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
