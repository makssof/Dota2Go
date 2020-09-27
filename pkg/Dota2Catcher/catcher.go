package Dota2Catcher

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var SchemeFile = os.TempDir() + "/scheme_" + strconv.FormatInt(time.Now().Unix(), 10)

func CatchReadyButton(closeCh chan bool) {
	scheme, err := FSByte(false, "/assets/scheme")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(SchemeFile, scheme, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	imagePart := robotgo.OpenBitmap(SchemeFile)
	shift := 1

handle:
	for {
		select {
		case <-closeCh:
			break handle
		default:
			screen := robotgo.CaptureScreen()
			x, y := robotgo.FindBitmap(imagePart, screen)

			if x > -1 && y > -1 {
				shift *= -1 // Trigger :hover of the Ready Button
				w, h := robotgo.GetScreenSize()
				robotgo.MoveClick(w/2, h/2+shift)
			}
			robotgo.FreeBitmap(screen)
			time.Sleep(700 * time.Millisecond)
		}
	}
}

func InitTrayIcon(closeCh chan bool) {
	systray.Run(func() {
		icon, err := FSByte(false, "/assets/icon.ico")
		if err != nil {
			fmt.Println(err)
			return
		}
		systray.SetIcon(icon)
		systray.SetTitle("Dota2Go Game Catcher")
		closeBtn := systray.AddMenuItem("Close", "Close Dota2Go")
		go func(closeCh chan bool) {
			<-closeBtn.ClickedCh
			systray.Quit()
			time.Sleep(100 * time.Millisecond)
			closeCh <- true
		}(closeCh)
	}, func() {})
}

func ClearResources() {
	err := os.Remove(SchemeFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}
