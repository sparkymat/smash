package event

import (
	"os"

	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
)

func CreateKeyHandlerChannel() chan termbox.Event {
	ch := make(chan termbox.Event)

	return ch
}

func HandleEvents(ch chan termbox.Event, app *spartan.App) {
	for {
		ev := <-ch
		go handleEvent(ev, app)
	}
}

func handleEvent(ev termbox.Event, app *spartan.App) {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyEsc {
			os.Exit(0)
		} else if ev.Key == termbox.KeyF5 {
			app.Redraw()
		}
	}
}
