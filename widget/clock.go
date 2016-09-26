package widget

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/size"
)

type Clock struct {
	spartan.View
}

func (c Clock) GetWidth() size.Size {
	return 9
}

func (c Clock) GetHeight() size.Size {
	return 1
}

func (c Clock) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	t := time.Now()
	tString := t.Format("15:04:05")

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', c.ForegroundColor, c.BackgroundColor)
		}
	}

	tLength := uint32(len(tString))
	if tLength > (right - left + 1) {
		tLength = right - left + 1
	}

	for i := left; i < (left + tLength); i++ {
		termbox.SetCell(int(i), int(top), rune(tString[i-left]), c.ForegroundColor, c.BackgroundColor)
	}
}
