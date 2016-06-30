package main

import (
	"log"
	"raspi-alarm/alarm"
	"time"
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
				diff := int32(time.Now().Unix()) - alarm.Alarm.ArmedAt
				log.Println("Arm diff: ", diff, " secs")
				if diff > 120 {
					ReportAlert("Move detected", "Sendor 1 detected move.")
				}
				go alarm.ExeCmdNoWait("/home/pi/w/go/src/raspi-alarm/alarm.sh")
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
