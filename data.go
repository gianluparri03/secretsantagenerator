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
	PicPath  string
	Language string
	Ideas    []*Idea
}

var (
	// Players contains all the players
	Players []*Player

	// Default values for the players
	defaultPicPath string = "pics/_missing.png"
	defaultLang    string = "en"
)

// NewPlayer creates a new player and adds it to the list
func NewPlayer() *Player {
	p := Player{
		Language: defaultLang,
		Ideas:    []*Idea{},
	}

	Players = append(Players, &p)
	return &p
}

// Validate check if a player is valid;
// if not, returns also the reason
func (p *Player) Validate() (bool, string) {
	// Restore default picpath if removed
	if p.PicPath == "" {
		p.PicPath = defaultPicPath
	}

	// Check attributes length
	if len(p.Name) < 1 {
		return false, "Name must be set"
	} else if _, err := mail.ParseAddress(p.Email); err != nil {
		return false, "Invalid email"
	} else if _, err := os.Stat(p.PicPath); err != nil {
		return false, "Invalid picture"
	} else if p.Language != "en" && p.Language != "it" {
		return false, "Unknown language"
	}

	// Checks if data is duplicated
	for _, p2 := range Players {
		if p != p2 {
			if p.Name == p2.Name {
				return false, "Name duplicated"
			} else if p.Email == p2.Email {
				return false, "Email duplicated"
			}
		}
	}

	// Adds missing data
	for _, i := range p.Ideas {
		if len(i.Description) < 1 && len(i.Link) > 1 {
			i.Description = "[link]"
		}
	}

	return true, ""
}

// Couple stores the relation giver-receiver
type Couple struct {
	Giver    *Player
	Receiver *Player
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
// It also makes sure all the players are valid
func GenerateCouples() ([]Couple, error) {
	for _, p := range Players {
		if isValid, reason := p.Validate(); !isValid {
			return nil, errors.New(p.Name + " is not valid (" + reason + ")")
		}
	}

	var couples []Couple
	for g, r := range generateDerangement(len(Players)) {
		couples = append(couples, Couple{Giver: Players[g], Receiver: Players[r]})
	}

	return couples, nil
}
