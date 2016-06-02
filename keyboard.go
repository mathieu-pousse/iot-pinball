package main

import (
	"github.com/nsf/termbox-go"
	"log"
)

type keyboardEventType int

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type direction int

const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       termbox.Key
}

func keyToDirection(k termbox.Key) direction {
	switch k {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowDown:
		return DOWN
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	default:
		return 0
	}
}

func listenToKeyboard(evChan chan keyboardEvent) {

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	for {
		log.Printf("waiting for intput")
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowDown:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowRight:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowUp:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyEsc:
				evChan <- keyboardEvent{eventType: END, key: ev.Key}
			default:
				if ev.Ch == 'r' {
					evChan <- keyboardEvent{eventType: RETRY, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
