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

func initMoveDetect(gbot *gobot.Gobot) {
	r := raspi.NewRaspiAdaptor("raspi")
	button := gpio.NewButtonDriver(r, "button", movementSensor)

	work := func() {
		gobot.On(button.Event("push"), func(data interface{}) {
			log.Println("detected")
		})

		gobot.On(button.Event("release"), func(data interface{}) {
			log.Println("releived")
		})

	}
	robot := gobot.NewRobot("detectBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)

	gbot.AddRobot(robot)
}
