package email

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendEmail(user, sendUserName, password, host, port, to, subject, body string) error {
	Port, _ := strconv.Atoi(port)
	m := gomail.NewMessage()
	m.SetHeader("From",  m.FormatAddress(user, sendUserName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, Port, user, password)
	err := d.DialAndSend(m)
	return err
}
