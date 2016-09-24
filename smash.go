package main

import (
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

	app.SetContent(&layout)
	app.Run()
}
