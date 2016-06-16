package alarm

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
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

	wg := new(sync.WaitGroup)
	wg.Add(1)
	exeCmd("/etc/init.d/uv4l_uvc start 046d:081b", wg)
	wg.Wait()
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
	wg := new(sync.WaitGroup)
	wg.Add(1)
	exeCmd("/etc/init.d/uv4l_uvc stop", wg)
	wg.Wait()

	Alarm.Armed = false
	return nil
}

func exeCmd(cmd string, wg *sync.WaitGroup) {
	log.Println(cmd)
	parts := strings.Fields(cmd)
	var out []byte
	var err error
	if len(parts) > 2 {
		out, err = exec.Command(parts[0], parts[1], parts[2]).Output()
	} else if len(parts) == 2 {
		out, err = exec.Command(parts[0], parts[1]).Output()
	} else {
		log.Println("Invalid arguments")
	}
	if err != nil {
		log.Println("error occured")
		log.Println(err.Error())
		log.Println(out)
	}
	log.Println(out)
	wg.Done()
}
