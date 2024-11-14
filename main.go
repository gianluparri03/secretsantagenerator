package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide a config and a players file")
	}

	log.Println("Secret Santa Generator")
	log.Println("======================")

	var err error
	log.Println("Reading configs...")
	if err = ReadConfigs(os.Args[1]); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	var rp []Player
	log.Println("Reading players...")
	if rp, err = ReadPlayers(os.Args[2]); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	if len(rp) < 2 {
		log.Fatalf("Error: too few players")
	}

	log.Println("Parsing players...")
	for _, p := range rp {
		if err = p.Save(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		} else {
			log.Printf("|  %s okay", p.Name)
		}
	}
	log.Println("done")

	log.Println("Generating couples...")
	couples := GenerateCouples()

	log.Println("Sending emails...")
	if err := SendMails(couples); err != nil {
		log.Fatalf("Error: %s", err.Error())
	} else {
		log.Println("All done! Enjoy!")
	}
}
