package message

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"

	"github.com/just-arun/office-today/internals/boot/config"
)

// Mail for sending main
func Mail(emailTo string, subject string, message string) error {
	m := gomail.NewMessage()

  // Set E-Mail sender
  m.SetHeader("From", config.MailName)

  // Set E-Mail receivers
  m.SetHeader("To", emailTo)

  // Set E-Mail subject
  m.SetHeader("Subject", subject)

  // Set E-Mail body. You can set plain text or html with text/html
  m.SetBody("text/plain", message)

  // Settings for SMTP server
  d := gomail.NewDialer("smtp.gmail.com", 587, config.MailName, config.MailPass)

  // This is only needed when SSL/TLS certificate is not valid on server.
  // In production this should be set to false.
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  // Now send E-Mail
  if err := d.DialAndSend(m); err != nil {
    fmt.Println(err)
    return err
  }

  return nil
}
