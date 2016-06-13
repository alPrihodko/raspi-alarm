package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hybridgroup/gobot"
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

	initButton(gbot)
	initMoveDetect(gbot)

	go gbot.Start()

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	//if stop != nil {
	//	stop <- true
	//}
}
