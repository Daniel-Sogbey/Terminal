package handlers

import (
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/models"
)

type AuthHandler struct {
	store models.UserStore
}

func NewAuthHandler(s models.UserStore) *AuthHandler {
	return &AuthHandler{
		store: s,
	}
}

func (h *AuthHandler) Signip(w http.ResponseWriter, r *http.Request) {
h.
}

// func Signup(w http.ResponseWriter, r *http.Request) {
// 	var user *models.User

// 	err := helpers.ReadJSON(r, &user)

// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Status: "falied",
// 			Error:  "Bad request from client",
// 		}

// 		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println(user)

// 	id, err := user.InsertUser()

// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Status: "failed",
// 			Error:  fmt.Sprintf("%v", err),
// 		}
// 		helpers.WriteErrorJSON(w, r, errorResponse, http.StatusBadRequest)
// 		return
// 	}

// 	response := models.DataResponse{
// 		Status:  "success",
// 		Message: "Signed up successfully",
// 		Data:    &id,
// 	}

// 	//Send email to newly signed up user
// 	ec := pkg.NewEmailClient()
// 	ec.SendEmail("Message from email signup", user.Email)

// 	helpers.WriteJSON(w, response, http.StatusOK)
// }
