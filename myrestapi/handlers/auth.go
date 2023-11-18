package handlers

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/helpers"
	"github.com/Daniel-Sogbey/myrestapi/models"
	"github.com/Daniel-Sogbey/myrestapi/pkg"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := helpers.ReadJSON(r, &user)

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
			Error:  fmt.Sprintf("%v",err),
		}
		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
		return
	}

	user, err = user.GetUserByID(id)

	if err != nil {
		errorResponse := models.ErrorResponse{
			Status: "failed",
			Error:  "Internal server error. Try again in a minute",
		}
		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusInternalServerError)
		return
	}

	response := models.DataResponse{
		Status:  "success",
		Message: "Signed up successfully",
		Data:    &user,
	}

	//Send email to newly signed up user
	ec := pkg.NewEmailClient()
	ec.SendEmail("Message from email signup", user.Email)

	helpers.WriteJSON(w, response, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user *models.User

	err := helpers.ReadJSON(r, &user)

	if err != nil {
		errorResponse := &models.ErrorResponse{
			Status: "failed",
			Error:  "Internal server error. Try again in a minute",
		}

		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
		return
	}

	u, err := user.GetUserByEmailAndPassword(user.Password)

	if err != nil {
		errorResponse := &models.ErrorResponse{
			Status: "failed",
			Error:  "Incorrect user credentials",
		}

		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
		return
	}

	response := &models.DataResponse{
		Status:  "success",
		Message: "Logged in successfully",
		Data:    u,
	}

	helpers.WriteJSON(w, response, http.StatusOK)

}
