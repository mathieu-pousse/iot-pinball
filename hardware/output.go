package hardware

type OutputEvent struct {
	Direction string
	OutputId string
}

type Output struct {
	Name      string
	listeners []EventHandler
}

func (input *Output) Configure() {

}

func (input *Output) AddEventHandler(eh EventHandler) {
	input.listeners = append(input.listeners, eh)
}

func (input *Output) Initialize() {

}

func (input *Output) OnEvent(event OutputEvent) {
	for _, eh := range input.listeners {
		eh.Handle(event)
	}
}
