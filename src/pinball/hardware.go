package main

import (
	"log"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

const (
	ADDRESS = 0x20
	IODIRA  = 0x00
	OLATA   = 0x14
	GPPUA   = 0x0C
)

type Hardware struct {
}

var table = Hardware{}

func (hardware *Hardware) init() {
	log.Println("Configuring hardware...")
	defer embd.CloseLED()
	for {
		embd.LEDToggle("LED0")
		time.Sleep(250 * time.Millisecond)
	}
}

func (hardware *Hardware) i2c() {
	bus := embd.NewI2CBus(1)
	defer bus.Close()
	bus.WriteByteToReg(ADDRESS, IODIRA, 0x80)
	bus.WriteByteToReg(ADDRESS, GPPUA, 0x80)
	var index byte = 0
	for ; index < 7; index++ {
		bus.WriteByteToReg(ADDRESS, OLATA, index)
		time.Sleep(250 * time.Millisecond)
	}

}

func (hardware *Hardware) direction(chip byte) {

}
