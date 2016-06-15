package main

import (

	//"os"
	//"time"
	//  "sync"
	"raspi-alarm/alarm"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func initLed(gbot *gobot.Gobot, r *raspi.RaspiAdaptor) *gpio.LedDriver {
	led := gpio.NewLedDriver(r, "led", led)

	work := func() {

	}

	robot := gobot.NewRobot("ledBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)
	gbot.AddRobot(robot)
	return led
}

func blink(quit chan struct{}) {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			Led.Toggle()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func setLed() {
	if alarm.Alarm.Armed == true {
		Led.On()
	} else {
		Led.Off()
	}
}
