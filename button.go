package main

import (
	"log"
	"time"
	//"os"
	//"time"
	//  "sync"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func initButton(gbot *gobot.Gobot, r *raspi.RaspiAdaptor) {
	button := gpio.NewButtonDriver(r, "button", activateButton)
	work := func() {
		var timer time.Time
		var quit chan struct{}

		gobot.On(button.Event("push"), func(data interface{}) {
			log.Println("button pressed")
			timer = time.Now()
			quit = make(chan struct{})
			go blink(quit)
		})

		gobot.On(button.Event("release"), func(data interface{}) {
			log.Println("button released")
			//if !isChanClosed(quit) {
			close(quit)
			//}

			if time.Since(timer) > 5*time.Second {
				Led.On()
			} else {
				Led.Off()
			}
		})

	}
	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)
	gbot.AddRobot(robot)

}

func isChanClosed(ch chan struct{}) bool {
	if len(ch) == 0 {
		select {
		case _, ok := <-ch:
			return !ok
		}
	}
	return false
}
