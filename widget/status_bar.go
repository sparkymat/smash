package widget

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/size"
)

type StatusBar struct {
	spartan.View
	clock Clock
}

func (sb *StatusBar) OnStart() {
	sb.clock = Clock{}
	sb.clock.ForegroundColor = sb.ForegroundColor
	sb.clock.BackgroundColor = sb.BackgroundColor

	sb.clock.OnStart()
}

func (sb StatusBar) GetWidth() size.Size {
	return size.MatchParent
}

func (sb StatusBar) GetHeight() size.Size {
	return 1
}

func (sb StatusBar) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', sb.ForegroundColor, sb.BackgroundColor)
		}
	}

	// Clock
	clockLeft := right - uint32(sb.clock.GetWidth()) + 1
	if clockLeft < left {
		clockLeft = left
	}
	sb.clock.Draw(clockLeft, top, right, bottom)
}
