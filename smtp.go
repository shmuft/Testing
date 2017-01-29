package main

import (
	"net/smtp"
	"html/template"
	"bytes"
	"log"
)

var auth smtp.Auth

func sendEMail(u *UserForEmail) {
	auth = smtp.PlainAuth("", "Your@Email.com", "yourPassword", "smtp.gmail.com")


	r := NewRequest([]string{"Email@addressForWho"}, "Hello Username!", "Hello World!")
	err := r.ParseTemplate("./templates/emailTemplate.html", u)
	if err != nil {
		log.Println(err)
		return
	}
	ok, _ := r.Send()
	if ok {
		log.Println("All is OK!")
	}
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject string, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
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

func (r *Request) Send() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "Your@EmaiAddress.com", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}
