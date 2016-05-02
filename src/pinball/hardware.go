package main

import (
	"log"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

type Hardware struct {
}

var table = Hardware{}

func (hardware *Hardware) init() {
	log.Println("Configuring hardware...")
	for {
		embd.LEDToggle("LED0")
		time.Sleep(250 * time.Millisecond)
	}

}
