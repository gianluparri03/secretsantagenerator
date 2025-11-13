package main

import (
	"errors"
	"flag"
	"time"
)

type Flags struct {
	Config string
	Email  string

	Parse bool
	Sleep time.Duration
	Test  bool
}

func ParseFlags() (f Flags, err error) {
	flag.StringVar(&f.Config, "config", "", "The config file")
	flag.StringVar(&f.Email, "email", "", "The email config file")

	flag.BoolVar(&f.Parse, "parse", false, "If set, the emails aren't sent")
	flag.DurationVar(&f.Sleep, "sleep", 400*time.Millisecond, "The amount of time slept between operations")
	flag.BoolVar(&f.Test, "test", false, "If set, everyone will be paired with themself")

	flag.Parse()

	if f.Config == "" {
		return f, errors.New("Missing config file (--config)")
	} else if f.Email == "" {
		return f, errors.New("Missing email config file (--email)")
	}

	return f, nil
}
