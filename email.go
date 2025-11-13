package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"path/filepath"
)

//go:embed templates
var templates embed.FS

// EmailConfigs contains the informations needed to send the emails
type EmailConfigs struct {
	Address  string // Address is the sender's address
	Host     string // Host is the mail server's host
	Port     int    // Port is the mail server's port
	Login    string // Login is used to log into the mail server
	Password string // Password is used to log into the mail server
}

// LoadEmailConfigs returns an EmailConfigs instance loaded from the given file
func LoadEmailConfigs(filename string) (EmailConfigs, error) {
	var ec EmailConfigs

	if data, err := os.ReadFile(filename); err != nil {
		return ec, errors.New("could not open file")
	} else if err = json.Unmarshal([]byte(data), &ec); err != nil {
		return ec, errors.New("could not load file")
	}

	if ec.Address == "" {
		return ec, errors.New("missing required field \"address\"")
	} else if ec.Host == "" {
		return ec, errors.New("missing required field host")
	} else if ec.Port == 0 {
		return ec, errors.New("missing required field port")
	} else if ec.Login == "" {
		return ec, errors.New("missing required field login")
	} else if ec.Password == "" {
		return ec, errors.New("missing required field password")
	}

	return ec, nil
}

// SendMails generates and sends all the mails
func (ec EmailConfigs) SendMails(configs Configs, couples []Couple) error {
	pool := []*gomail.Message{}
	tmpl, _ := template.ParseFS(templates, "templates/base.html", "templates/"+configs.Lang+".html")

	for _, c := range couples {
		m := gomail.NewMessage()
		m.SetHeader("From", ec.Address)
		m.SetHeader("To", c.Giver.Email)
		m.SetHeader("Subject", configs.Subject)

		var buffer bytes.Buffer
		tmpl.Execute(&buffer, map[string]any{
			"GiverName":       c.Giver.Name,
			"GiverPicPath":    filepath.Base(c.Giver.PicPath),
			"Ideas":           c.Receiver.Ideas,
			"Notes":           configs.Notes,
			"ReceiverName":    c.Receiver.Name,
			"ReceiverPicPath": filepath.Base(c.Receiver.PicPath),
		})

		m.SetBody("text/html", buffer.String())
		m.Embed(c.Giver.PicPath)
		m.Embed("pics/_arrow.png")
		if c.Receiver.PicPath != c.Giver.PicPath {
			m.Embed(c.Receiver.PicPath)
		}

		pool = append(pool, m)
	}

	d := gomail.NewDialer(ec.Host, ec.Port, ec.Login, ec.Password)
	return d.DialAndSend(pool...)
}
