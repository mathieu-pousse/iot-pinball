package main

import (
	"log"
	"os/user"
)

func CheckSystem() {
	checkUID()
	log.Printf("All check passed")
}

func checkUID() {
	log.Printf("Checking user id...")
	var user, err = user.Current()
	if err != nil {
		log.Fatal("Can check if we are root")
	}
	if user.Uid != "0" {
		log.Fatalf("Must be root to interact properly with hardware (currently %s)", user.Name)
	}
}
