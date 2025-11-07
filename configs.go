package main

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	DEFAULT_SUBJECT = "SecretSantaGenerator"
	DEFAULT_LANG    = "en"
	DEFAULT_PICPATH = "pics/_missing.png"
)

// Configs contains the informations read from the config file
type Configs struct {
	Subject string // optional
	Lang    string // optional
	Players []Player
}

// LoadConfigs returns an EmailConfigs instance loaded from the given file
func LoadConfigs(filename string) (Configs, error) {
	var c Configs

	if data, err := os.ReadFile(filename); err != nil {
		return c, errors.New("could not open file")
	} else if err = json.Unmarshal([]byte(data), &c); err != nil {
		return c, errors.New("could not load file")
	}

	if c.Subject == "" {
		c.Subject = DEFAULT_SUBJECT
	}
	if c.Lang == "" {
		c.Lang = DEFAULT_LANG
	}

	if _, err := os.Stat("templates/" + c.Lang + ".html"); err != nil {
		return c, errors.New("unknown lang")
	} else if len(c.Players) < 2 {
		return c, errors.New("too few players")
	}

	for j, p := range c.Players {
		if p.Name == "" {
			return c, errors.New("found a player without a name")
		} else if p.Email == "" {
			return c, errors.New("found a player without an email")
		}

		if p.PicPath == "" {
			c.Players[j].PicPath = DEFAULT_PICPATH
		}

		for _, i := range p.Ideas {
			if i.Description == "" {
				return c, errors.New("found an idea without a description")
			}
		}

		for jj := j + 1; jj < len(c.Players); jj++ {
			if c.Players[jj].Name == p.Name {
				return c, errors.New("found a duplicated name")
			} else if c.Players[jj].Email == p.Email {
				return c, errors.New("found a duplicated email")
			}
		}
	}

	return c, nil
}
