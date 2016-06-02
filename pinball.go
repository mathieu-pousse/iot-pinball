package main

import (
	"pinball/hardware"
	"log"
	"github.com/nsf/termbox-go"
)

type Pinball struct {
	Hardware      *hardware.Hardware
	Inputs        []hardware.Input
	inputsChannel chan hardware.Event
	outputsChannel chan hardware.Event
}

var pinball = NewPinball()

func NewPinball() *Pinball {
	p := Pinball{}
	p.initialize()
	return &p
}

func (pinball *Pinball) initialize() {
	pinball.Hardware = new(hardware.Hardware)
	pinball.createInputs()
	pinball.inputsChannel = make(chan hardware.Event)
	if global.NoHardware {
		go listenToKeyboard(pinball.inputsChannel)
	} else {
		//go handleWithActualInputs
		log.Panic("not yet implemented")
	}
}

func (pinball *Pinball) eventLoop() {
	defer pinball.release()

	eventLoop:
	for {
		select {
		//case p := <-pointsChan:
		//	g.addPoints(p)
		case e := <-keyboardEventsChan:
			log.Printf("got %+v", e)
			switch e.eventType {
			case MOVE:
				log.Printf("move %v", e.eventType)
				keyToDirection(e.key)
			case END:
				log.Printf("break %v", e.eventType)
				break eventLoop
			default:
				log.Printf("default %v", e.eventType)
			}
		default:

		}

		log.Printf("event loop is over")
	}

}

func (pinball *Pinball) release() {
	termbox.Close()
}

func (pinball *Pinball) createInputs() {
	pinball.createInput("left flipper", hardware.PulseWhilePressed{OutputId: "leftFlipper"})
	pinball.createInput("right flipper", hardware.PulseWhilePressed{OutputId: "rightFlipper"})
	pinball.createInput("left inlane", hardware.Score{Plus: 1000})
	pinball.createInput("right inlane", hardware.Score{Plus: 1000})

	pinball.createInput("left outlane")  // replay before 30 seconds
	pinball.createInput("right outlane") // replay before 30 seconds
	pinball.createInput("drain") // replay before 30 seconds

	pinball.createInput("bumper1",
		hardware.PulseOnOutput{OutputId: "bumper1", Delay: hardware.SolenoidDelay()},
		hardware.Score{Plus: 1000})
	pinball.createInput("bumper2",
		hardware.PulseOnOutput{OutputId: "bumper2", Delay: hardware.SolenoidDelay()},
		hardware.Score{Plus: 1000})
	pinball.createInput("bumper3",
		hardware.PulseOnOutput{OutputId: "bumper3", Delay: hardware.SolenoidDelay()},
		hardware.Score{Plus: 1000})

	pinball.createInput("targetB")
	pinball.createInput("targetU")
	pinball.createInput("targetG")

}

func (pinball *Pinball) createInput(name string, ehs ...hardware.EventHandler) {
	input := hardware.Input{Name: name}
	for _, eh := range ehs {
		input.AddEventHandler(eh)
	}
	pinball.Inputs = append(pinball.Inputs, input)
}
