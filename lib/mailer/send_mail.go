package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type ConfirmationBody struct {
	Token    string
	FullName string
}

func buildTemplateBody(subject string, templateFile string, templateBody any) []byte {
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+"\n%s\n\n", mimeHeaders)))

	t, _ := template.ParseFiles("./domain/core/templates/email/" + templateFile)
	er := t.Execute(&body, templateBody)
	if er != nil {
		log.Fatalln(er)
	}
	return body.Bytes()
}

func SendMail(email string, subject string, templateFile string, body any) {

	//auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	err := smtp.SendMail(
		os.Getenv("smtp_host")+":"+os.Getenv("SMTP_PORT"),
		nil,
		"vincentvad@gmail.com",
		[]string{email},
		buildTemplateBody(subject, templateFile, body),
	)

	if err != nil {
		log.Fatal(err)
	}
}
