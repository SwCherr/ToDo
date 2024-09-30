package service

import (
	"log"

	"net/smtp"
)

func (s *AuthService) sendEmail(email string) { // Connect to the remote SMTP server.
	from := "example@gmail.com"
	pass := "very_secret_pass"
	to := email
	msg := "There was an attempt to reissue a token from a different IP address."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
