package widget

import (
	"time"

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

	sb.drawClocK(right-9, top, right, bottom)
}

func (sb StatusBar) drawClocK(left uint32, top uint32, right uint32, bottom uint32) {
	t := time.Now()
	timeText := t.Format(" 15:04:05")

	width := right - left + 1
	length := uint32(len(timeText))

	if length > width {
		length = width
	}

	for i := left; i < (left + length); i++ {
		termbox.SetCell(int(i), int(top+(bottom-top)/2), rune(timeText[i-left]), sb.ForegroundColor, sb.BackgroundColor)
	}
}

func (sb StatusBar) OnTick() {
}
