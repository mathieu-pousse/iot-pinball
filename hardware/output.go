package hardware

type OutputEvent struct {
	Direction string
	OutputId string
}

type Output struct {
	Name      string
	listeners []OutputEventHandler
}

func (input *Output) Configure() {

}

func (input *Output) AddEventHandler(eh OutputEventHandler) {
	input.listeners = append(input.listeners, eh)
}

func (input *Output) Initialize() {

}

func (input *Output) OnEvent(event OutputEvent) {
	for _, eh := range input.listeners {
		eh.Handle(event)
	}
}
