package main

import . "github.com/makssof/Dota2Go/pkg/Dota2Catcher"

func main() {
	go CatchReadyButton()
	select {}
}
