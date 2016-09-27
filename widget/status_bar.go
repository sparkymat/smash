package widget

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
)

type StatusBar struct {
	spartan.View
}

func (sb StatusBar) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', sb.ForegroundColor, sb.BackgroundColor)
		}
	}
}
