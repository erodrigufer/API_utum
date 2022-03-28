package main

import (
	"fmt"
	"net/http"
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

	fmt.Fprintf(w, "show the details of company %d\n", id)
}
