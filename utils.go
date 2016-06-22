package main

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

/*
ReportAlert sends message to the system owner "emailTo"
*/
func ReportAlert(s string, b string, files ...string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)

	m.SetHeader("Subject", s)
	m.SetBody("text/html", b)
	for _, file := range files {
		m.Attach(file)
	}

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, googleAccountName, googleAPIKey)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
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
