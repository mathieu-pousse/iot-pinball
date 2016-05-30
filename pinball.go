package main

import "pinball/hardware"

type Pinball struct {
	Hardware *hardware.Hardware
	Inputs   []hardware.Input
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
}

func (pinball *Pinball) eventLoop() {

}

func (pinball *Pinball) release() {
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
		hardware.Score{Plus: 1000}
	)
	pinball.createInput("bumper2",
		hardware.PulseOnOutput{OutputId: "bumper2", Delay: hardware.SolenoidDelay()},
		hardware.Score{Plus: 1000}
	)
	pinball.createInput("bumper3",
		hardware.PulseOnOutput{OutputId: "bumper3", Delay: hardware.SolenoidDelay()},
		hardware.Score{Plus: 1000}
	)

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
