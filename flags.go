package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Flags contains the flags parsed by the command line
type Flags struct {
	Config  string
	DataDir string
	Email   string

	DontSend  bool
	SleepBase time.Duration
	Test      bool
}

// init overwrites the usage
func init() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Usage:
	`+os.Args[0]+` [--data-dir <data_dir>] [--configs-file <configs_file>]
	               [--email-configs <email_file>]
	               [--dont-send] [--sleep <duration>] [--test]

Where:
* <data_dir> is the directory base directory from which <config_file> and the
  players' pictures paths relative paths are derived. Default: ".".
* <configs_file> is the path of the config file, relative to <data_dir>.
  Default: "configs.json".
* <email_file> is the path to the email config file. Default: "email.json".
* --dont-send, if set, will skip the email sending stage. Not set by default.
* --sleep can set a custom sleep time: --sleep 0 will be blazingly-fast!
  Default: "400ms". See time.ParseDuration for the format.
* --test, if set, will pair everyone with themself. Not set by default.`)
	}
}

// ParseFlags returns an instance of Flags and an error
func ParseFlags() (f Flags, err error) {
	flag.StringVar(&f.Config, "configs-file", "configs.json", "")
	flag.StringVar(&f.DataDir, "data-dir", ".", "")
	flag.StringVar(&f.Email, "email-configs", "email.json", "")

	flag.BoolVar(&f.DontSend, "dont-send", false, "")
	flag.DurationVar(&f.SleepBase, "sleep", 400*time.Millisecond, "")
	flag.BoolVar(&f.Test, "test", false, "")

	flag.Parse()

	return f, nil
}

// Sleep sleeps for some time
func (f Flags) Sleep(times int) {
	time.Sleep(f.SleepBase * time.Duration(times))
}
