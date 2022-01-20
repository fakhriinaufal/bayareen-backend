package service

import (
	"bayareen-backend/features/email"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/BurntSushi/toml"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type EmailService struct {
	Server   string
	Port     int
	Email    string
	Password string
}

func NewEmailService(path string) email.Service {
	es := EmailService{}
	if _, err := toml.DecodeFile(path, &es); err != nil {
		log.Fatal(err)
	}

	return &es
}

func (es *EmailService) Send(templateName string, r *email.Request, items interface{}) error {
	err := es.parseTemplate(templateName, r, items)
	if err != nil {
		log.Fatal(err)
	}
	if err = es.sendMail(r); err != nil {
		return err
	}
	return nil
}

func (es *EmailService) sendMail(r *email.Request) error {
	body := "To: " + r.To[0] + "\r\nSubject: " + r.Subject + "\r\n" + MIME + "\r\n" + r.Body
	SMTP := fmt.Sprintf("%s:%d", es.Server, es.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", es.Email, es.Password, es.Server), es.Email, r.To, []byte(body)); err != nil {
		return err
	}
	return nil
}

func (es *EmailService) parseTemplate(fileName string, r *email.Request, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.Body = buffer.String()
	return nil
}
