package main

import "math/rand"

// Player is both a giver and a receiver
type Player struct {
	Name    string
	Email   string
	PicPath string `json:"pic_path"` // optional
	Ideas   []Idea // optional
}

// Idea is an idea proposed by the receiver to the giver
type Idea struct {
	Description string
	Link        string // optional
}

// Couple stores the relation giver-receiver
type Couple struct {
	Giver    Player
	Receiver Player
}

// GenerateCouples creates all the couples
func (c Configs) GenerateCouples() []Couple {
	var derangement []int
	for valid := false; !valid; {
		valid = true
		derangement = rand.Perm(len(c.Players))
		for a, b := range derangement {
			if a == b {
				valid = false
				break
			}
		}
	}

	var couples []Couple
	for g, r := range derangement {
		couples = append(couples, Couple{Giver: c.Players[g], Receiver: c.Players[r]})
	}

	return couples
}
