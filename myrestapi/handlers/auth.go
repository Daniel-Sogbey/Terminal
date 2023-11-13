package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/models"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	id, err := user.InsertUser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err = user.GetUserByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "Signed up successfully",
		Data:    &user,
	}

	b, err := json.Marshal(&response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
