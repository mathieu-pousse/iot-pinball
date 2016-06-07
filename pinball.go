package main

import (
	"pinball/hardware"
	"log"
	"github.com/nsf/termbox-go"
)

type Pinball struct {
	Hardware       *hardware.Hardware
	Inputs         []hardware.Input
	inputsChannel  chan hardware.InputEvent
	outputsChannel chan hardware.OutputEvent
}

func NewPinball() *Pinball {
	p := Pinball{}
	p.initialize()
	return &p
}

func (pinball *Pinball) initialize() {
	pinball.Hardware = new(hardware.Hardware)
	pinball.createInputs()
	pinball.inputsChannel = make(chan hardware.InputEvent)
	if global.NoHardware {
		go inputEventLoop(pinball.inputsChannel)
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
		case e := <-pinball.inputsChannel:
			if e.InputId == "power" && e.Direction == hardware.Rising {
				break eventLoop
			}
			pinball.dispatchInputEvent(e)

		default:

		}

	}
	log.Printf("event loop is over\n")
}

func (pinball *Pinball) dispatchInputEvent(e hardware.InputEvent) {
	log.Printf("### InputEvent %+v\n", e)
	for _, input := range pinball.Inputs {
		if input.Name == e.InputId {
			input.OnEvent(e)
			return
		}
	}
	log.Printf("no body care about this event on %s\n", e.InputId)
}

func (pinball *Pinball) release() {
	termbox.Close()
}

func (pinball *Pinball) createInputs() {
	pinball.createInput("leftFlipper", hardware.PulseWhilePressed{OutputId: "leftFlipper"})
	pinball.createInput("rightFlipper", hardware.PulseWhilePressed{OutputId: "rightFlipper"})
	pinball.createInput("leftInlane", hardware.Score{Plus: 1000})
	pinball.createInput("rightInlane", hardware.Score{Plus: 1000})

	pinball.createInput("leftOutlane")  // replay before 30 seconds
	pinball.createInput("rightOutlane") // replay before 30 seconds
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

func (pinball *Pinball) createInput(name string, ehs ...hardware.InputEventHandler) {
	input := hardware.Input{Name: name}
	for _, eh := range ehs {
		input.AddEventHandler(eh)
	}
	pinball.Inputs = append(pinball.Inputs, input)
}
