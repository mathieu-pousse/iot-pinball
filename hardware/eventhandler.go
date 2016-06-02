package hardware

import (
	"log"
	"time"
)

func SolenoidDelay() time.Duration {
	return 1 * time.Millisecond
}

type EventHandler interface {
	Handle(event InputEvent)
}

type Score struct {
	Plus int
}

func (eh Score) Handle(event InputEvent) {
	log.Printf("score +%v", eh.Plus)
}

type PulseOnOutput struct {
	OutputId string
	Delay    time.Duration
}

func (eh PulseOnOutput) Handle(event InputEvent) {
	log.Printf("setting %s to 1", eh.OutputId)
	go func() {
		time.Sleep(eh.Delay)
		log.Printf("setting %s to 0", eh.OutputId)
	}()
}

type PulseWhilePressed struct {
	OutputId string
}

func (eh PulseWhilePressed) Handle(event InputEvent) {
	if event.Direction == "rising" {
		log.Printf("setting %s to 1", eh.OutputId)
	} else if event.Direction == "falling" {
		log.Printf("setting %s to 0", eh.OutputId)
	}
}
