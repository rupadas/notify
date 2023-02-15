package email

import (
	"log"
	"net/smtp"
	"os"

	common "github.com/rupadas/notify/config"
)

func Send(EmailBody common.EmailBody) {
	from := "***@gmail.com"
	pass := "***"
	to := EmailBody.To
	body := EmailBody.Body
	title := EmailBody.Title
	host := "smtp.gmail.com"
	toList := []string{to}

	// Its the default port of smtp server
	port := "587"

	msg := []byte(title + body)
	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, pass, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(host+":"+port, auth, from, toList, msg)

	// handling the errors
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("email sent successfully")
}
