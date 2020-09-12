package main

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	imagePart := robotgo.OpenBitmap("scheme")

	counter := 0
	for true {
		screen := robotgo.CaptureScreen()
		x, y := robotgo.FindBitmap(imagePart, screen)

		if x > -1 && y > -1 {
			w, h := robotgo.GetScreenSize()
			robotgo.MoveClick(w/2, h/2)
			counter++
		}
		robotgo.FreeBitmap(screen)
		time.Sleep(700 * time.Millisecond)
	}
}
