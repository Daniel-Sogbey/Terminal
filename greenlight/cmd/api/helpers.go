package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (app *application) readIDParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encounted a problem and could not process your request", http.StatusInternalServerError)
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	app.logger.Println("headers : ", headers)

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.WriteHeader(status)

	_, err = w.Write(js)

	if err != nil {
		app.logger.Println(err)
		return err
	}

	return nil

}
