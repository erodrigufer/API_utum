package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/company", app.addCompanyDataHandler)
	router.HandlerFunc(http.MethodGet, "/v1/company/:id", app.showCompanyHandler)
	// Return the httprouter instance.
	return router
}
