package main

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/smash/widget"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/size"
)

func main() {
	app := spartan.New()

	layout := spartan.LinearLayout{
		Direction:  direction.Vertical,
		IsBordered: false,
	}
	layout.Height = size.MatchParent
	layout.Width = size.MatchParent

	mainContentLayout := spartan.LinearLayout{
		Direction:  direction.Vertical,
		IsBordered: false,
	}
	mainContentLayout.Height = size.MatchParent
	mainContentLayout.Width = size.MatchParent

	sb := widget.StatusBar{}
	sb.ForegroundColor = termbox.ColorBlack
	sb.BackgroundColor = 0x667

	commandArea := widget.CommandArea{}

	layout.AddChild(&mainContentLayout)
	layout.AddChild(&sb)
	layout.AddChild(&commandArea)

	app.SetContent(&layout)
	app.Run()
}
