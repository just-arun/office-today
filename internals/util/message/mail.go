package message

import (
	"net/smtp"

	"github.com/just-arun/office-today/internals/boot/config"
)

// Mail for sending main
func Mail(emailTo string, message string) error {
	from := config.MailName
	pass := config.MailPass
	to := emailTo

	msg := message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("this is id", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
