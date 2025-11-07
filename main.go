package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	print("Secret Santa Generator\n======================\n")

	if len(os.Args) < 3 {
		fatal("Please provide an email_file and a config_file")
	}

	print("Reading email_file... ")
	emailconfigs, err := LoadEmailConfigs(os.Args[1])
	if err != nil {
		fatal(err.Error())
	} else {
		print("ok.\n")
	}

	print("Reading configs_file... ")
	configs, err := LoadConfigs(os.Args[2])
	if err != nil {
		fatal(err.Error())
	} else {
		print("ok.\n")
	}

	print("\nPlayers found:\n")
	for _, p := range configs.Players {
		print("|  " + p.Name + "\n")
	}
	print("\n")

	print("Generating couples... ")
	couples := configs.GenerateCouples()
	print("ok.\n")
	print("\n")

	print("Sending emails... ")
	if err := emailconfigs.SendMails(couples, configs.Subject, configs.Lang); err != nil {
		fatal(err.Error())
	} else {
		print("ok.\n")
	}

	print("\nDone! Bye bye!\n")
}

// fatal prints an error then exits
func fatal(msg string) {
	fmt.Printf("%s.\nExiting.\n", msg)
	os.Exit(1)
}

// print prints a message
func print(msg string) {
	fmt.Print(msg)
	time.Sleep(400 * time.Millisecond) // just for fun
}
