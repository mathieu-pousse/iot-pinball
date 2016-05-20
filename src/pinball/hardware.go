package main

import (
	"log"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

const (
	ADDRESS  = 0x20 // base chip address
	IODIRA   = 0x00 // set IO direction (1 = Input / 0 = Output)
	IODIRB   = 0x01 // set IO direction (1 = Input / 0 = Output)
	GPINTENA = 0x04 // enable interruption
	GPINTENB = 0x05 // enable interruption
	DEFVALA  = 0x06 // default value to compore for interruption
	DEFVALB  = 0x07 // default value to compore for interruption
	INTCONA  = 0x08 // compare against DEFVAL to interrupt if set to 1
	INTCONB  = 0x09 // compare against DEFVAL to interrupt if set to 1
	GPPUA    = 0x0C // configure pull up resistor
	GPPUB    = 0x0D // configure pull up resistor
	GPIOA    = 0x12 // instant read
	GPIOB    = 0x13 // instant read
	OLATA    = 0x14 // set output data
	OLATB    = 0x15 // set output data
	INTCAPA  = 0x0C // read captured value when interrupt
	INTCAPB  = 0x0D // read captured value when interrupt

)

type Hardware struct {
	bus embd.I2CBus
}

var table = Hardware{}

func (hardware *Hardware) init() {
	if hardware.bus != nil {
		return
	}
	hardware.bus = embd.NewI2CBus(1)
}

func (hardware *Hardware) OnBoardLeds() {
	log.Println("Configuring hardware...")
	defer embd.CloseLED()
	for index := 0; index < 20; index++ {
		embd.LEDToggle("LED0")
		time.Sleep(250 * time.Millisecond)
	}
	embd.LEDToggle("LED0")
	time.Sleep(2500 * time.Millisecond)
}

func (hardware *Hardware) I2CLeds() {
	hardware.bus.WriteByteToReg(ADDRESS, IODIRA, 0x80)
	hardware.bus.WriteByteToReg(ADDRESS, GPPUA, 0x80)
	hardware.bus.WriteByteToReg(ADDRESS, OLATA, 0x00)
	time.Sleep(250 * time.Millisecond)
	var index byte = 0
	for ; index < 7; index++ {
		hardware.bus.WriteByteToReg(ADDRESS, OLATA, index)
		time.Sleep(250 * time.Millisecond)
	}
}

func (hardware *Hardware) loop() {
	hardware.bus.WriteByteToReg(ADDRESS, IODIRA, 0x80) // input for A7
	hardware.bus.WriteByteToReg(ADDRESS, GPPUA, 0x80)  // enable pull up resistor for A7
	counter := 5
	for counter > 0 {
		value, err := hardware.bus.ReadByteFromReg(ADDRESS, GPIOA)
		if err != nil {
			log.Fatal("error while reading")
		}
		if value&0x80 != 0 {
			log.Println("Got a 1 !")
			counter--
		}
		time.Sleep(10 * time.Millisecond)
	}
}
