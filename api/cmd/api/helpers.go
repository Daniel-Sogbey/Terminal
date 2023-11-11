package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxByte := 1048576 // one megabyte of incomming data

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {

	out, err := json.Marshal(&data)

	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {

	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse

	payload.Error = true
	payload.Message = err.Error()

	app.writeJSON(w, statusCode, payload)

}
