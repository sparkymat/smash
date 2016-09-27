package widget

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/size"
)

type Clock struct {
	spartan.View
	TimeString string

	Left   uint32
	Right  uint32
	Top    uint32
	Bottom uint32
}

func (c Clock) GetWidth() size.Size {
	return 9
}

func (c Clock) GetHeight() size.Size {
	return 1
}

func (c *Clock) OnStart() {
	c.UpdateTime()
	go c.RunTimerLoop()
}

func (c Clock) RunTimerLoop() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for _ = range ticker.C {
		c.UpdateTime()
	}
}

func (c *Clock) UpdateTime() {
	t := time.Now()
	c.TimeString = t.Format("15:04:05")
	c.Draw(c.Left, c.Top, c.Right, c.Bottom)
}

func (c *Clock) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	c.Left = left
	c.Right = right
	c.Top = top
	c.Bottom = bottom

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', c.ForegroundColor, c.BackgroundColor)
		}
	}

	tLength := uint32(len(c.TimeString))
	if tLength > (right - left + 1) {
		tLength = right - left + 1
	}

	for i := left; i < (left + tLength); i++ {
		termbox.SetCell(int(i), int(top), rune(c.TimeString[i-left]), c.ForegroundColor, c.BackgroundColor)
	}
}
