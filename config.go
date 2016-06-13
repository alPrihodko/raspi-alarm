package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	Active   string `json:"PumpState, string"`
	Disarmed string `json:"HeaterState, string"`
}

/*Config contains user preferences*/
type Config config

func (q *Config) saveConfig() error {
	b, err := json.Marshal(q)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(configFileName, b, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (q *Config) setDefault() {
	q.Active = "False"
	q.Disarmed = "True"
}

func (q *Config) loadConfig() error {

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Println("Cannot open config: ", err.Error())
		q.setDefault()
		return err
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(q); err != nil {
		log.Println("Error parsing config: ", err.Error())
		q.setDefault()
		return err
	}

	return nil
}
