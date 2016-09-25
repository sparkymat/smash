package widget

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/size"
)

type CommandArea struct {
	spartan.View
}

func (ca CommandArea) GetWidth() size.Size {
	return size.MatchParent
}

func (ca CommandArea) GetHeight() size.Size {
	return 3
}

func (ca CommandArea) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	msg := "[Press Space]"
	length := uint32(len(msg))
	width := right - left + 1
	height := bottom - top + 1

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', ca.ForegroundColor, ca.BackgroundColor)
		}
	}

	if length <= width {
		horizontalMidPoint := left + width/2
		msgLeft := horizontalMidPoint - length/2
		verticalMidPoint := top + height/2

		for i := msgLeft; i < msgLeft+length; i++ {
			termbox.SetCell(int(i), int(verticalMidPoint), rune(msg[i-msgLeft]), ca.ForegroundColor, ca.BackgroundColor)
		}
	}
}
