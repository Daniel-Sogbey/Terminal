package handlers

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/helpers"
	"github.com/Daniel-Sogbey/myrestapi/models"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := helpers.ReadJSON(r, user)

	if err != nil {
		errorResponse := models.ErrorResponse{
			Status: "falied",
			Error:  "Bad request from client",
		}

		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	id, err := user.InsertUser()

	if err != nil {
		errorResponse := models.ErrorResponse{
			Status: "failed",
			Error:  fmt.Sprintf("%v", err),
		}
		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusInternalServerError)
		return
	}

	user, err = user.GetUserByID(id)

	if err != nil {
		errorResponse := models.ErrorResponse{
			Status: "failed",
			Error:  fmt.Sprintf("%v", err),
		}
		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusInternalServerError)
		return
	}

	response := models.DataResponse{
		Status:  "success",
		Message: "Signed up successfully",
		Data:    &user,
	}

	helpers.WriteJSON(w, response, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user *models.User

	err := helpers.ReadJSON(r, &user)

	if err != nil {
		errorResponse := &models.ErrorResponse{
			Status: "failed",
			Error:  fmt.Sprintf("%v", err),
		}

		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)

	}

	u, err := user.GetUserByUsernameAndPassword(user.Password)

	if err != nil {
		errorResponse := &models.ErrorResponse{
			Status: "failed",
			Error:  fmt.Sprintf("%v", err),
		}

		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
	}

	response := &models.DataResponse{
		Status:  "success",
		Message: "Logged in successfully",
		Data:    u,
	}

	helpers.WriteJSON(w, response, http.StatusOK)

}
