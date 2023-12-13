package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	var (
		username   = "daniel.sogbey@hubtel.com"
		smtpServer = "smtp.office365.com"
		// smtpPort   = 587
		password = "#12maN@gh"
		identity = ""
	)

	auth := smtp.PlainAuth(identity, username, password, smtpServer)

	fmt.Println(auth)

	err := smtp.SendMail(smtpServer+":587", auth, username, []string{"sogbeydaniel28@gmail.com"}, []byte("Hello,World!"))

	if err != nil {
		log.Fatal(err)
	}

}
