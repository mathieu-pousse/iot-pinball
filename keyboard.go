package main

import (
	"pinball/hardware"
	"log"
	"github.com/MarinX/keylogger"
)

func findInputForKeyCode(key uint16) string {

}

func inputEventLoop(evChan chan hardware.InputEvent) {
	devs, err := keylogger.NewDevices()
	if err != nil {
		log.Fatal(err)
		return
	}
	//for _, val := range devs {
	//	fmt.Println("Id->", val.Id, "Device->", val.Name)
	//}
	rd := keylogger.NewKeyLogger(devs[4])
	in, err := rd.Read()
	if err != nil {
		log.Fatalf("oops %v", err)
		return
	}
	for i := range in {
		//listen only key stroke event
		if i.Type == keylogger.EV_KEY {
			log.Printf("key %sv %v", i.Code, i.Value)
			key := findInputForKeyCode(i.Code)
			if key == nil {
				continue
			}
			switch i.Value {
			case 0:
				evChan <- hardware.InputEvent{InputId:key, Direction: hardware.Falling}
			case 1:
				evChan <- hardware.InputEvent{InputId:key, Direction: hardware.Rising}
				break
			case 2:
				evChan <- hardware.InputEvent{InputId:key, Direction: hardware.Rising}
				break
			}

		}
	}
}
