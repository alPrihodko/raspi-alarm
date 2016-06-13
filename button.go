package main

import (
	"log"
	//"os"
	//"time"
	//  "sync"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func initButton(gbot *gobot.Gobot) {
	r := raspi.NewRaspiAdaptor("raspi")
	button := gpio.NewButtonDriver(r, "button", activateButton)

	work := func() {
		gobot.On(button.Event("push"), func(data interface{}) {
			log.Println("button pressed")
		})

		gobot.On(button.Event("release"), func(data interface{}) {
			log.Println("button released")
		})

	}
	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)
	gbot.AddRobot(robot)

}
