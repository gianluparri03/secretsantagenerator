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
	Lang    string
	Players []Player
	Subject string
	Notes   string
}

// LoadConfigs returns an EmailConfigs instance loaded from the given file
func LoadConfigs(filename string) (Configs, error) {
	var c Configs

	if data, err := os.ReadFile(filename); err != nil {
		return c, errors.New("could not open file")
	} else if err = json.Unmarshal([]byte(data), &c); err != nil {
		return c, errors.New("could not load file")
	}

	if c.Lang == "" {
		c.Lang = DEFAULT_LANG
	}
	if c.Subject == "" {
		c.Subject = DEFAULT_SUBJECT
	}

	if f, err := templates.Open("templates/" + c.Lang + ".html"); err == nil {
		f.Close()
	} else if err != nil {
		return c, errors.New("unknown lang")
	} else if len(c.Players) < 2 {
		return c, errors.New("too few players")
	}

	for p, _ := range c.Players {
		if c.Players[p].Name == "" {
			return c, errors.New("found a player without a name")
		} else if c.Players[p].Email == "" {
			return c, errors.New("found a player without an email")
		}

		if c.Players[p].PicPath == "" {
			c.Players[p].PicPath = DEFAULT_PICPATH
		}

		if _, err := os.Stat(c.Players[p].PicPath); err != nil {
			return c, errors.New("pic not found")
		}

		for _, i := range c.Players[p].Ideas {
			if i.Name == "" {
				return c, errors.New("found an idea without a name")
			}
		}

		for p2 := p + 1; p2 < len(c.Players); p2++ {
			if c.Players[p2].Name == c.Players[p].Name {
				return c, errors.New("found a duplicated name")
			} else if c.Players[p2].Email == c.Players[p].Email {
				return c, errors.New("found a duplicated email")
			}
		}
	}

	return c, nil
}
