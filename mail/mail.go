package mail

import (
	"log"
	"net/smtp"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

func SendMail(to []string, msg string) {
	auth := smtp.PlainAuth("", "khanhvn@eupgroup.net", "rwjyecotmmaohgud", smtpAuthAddress)
	// Here we do it all: connect to our server, set up a message and send it
	err := smtp.SendMail(smtpServerAddress, auth, "abcxyz@gmail.com", to, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

}
