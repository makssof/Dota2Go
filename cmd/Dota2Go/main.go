package main

import . "github.com/makssof/Dota2Go/pkg/Dota2Catcher"

func main() {
	closeCh := make(chan bool)
	go InitTrayIcon(closeCh)
	go CatchReadyButton(closeCh)
	<-closeCh
	ClearResources()
}
