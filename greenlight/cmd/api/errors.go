package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(w http.ResponseWriter, status int, message interface{}) {

	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, err error) {
	app.logError(err)

	message := "the server encounted a problem and could not process your request"

	app.errorResponse(w, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resourse could not be found"

	app.errorResponse(w, http.StatusNotFound, message)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resourse", r.Method)

	app.errorResponse(w, http.StatusMethodNotAllowed, message)
}
