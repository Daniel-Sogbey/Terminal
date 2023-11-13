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
		log.Println(err)
		errorResponse := models.ErrorResponse{
			Status: "failed",
			Error:  fmt.Sprintf("%v", err),
		}

		b, _ := json.Marshal(&errorResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}

	user, err = user.GetUserByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// er := &pkg.EmailRequest{
	// 	Personalization: []pkg.Personalization{
	// 		{
	// 			To: []pkg.Email{
	// 				{
	// 					Email: "mathematics06physics@gmail.com",
	// 					Name:  "Daniel",
	// 				},
	// 			},
	// 			Subject: "Hello, World",
	// 		},
	// 	},
	// 	Content: []pkg.Content{
	// 		{
	// 			Type:  "text/plain",
	// 			Value: "Hello, From Go Server",
	// 		},
	// 	},
	// 	From: pkg.From{
	// 		Email: "gift.alchemy.developer@gmail.com",
	// 		Name:  "Daniel",
	// 	},
	// 	ReplyTo: pkg.ReplyTo{
	// 		Email: "gift.alchemy.developer@gmail.com",
	// 		Name:  "Daniel",
	// 	},
	// }

	// err = pkg.SendEmail(os.Getenv("SEND_GRID_API_KEY"), er)

	if err != nil {
		return
	}

	response := models.DataResponse{
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
