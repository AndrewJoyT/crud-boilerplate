package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"

	emailDomain "github.com/AndrewJoyT/crud-boilerplate/domain/email"
)

// SendEmail function for send email with mailgun
func SendEmail(recipients []string, subject string, htmlString string) error {
	SmtpHost := os.Getenv("SMTP_HOST")
	SmtpPort := os.Getenv("SMTP_PORT")
	SmtpUsername := os.Getenv("SMTP_USERNAME")
	SmtpPassword := os.Getenv("SMTP_PASSWORD")
	SmtpFrom := os.Getenv("SMTP_FROM")

	newEmail := &email.Email{
		To:      recipients,
		From:    SmtpFrom,
		Subject: subject,
		HTML:    []byte(htmlString),
	}
	fmt.Printf("%s:%s, %s, %s, %s, %s", SmtpHost, SmtpPort, SmtpUsername, SmtpPassword, SmtpHost, recipients)
	return newEmail.Send(fmt.Sprintf("%s:%s", SmtpHost, SmtpPort),
		smtp.PlainAuth("", SmtpUsername, SmtpPassword, SmtpHost),
	)
}

// CreateTemplate -
func CreateTemplate(input *emailDomain.TemplateEmailRequest) (string, error) {
	templateFile := fmt.Sprintf("utils/templates/%s", input.TemplateName)
	_, err := os.Stat(templateFile)
	if err == nil {
		htmlTemplate := template.Must(template.ParseFiles(templateFile))
		var buffer bytes.Buffer
		err := htmlTemplate.Execute(&buffer, input.Data)
		return buffer.String(), err
	}
	return "", err
}
