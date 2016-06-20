package main

import (
	"log"
	"raspi-alarm/alarm"
	//"os"
	//"time"
	//  "sync"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func initMoveDetect(gbot *gobot.Gobot, r *raspi.RaspiAdaptor) {
	button := gpio.NewButtonDriver(r, "button", movementSensor)

	work := func() {
		gobot.On(button.Event("push"), func(data interface{}) {
			log.Println("detected")
			if alarm.Alarm.Armed {
				ReportAlert("Move detected", "Sendor 1 detected move.", nil)
			}
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
