package main

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"path/filepath"
)

// SendMails generates and sends all the mails
func SendMails(couples []Couple) error {
	pool := []*gomail.Message{}
	for _, c := range couples {
		pool = append(pool, createMail(c))
	}

	return sendAll(pool)
}

// createMail creates the message for the giver
func createMail(c Couple) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", Runtime.Email.Address)
	m.SetHeader("To", c.Giver.Email)
	m.SetHeader("Subject", Runtime.Subject)

	var buffer bytes.Buffer
	tmpl, _ := template.ParseFiles("templates/base.html", "templates/"+Runtime.Template+".html")
	tmpl.Execute(&buffer, map[string]any{
		"GiverName":       c.Giver.Name,
		"GiverPicPath":    filepath.Base(c.Giver.PicPath),
		"ReceiverName":    c.Receiver.Name,
		"ReceiverPicPath": filepath.Base(c.Receiver.PicPath),
		"IdeasEnabled":    Runtime.IdeasEnabled,
		"Ideas":           c.Receiver.Ideas,
		"Budget":          Runtime.Budget,
		"Currency":        Runtime.Currency,
	})

	m.SetBody("text/html", buffer.String())
	m.Embed(c.Giver.PicPath)
	m.Embed("pics/_arrow.png")
	if c.Receiver.PicPath != c.Giver.PicPath {
		m.Embed(c.Receiver.PicPath)
	}

	return m
}

// sendAll sends all the messages in the pool
func sendAll(messages []*gomail.Message) error {
	d := gomail.NewDialer(Runtime.Email.Host, Runtime.Email.Port, Runtime.Email.Login, Runtime.Email.Password)
	return d.DialAndSend(messages...)
}
