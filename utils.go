package main

import (
	"gopkg.in/gomail.v2"
)

/*
ReportAlert sends message to the system owner "emailTo"
*/
func ReportAlert(params ...string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)

	m.SetHeader("Subject", params[1])
	m.SetBody("text/html", params[2])
	if len(params) > 2 {
		m.Attach(params[3])
	}

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, googleAccountName, googleAPIKey)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
