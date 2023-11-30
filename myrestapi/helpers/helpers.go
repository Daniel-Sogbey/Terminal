package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, payload interface{}, statusCode int, headers ...http.Header) error {

	b, err := json.Marshal(&payload)

	if err != nil {

		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)

	return nil
}

func WriteErrorJSON(w http.ResponseWriter, r *http.Request, payload interface{}, statusCode int, headers ...http.Header) error {

	b, err := json.Marshal(&payload)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)

	return nil
}

func ReadJSON(r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return err
	}

	fmt.Println("**", data)

	return nil

}
