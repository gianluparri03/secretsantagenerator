package main

import (
	"embed"
	"fmt"
	"os"
)

//go:embed templates pics/_missing.png pics/_arrow.png
var assets embed.FS

func main() {
	flags, err := ParseFlags()
	checkErr(err, false)

	fmt.Println("Secret Santa Generator")
	fmt.Println("======================")
	if flags.DontSend && flags.Test {
		fmt.Println("Running in dont-send + test mode\n")
	} else if flags.DontSend {
		fmt.Println("Running in dont-send mode\n")
	} else if flags.Test {
		fmt.Println("Running in test mode\n")
	}
	flags.Sleep(2)

	fmt.Print("Reading email_file... ")
	emailconfigs, err := LoadEmailConfigs(flags.Email)
	flags.Sleep(1)
	checkErr(err, true)

	fmt.Print("Reading configs_file... ")
	configs, err := LoadConfigs(flags.Config)
	flags.Sleep(1)
	checkErr(err, true)

	flags.Sleep(2)

	fmt.Print("\nPlayers found:\n")
	for _, p := range configs.Players {
		flags.Sleep(1)
		fmt.Printf("|  %s <%s>\n", p.Name, p.Email)
	}

	fmt.Print("\n")
	flags.Sleep(2)

	fmt.Print("Generating couples... ")
	var couples []Couple
	if !flags.Test {
		couples = configs.GenerateCouples()
	} else {
		couples = configs.GenerateTestCouples()
	}
	flags.Sleep(4)
	checkErr(nil, true)

	fmt.Print("Building emails... ")
	mails := emailconfigs.BuildMails(configs, couples)
	flags.Sleep(3)
	checkErr(nil, true)

	fmt.Print("Connecting to the mail server... ")
	err = emailconfigs.TryConnect()
	flags.Sleep(1)
	checkErr(err, true)

	if !flags.DontSend {
		fmt.Print("Sending emails... ")
		err = emailconfigs.SendMails(mails)
		flags.Sleep(1)
		checkErr(err, true)
	}

	flags.Sleep(1)
	fmt.Print("\nDone! Merry Christmas everybody!\n")
}

// checkErr prints the error and exits if there's an error,
// otherwise if printOk is set it will print "ok.\n"
func checkErr(err error, printOk bool) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else if printOk {
		fmt.Println("ok.")
	}
}
