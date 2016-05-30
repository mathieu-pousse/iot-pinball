package hardware

import (
	"log"
	"sync"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

const (
	ADDRESS  = 0x20 // base chip address
	IODIRA   = 0x00 // set IO direction (1 = Input / 0 = Output)
	IODIRB   = 0x01 // set IO direction (1 = Input / 0 = Output)
	IPOLA    = 0x02
	IPOLB    = 0x03
	GPINTENA = 0x04 // enable interruption
	GPINTENB = 0x05 // enable interruption
	DEFVALA  = 0x06 // default value to compore for interruption
	DEFVALB  = 0x07 // default value to compore for interruption
	INTCONA  = 0x08 // compare against DEFVAL to interrupt if set to 1
	INTCONB  = 0x09 // compare against DEFVAL to interrupt if set to 1
	IOCON    = 0x0A
	//	IOCON    = 0x0B
	GPPUA   = 0x0C // configure pull up resistor
	GPPUB   = 0x0D // configure pull up resistor
	INTFA   = 0x0E
	INTFB   = 0x0F
	INTCAPA = 0x10 // read captured value when interrupt
	INTCAPB = 0x11 // read captured value when interrupt
	GPIOA   = 0x12 // instant read
	GPIOB   = 0x13 // instant read
	OLATA   = 0x14 // set output data
	OLATB   = 0x15 // set output data

)

type Hardware struct {
	bus    embd.I2CBus
	gpio21 embd.DigitalPin
	wg     sync.WaitGroup
}

func (hardware *Hardware) init() {
	if hardware.bus != nil {
		return
	}

	hardware.bus = embd.NewI2CBus(1)
	gpio21, err := embd.NewDigitalPin(21)
	if err != nil {
		log.Fatal(err)
	}
	hardware.gpio21 = gpio21
	hardware.gpio21.SetDirection(embd.In)
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
	hardware.bus.WriteByteToReg(ADDRESS, IODIRA, 0x80)   // input for A7
	hardware.bus.WriteByteToReg(ADDRESS, GPPUA, 0x80)    // enable pull up resistor for A7
	hardware.bus.WriteByteToReg(ADDRESS, IPOLA, 0x80)    // inverse polarity
	hardware.bus.WriteByteToReg(ADDRESS, OLATA, 0x00)    // switch off everything
	hardware.bus.WriteByteToReg(ADDRESS, GPINTENA, 0x80) // interrupt for A7
	// hardware.bus.WriteByteToReg(ADDRESS, INTCONA, 0x80)  // compare against DEFVAL
	// hardware.bus.WriteByteToReg(ADDRESS, DEFVALA, 0x80)  // DEFVAL is 1

	// clear previous interruption if any
	hardware.bus.ReadByteFromReg(ADDRESS, INTCAPA)
	//	hardware.sequential()
	hardware.wInterrupt()
}

func (hardware *Hardware) wInterrupt() {

	quit := make(chan interface{})

	hardware.gpio21.Watch(embd.EdgeFalling, func(btn embd.DigitalPin) {
		log.Println("interrupted")
		value, _ := hardware.bus.ReadByteFromReg(ADDRESS, GPIOA)

		log.Printf("Got a %v !", value&0x80 != 0)
		if value&0x80 != 0 {
			hardware.bus.WriteByteToReg(ADDRESS, OLATA, 0x07)
			//		hardware.bus.WriteByteToReg(ADDRESS, OLATA, 0x07)
		} else {
			hardware.bus.WriteByteToReg(ADDRESS, OLATA, 0x00)
			//hardware.bus.WriteByteToReg(ADDRESS, OLATA, 0x00)
		}

	})

	log.Print("ready to wait")
	log.Printf("over %v", <-quit)
}

func (hardware *Hardware) sequential() {
	log.Printf("reading gpio21")
	count := 5
	for count > 0 {
		v, _ := hardware.gpio21.Read()
		//log.Printf("got %v %v", v, err)
		if v != 1 {
			value, _ := hardware.bus.ReadByteFromReg(ADDRESS, GPIOA)
			log.Printf("got %v %v", v, value)
			count -= 1
		}
	}
}
