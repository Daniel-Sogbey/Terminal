package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
}

type UserMetal struct {
	Visits int
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
