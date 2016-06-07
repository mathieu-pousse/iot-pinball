package hardware

import "log"

type InputDirection int

const (
	Rising InputDirection = 1
	Falling InputDirection = 2
	Low InputDirection = 3
	High InputDirection = 4
)

type InputEvent struct {
	Direction InputDirection
	InputId   string
}

type Input struct {
	Name      string
	listeners []InputEventHandler
}

func (input *Input) Configure() {

}

func (input *Input) AddEventHandler(eh InputEventHandler) {
	input.listeners = append(input.listeners, eh)
}

func (input *Input) Initialize() {

}

func (input *Input) OnEvent(event InputEvent) {
	log.Printf("Dispatching event to listeners of %s\n", event.InputId)
	for _, eh := range input.listeners {
		eh.Handle(event)
	}
}

type InputEventProducer interface {
	produce(inputEventChannel chan InputEvent)
}