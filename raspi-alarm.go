package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"raspi-alarm/alarm"
	"sync"
	"syscall"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"golang.org/x/net/websocket"
)

const configFileName = "/etc/raspi-alarm.conf"

const activateButton = "13" //27
const movementSensor = "11"
const led = "12"

/*
INTERVAL  Check sensors status with interval
*/
var INTERVAL int

//var err error

var conf Config

/*
Led is alarm led indicator
*/
var Led *gpio.LedDriver

type socketConns struct {
	ws   map[int32]*websocket.Conn
	lock *sync.Mutex
}

var conns socketConns
var rconns socketConns

//var stop chan bool

func main() {

	err := conf.loadConfig()
	if err != nil {
		log.Println("Likely use default configuration")
	}

	gbot := gobot.NewGobot()
	ra := raspi.NewRaspiAdaptor("raspi")

	conns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}

	flag.IntVar(&INTERVAL, "timeout", 60, "Timeout?")
	flag.Parse()

	log.Println("Timeout interval to track sensors: ", INTERVAL)

	//http.Handle("/echo", websocket.Handler(echoHandler))

	http.Handle("/", http.FileServer(http.Dir("/home/pi/w/go/src/raspi-alarm")))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, os.Kill)
	signal.Notify(c, syscall.SIGABRT)

	go func() {
		<-c
		//TODO: return relay to initial state
		os.Exit(1)
	}()

	initMoveDetect(gbot, ra)
	initButton(gbot, ra)
	Led = initLed(gbot, ra)

	if alarm.Alarm.Armed == true {
		err = alarm.Alarm.Disarm()
		if err != nil {
			log.Println("Cannot disarm system")
		}
	} else {
		err = alarm.Alarm.Arm()
		if err != nil {
			log.Println("Cannot arm system")
		}
	}

	go gbot.Start()

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	//if stop != nil {
	//	stop <- true
	//}
}
