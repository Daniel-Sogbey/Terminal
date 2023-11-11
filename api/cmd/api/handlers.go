package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Username string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		//send back an error message
		app.errorLog.Println("invalid json", err)

		payload.Error = true
		payload.Message = "Invalid request body"

		out, err := json.Marshal(&payload)

		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(out)
		return

	}

	//TODO: authenticate
	app.infoLog.Printf("%+v", creds)

	//TODO: send back response
	payload.Error = false
	payload.Message = "Signed in"

	out, err := json.Marshal(&payload)

	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
