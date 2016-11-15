package app

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/smash/app/mode"
	"github.com/sparkymat/smash/event"
	"github.com/sparkymat/smash/widget"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/size"
)

type SmashApp struct {
	commandArea         widget.CommandArea
	contentArea         widget.ContentArea
	eventHandlerChannel chan termbox.Event
	mainLayout          spartan.LinearLayout
	spartanApp          spartan.App
	statusBar           widget.StatusBar
	mode                mode.Mode
	ticker              *time.Ticker
}

func New() *SmashApp {
	sa := SmashApp{}

	sa.spartanApp = spartan.New()

	sa.eventHandlerChannel = make(chan termbox.Event)
	sa.mode = mode.Idle

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

	sa.ticker = time.NewTicker(time.Millisecond * 200)

	return &sa
}

func (sa *SmashApp) Run() {
	go event.Handler(sa.eventHandlerChannel)
	go func() {
		for _ = range sa.ticker.C {
			sa.OnTick()
		}
	}()

	sa.spartanApp.Run(sa.eventHandlerChannel)
}

func (sa *SmashApp) CleanupForTermination() {
	sa.ticker.Stop()
}

func (sa *SmashApp) OnTick() {
	sa.contentArea.OnTick()
	sa.statusBar.OnTick()
	sa.commandArea.OnTick()
}
