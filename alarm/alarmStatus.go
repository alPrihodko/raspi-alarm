package alarm

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const armedFile = "/etc/raspi-alarm-armed.lock"
const binpath = "/usr/local/bin/"

type alarm struct {
	Armed   bool
	ArmedAt int32
}

/*
Alarm signals that system armed and should alarm
*/
var Alarm alarm

func init() {
	var err error
	Alarm.Armed, err = Exists(armedFile)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Alarm system state ", Alarm.Armed)
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
	ExeCmd("arm.sh", wg)
	wg.Wait()
	Alarm.Armed = true
	Alarm.ArmedAt = int32(time.Now().Unix())
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
	wg := new(sync.WaitGroup)
	wg.Add(1)
	ExeCmd("disarm.sh", wg)
	wg.Wait()
	return nil
}

/*
ExeCmd execs shell script
*/
func ExeCmd(cmd string, wg *sync.WaitGroup) {
	log.Println(cmd)
	parts := strings.Fields(cmd)
	var out []byte
	var err error
	if len(parts) > 2 {
		out, err = exec.Command(binpath+parts[0], parts[1], parts[2]).Output()
	} else if len(parts) == 2 {
		out, err = exec.Command(binpath+parts[0], parts[1]).Output()
	} else if len(parts) == 1 {
		out, err = exec.Command(binpath + parts[0]).Output()
	} else {
		log.Println("Invalid arguments")
	}
	if err != nil {
		log.Println("error occured")
		log.Println(err.Error())
		log.Println(string(out))
	}
	log.Println(string(out))
	wg.Done()
}

/*
ExeCmdNoWait execs shell script
*/
func ExeCmdNoWait(cmd string) {
	log.Println(cmd)
	parts := strings.Fields(cmd)
	var out []byte
	var err error
	if len(parts) > 2 {
		out, err = exec.Command(binpath+parts[0], parts[1], parts[2]).Output()
	} else if len(parts) == 2 {
		out, err = exec.Command(binpath+parts[0], parts[1]).Output()
	} else if len(parts) == 1 {
		out, err = exec.Command(binpath + parts[0]).Output()
	} else {
		log.Println("Invalid arguments")
	}
	if err != nil {
		log.Println("error occured")
		log.Println(err.Error())
		log.Println(string(out))
	}
	log.Println(string(out))
}

/*
Exists returns file exists
*/
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)

	//log.Println(v)

	if err != nil {
		log.Println(err.Error())
	}

	if os.IsNotExist(err) {
		log.Println("seems no file: ", name)
		return false, nil
	}

	return err == nil, err
}
