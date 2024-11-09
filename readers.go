package main

import (
	"encoding/json"
	"os"
)

// Config contains the informations read from the config file
type Config struct {
	// Budget is shown at the end of the email, alongside currency
	Budget int

	// Currency is shown at the end of the email, alongside budget
	Currency string

	// IdeasEnabled is used to determine whether to show or not the
	// ideas box in the emails
	IdeasEnabled bool `json:"ideas_enabled"`

	// Email contains the creadentials used when sending emails
	Email struct {
		Address  string
		Host     string
		Port     int
		Login    string
		Password string
	}
}

// Runtime contains the read config
var Runtime Config

// ReadConfigs reads the config file and stores the result into Runtime
func ReadConfigs(path string) error {
	if data, err := os.ReadFile(path); err != nil {
		return err
	} else {
		return json.Unmarshal([]byte(data), &Runtime)
	}
}

// ReadPlayers reads the players file and creates them
func ReadPlayers(path string) ([]Player, error) {
	players := []Player{}

	if data, err := os.ReadFile(path); err != nil {
		return nil, err
	} else {
		return players, json.Unmarshal(data, &players)
	}
}
