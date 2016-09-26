package widget

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/size"
)

type StatusBar struct {
	spartan.View
}

func (sb StatusBar) GetWidth() size.Size {
	return size.MatchParent
}

func (sb StatusBar) GetHeight() size.Size {
	return 1
}

func (sb StatusBar) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	clock := Clock{}
	clock.ForegroundColor = sb.ForegroundColor
	clock.BackgroundColor = sb.BackgroundColor

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', sb.ForegroundColor, sb.BackgroundColor)
		}
	}

	// Clock
	clockLeft := right - uint32(clock.GetWidth()) + 1
	if clockLeft < left {
		clockLeft = left
	}
	clock.Draw(clockLeft, top, right, bottom)
}
