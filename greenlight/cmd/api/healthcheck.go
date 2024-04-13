package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	headers := http.Header{}

	headers.Add("header1", "value1")
	headers.Add("header2", "value2")
	headers.Add("header3", "value3")

	err := app.writeJSON(w, http.StatusOK, env, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
