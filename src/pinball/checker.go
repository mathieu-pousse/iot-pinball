package main

import (
	"log"
	"os/exec"
	"strings"
)

func CheckSystem() {
	checkUID()
	log.Printf("All check passed")
}

func checkUID() {
	log.Printf("Checking user id...")
	cmd := exec.Command("id", "-un")
	output, err := cmd.Output()

	if err != nil {
		log.Fatal("Cannot determine user id, assume not root", err)
	}
	if strings.TrimSpace(string(output)) != "root" {
		log.Fatalf("not running as root but as %s", output)
	}
}
