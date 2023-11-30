package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) WriteJSON(w http.ResponseWriter, data interface{}, status int) error {
	js, err := json.Marshal(&data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
