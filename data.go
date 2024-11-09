package main

import (
	"errors"
	"math/rand"
	"net/mail"
	"os"
)

// Idea is an idea given by the gift-receiver,
// with a name and an optional link
type Idea struct {
	Description string
	Link        string
}

// Player is both a gift-giver and a gift-receiver.
// It has a name, an email, a language, a list of ideas
// and an optional pic path
type Player struct {
	Name     string
	Email    string
	PicPath  string `json:"pic_path"`
	Language string
	Ideas    []Idea
}

// Players contains all the players
var Players []Player

// Validate check if a player is valid;
// if not, returns also the reason
func (p *Player) Save() error {
	// Restore default picpath if removed
	if p.PicPath == "" {
		p.PicPath = "pics/_missing.png"
	}

	// Check attributes length
	if len(p.Name) < 1 {
		return errors.New("Name must be set")
	} else if _, err := mail.ParseAddress(p.Email); err != nil {
		return errors.New("Invalid email")
	} else if _, err := os.Stat(p.PicPath); err != nil {
		return errors.New("Invalid picture")
	} else if p.Language != "en" && p.Language != "it" {
		return errors.New("Unknown language")
	}

	// Checks if data is duplicated
	for _, p2 := range Players {
		if p.Name == p2.Name {
			return errors.New("Name duplicated")
		} else if p.Email == p2.Email {
			return errors.New("Email duplicated")
		}
	}

	// Adds missing data
	for _, i := range p.Ideas {
		if len(i.Description) < 1 && len(i.Link) > 1 {
			i.Description = "[link]"
		}
	}

	// Adds it to the list
	Players = append(Players, *p)
	return nil
}

// Couple stores the relation giver-receiver
type Couple struct {
	Giver    Player
	Receiver Player
}

// generateDerangement generates a derangement
func generateDerangement(n int) (d []int) {
	isValid := false
	for !isValid {
		d = rand.Perm(n)

		isValid = true
		for a, b := range d {
			if a == b {
				isValid = false
				break
			}
		}
	}

	return d
}

// GenerateCouples creates all the couples.
func GenerateCouples() []Couple {
	var couples []Couple
	for g, r := range generateDerangement(len(Players)) {
		couples = append(couples, Couple{Giver: Players[g], Receiver: Players[r]})
	}

	return couples
}
