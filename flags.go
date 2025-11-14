package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

// Flags contains the flags parsed by the command line
type Flags struct {
	Config string
	Email  string

	DontSend  bool
	SleepBase time.Duration
	Test      bool
}

// init overwrites the usage
func init() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Usage:
	`+os.Args[0]+` -config <config_file> --email <email_file> [--dont-send] [--sleep <duration>] [--test]

Where:
* <config_file> is the path to the config file
* <email_file> is the path to the email config file
* --dont-send, if set, will skip the email sending stage
* --sleep can set a custom sleep time; --sleep 0 will be blazingly-fast!
* --test, if set, will pair everyone with themself`)
	}
}

// ParseFlags returns an instance of Flags and an error
func ParseFlags() (f Flags, err error) {
	flag.StringVar(&f.Config, "config", "", "")
	flag.StringVar(&f.Email, "email", "", "")

	flag.BoolVar(&f.DontSend, "dont-send", false, "")
	flag.DurationVar(&f.SleepBase, "sleep", 400*time.Millisecond, "")
	flag.BoolVar(&f.Test, "test", false, "")

	flag.Parse()

	if f.Config == "" {
		return f, errors.New("Missing config file (--config)")
	} else if f.Email == "" {
		return f, errors.New("Missing email config file (--email)")
	}

	return f, nil
}

// Sleep sleeps for some time
func (f Flags) Sleep(times int) {
	time.Sleep(f.SleepBase * time.Duration(times))
}
