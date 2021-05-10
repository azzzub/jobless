package utils

import (
	"bytes"
	"errors"
	"html/template"
	"net/smtp"
)

type templateData struct {
	Name string
	URL  string
}

type request struct {
	from    string
	to      []string
	subject string
	body    string
}

var auth smtp.Auth

func SendMail(receiverMail string, subject string, body string) error {
	auth = smtp.PlainAuth(
		GetEnv("MAIL_ID", ""),
		GetEnv("MAIL_USER", "your_mail_address@mail.com"),
		GetEnv("MAIL_PASS", "your_mail_password"),
		GetEnv("MAIL_HOST", "your_mail_host"),
	)

	templateData := &templateData{
		Name: "Jobless",
		URL:  "http://todo.todo",
	}

	r := newRequest([]string{receiverMail}, subject, body)
	if err := r.ParseTemplate("emailVerification.html", templateData); err == nil {
		ok, err := r.send()
		if !ok {
			return errors.New("sending email failed")
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func newRequest(to []string, subject, body string) *request {
	return &request{
		from:    "juanaayubs@gmail.com",
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *request) send() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, r.from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}
