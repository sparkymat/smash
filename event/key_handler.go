package event

import (
	"os"

	termbox "github.com/nsf/termbox-go"
)

func CreateKeyHandlerChannel() chan termbox.Event {
	ch := make(chan termbox.Event)

	return ch
}

func HandleEvents(ch chan termbox.Event) {
	for {
		ev := <-ch
		go handleEvent(ev)
	}
}

func handleEvent(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyEsc {
			os.Exit(0)
		}
	}
}
