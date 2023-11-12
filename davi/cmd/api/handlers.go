package main

import "net/http"

func (a *application) GetAllHandler(w http.ResponseWriter, r *http.Request) {
users, err := a.

	if err != nil {
		a.errorLog.Panicln(err)
		return
	}

	a.infoLog.Println(users[0].ID)
}
