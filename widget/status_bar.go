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
	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', sb.ForegroundColor, sb.BackgroundColor)
		}
	}

	// Clock
	time := " 4:00 PM"
	tLength := uint32(len(time))

	for i := left; i < (left + tLength); i++ {
		termbox.SetCell(int(i), int(top), rune(time[i]), sb.ForegroundColor, sb.BackgroundColor)
	}
}
