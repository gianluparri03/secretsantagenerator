package main

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"path/filepath"
)

// messagePool contains all the messages to be sent
var messagePool []*gomail.Message

// CreateMail creates the message for the giver, without sending it
func CreateMail(c Couple) {
	m := gomail.NewMessage()
	m.SetHeader("From", EmailAddress)
	m.SetHeader("To", c.Giver.Email)
	m.SetHeader("Subject", "Secret Santa Generator")

	var buffer bytes.Buffer
	tmpl, _ := template.ParseFiles("templates/base.html", "templates/"+c.Giver.Language+".html")
	tmpl.Execute(&buffer, map[string]any{
		"GiverName":       c.Giver.Name,
		"GiverPicPath":    filepath.Base(c.Giver.PicPath),
		"ReceiverName":    c.Receiver.Name,
		"ReceiverPicPath": filepath.Base(c.Receiver.PicPath),
		"IdeasEnabled":    IdeasEnabled,
		"Ideas":           c.Receiver.Ideas,
		"Budget":          Budget,
		"Currency":        Currency,
	})

	m.SetBody("text/html", buffer.String())
	m.Embed(c.Giver.PicPath)
	m.Embed("pics/_arrow.png")
	m.Embed(c.Receiver.PicPath)

	messagePool = append(messagePool, m)
}

// SendAll sends all the messages in the pool
func SendAll() {
	d := gomail.NewDialer(EmailHost, EmailPort, EmailLogin, EmailPassword)
	if err := d.DialAndSend(messagePool...); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
