package app

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/smash/event"
	"github.com/sparkymat/smash/widget"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/size"
)

type smashApp struct {
	spartanApp          spartan.App
	eventHandlerChannel chan termbox.Event

	mainLayout  spartan.LinearLayout
	contentArea widget.ContentArea
	statusBar   widget.StatusBar
	commandArea widget.CommandArea
}

func New() *smashApp {
	sa := smashApp{}
	sa.spartanApp = spartan.New()
	sa.eventHandlerChannel = make(chan termbox.Event)

	sa.mainLayout = spartan.LinearLayout{}
	sa.mainLayout.Direction = direction.Vertical
	sa.mainLayout.Width = size.MatchParent
	sa.mainLayout.Height = size.MatchParent

	sa.contentArea = widget.ContentArea{}
	sa.contentArea.Width = size.MatchParent
	sa.contentArea.Height = size.MatchParent
	sa.mainLayout.AddChild(&sa.contentArea)

	sa.statusBar = widget.StatusBar{}
	sa.statusBar.Width = size.MatchParent
	sa.statusBar.Height = 1
	sa.statusBar.BackgroundColor = termbox.ColorBlue
	sa.statusBar.ForegroundColor = termbox.ColorWhite | termbox.AttrBold
	sa.mainLayout.AddChild(&sa.statusBar)

	sa.commandArea = widget.CommandArea{}
	sa.commandArea.Width = size.MatchParent
	sa.commandArea.Height = 3
	sa.mainLayout.AddChild(&sa.commandArea)

	sa.spartanApp.SetContent(sa.mainLayout)

	return &sa
}

func (sa *smashApp) Run() {
	go event.Handler(sa.eventHandlerChannel)
	sa.spartanApp.Run(sa.eventHandlerChannel)
}
