package hardware

import (
	"log"
)

type OutputEventHandler interface {
	Handle(event OutputEvent)
}

type LogOutputEventHandler struct {

}

func (eh LogOutputEventHandler) Handle(event OutputEvent) {
	log.Printf("%+v\n")
}
