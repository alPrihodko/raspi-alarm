package alarm

import (
	"log"
	"os"
)

const armedFile = "/var/lock/raspi-alarm-armed.lock"

type alarm struct {
	Armed bool
}

/*
Alarm signals that system armed and should alarm
*/
var Alarm alarm

func (q alarm) init() {
	var err error
	Alarm.Armed, err = Exists(armedFile)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("System state ", Alarm.Armed)
}

/*
Exists returns file exists
*/
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

/*
Arm arms system
*/
func (q alarm) Arm() error {
	newFile, err := os.Create(armedFile)
	if err != nil {
		return err
	}
	newFile.Close()
	Alarm.Armed = true
	return nil
}

/*
Disarm disarms system
*/
func (q alarm) Disarm() error {
	err := os.Remove(armedFile)
	if err != nil {
		return err
	}
	Alarm.Armed = false
	return nil
}
