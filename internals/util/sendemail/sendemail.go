package sendemail

import (
	"github.com/just-arun/office-today/internals/boot/config"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Payload for email
type Payload struct {
	FromName    string
	FromEmail   string
	ToName      string
	ToEmail     string
	Subject     string
	ContentHTML string
	ContentText string
}

// SendMail send mail
func (p *Payload) SendMail() (*rest.Response, error) {
	from := mail.NewEmail(p.FromName, p.FromEmail)
	subject := p.Subject
	to := mail.NewEmail(p.ToName, p.ToEmail)
	plainTextContent := p.ContentText
	htmlContent := p.ContentHTML
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendGridAPIKey)
	response, err := client.Send(message)
	return response, err
}
