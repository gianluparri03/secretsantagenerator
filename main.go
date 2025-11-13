package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Print("Secret Santa Generator\n======================\n")

	flags, err := ParseFlags()
	if err != nil {
		fmt.Printf("%s.\nExiting.\n", err.Error())
		os.Exit(1)
	} else {
		time.Sleep(flags.Sleep)
	}

	fmt.Print("Reading email_file... ")
	emailconfigs, err := LoadEmailConfigs(flags.Email)
	if err != nil {
		fmt.Printf("%s.\nExiting.\n", err.Error())
		os.Exit(1)
	} else {
		time.Sleep(flags.Sleep)
		fmt.Print("ok.\n")
	}

	fmt.Print("Reading configs_file... ")
	configs, err := LoadConfigs(flags.Config)
	if err != nil {
		fmt.Printf("%s.\nExiting.\n", err.Error())
		os.Exit(1)
	} else {
		time.Sleep(flags.Sleep)
		fmt.Print("ok.\n")
	}

	time.Sleep(2 * flags.Sleep)

	fmt.Print("\nPlayers found:\n")
	for _, p := range configs.Players {
		time.Sleep(flags.Sleep)
		fmt.Print("|  " + p.Name + "\n")
	}
	fmt.Print("\n")

	time.Sleep(2 * flags.Sleep)

	if !flags.Parse {
		fmt.Print("Generating couples... ")
		var couples []Couple
		if !flags.Test {
			couples = configs.GenerateCouples()
		} else {
			couples = configs.GenerateTestCouples()
		}
		time.Sleep(5 * flags.Sleep)
		fmt.Print("ok.\n\n")

		fmt.Print("Sending emails... ")
		if err := emailconfigs.SendMails(couples, configs.Subject, configs.Lang); err != nil {
			fmt.Printf("%s.\nExiting.\n", err.Error())
			os.Exit(1)
		} else {
			fmt.Print("ok.\n")
		}
	}

	fmt.Print("\nDone! Merry Christmas everybody!\n")
}
