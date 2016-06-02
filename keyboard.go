package main

import (
	"github.com/nsf/termbox-go"
	"pinball/hardware"
	"time"
	"log"
)

func findInputIdForKey(key termbox.Key) string {
	switch key {
	case termbox.KeyArrowLeft:
		return "leftFlipper"
	case termbox.KeyArrowRight:
		return "rightFlipper"
	case termbox.KeyEsc:
		return "power"
	default:
		return ""
	}
}

func listenToKeyboard(evChan chan hardware.InputEvent) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	var previous *termbox.Key

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			direction := hardware.Rising
			if previous != nil {
				if previous == &ev.Key {
					// already up
					direction = hardware.High
				} else {
					//  mark previous as falling
					evChan <- hardware.InputEvent{InputId: findInputIdForKey(*previous), Direction: hardware.Falling}
				}
				previous = nil
			}

			if inputId := findInputIdForKey(ev.Key); inputId != "" {
				evChan <- hardware.InputEvent{InputId: inputId, Direction: direction}
				previous = &ev.Key
			}

		case termbox.EventError:
			panic(ev.Err)
		}
		log.Printf("waiting...")
		time.Sleep(10)
	}
}
