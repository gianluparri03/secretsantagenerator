package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"gopkg.in/gomail.v2"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

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

// BuildMails builds and returns all the mails
func (ec EmailConfigs) BuildMails(configs Configs, couples []Couple) (pool []*gomail.Message) {
	tmpl, _ := template.ParseFS(assets, "templates/base.html", "templates/"+configs.Lang+".html")

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
		embedFile(m, c.Giver.PicPath)
		embedFile(m, "pics/_arrow.png")
		if c.Receiver.PicPath != c.Giver.PicPath {
			embedFile(m, c.Receiver.PicPath)
		}

		pool = append(pool, m)
	}

	return pool
}

// SendMails sends the mails
func (ec EmailConfigs) SendMails(pool []*gomail.Message) error {
	dialer := gomail.NewDialer(ec.Host, ec.Port, ec.Login, ec.Password)
	err := dialer.DialAndSend(pool...)
	return err
}

// TryConnect tries to connect to the mail server
func (ec EmailConfigs) TryConnect() error {
	dialer := gomail.NewDialer(ec.Host, ec.Port, ec.Login, ec.Password)
	sender, err := dialer.Dial()
	if err == nil {
		sender.Close()
	}
	return err
}

// embedFile embeds a file into a message; it firstly tries to get the file
// from the embedded assets, and if does not find it it fetches it from the
// file system.
func embedFile(msg *gomail.Message, path string) {
	msg.Embed(path, gomail.SetCopyFunc(func(w io.Writer) error {
		bytes, err := assets.ReadFile(path)
		if err != nil {
			bytes, err = os.ReadFile(path)
			if err != nil {
				return err
			}
		}

		_, err = w.Write(bytes)
		return err
	}))
}
