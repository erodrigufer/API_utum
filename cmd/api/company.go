package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/erodrigufer/API_utum/internal/data"
)

func (app *application) addCompanyDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "add new company data")
}
func (app *application) showCompanyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	company := data.Company{
		ID:            id,
		CreatedAt:     time.Now(),
		Name:          "Siemens",
		Year:          1905,
		Valuation:     1000000,
		Headcount:     50000,
		IndustryField: []string{"energy", "industrial automation", "mobility"},
		Timestamp:     time.Now(),
		Version:       1,
	}
	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, company, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
