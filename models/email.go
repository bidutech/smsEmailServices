package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/gomail.v2"
)

type Email struct {
	From           string   `json:"from" `
	To             []string `json:"to" `
	Cc             string   `json:"cc" `
	Subject        string   `json:"subject"`
	BodyType       string   `json:"bodytype"`
	Body           string   `json:"body"`
	AttachFileName string   `json:"attachfilename"`
	D              *gomail.Dialer
}

//func init() {
//	host := beego.AppConfig.String("emailHost")
//	port, _ := beego.AppConfig.Int("emailPort")
//	uname := beego.AppConfig.String("emailUser")
//	pwd := beego.AppConfig.String("emailPassword")
//	fmt.Println(host, port, uname, pwd)
//	D = gomail.NewDialer(host, port, uname, pwd)
//}

func (e *Email) FalconSendEmail() (err error) {
	logs.Info("send email")

	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	for _, to := range e.To {
		m.SetHeader("To", to)
		logs.Info("to:%s", to)
	}

	//	if len(e.Cc) > 0 {
	//		//m.SetAddressHeader("Cc", "dan@example.com", "dan")
	//	}

	m.SetHeader("Subject", e.Subject)
	m.SetBody(e.BodyType, e.Body)
	if len(e.AttachFileName) > 0 {
		m.Attach(e.AttachFileName)
	}
	logs.Info("%v", m)
	host := beego.AppConfig.String("emailHost")
	port, _ := beego.AppConfig.Int("emailPort")
	uname := beego.AppConfig.String("emailUser")
	pwd := beego.AppConfig.String("emailPassword")
	fmt.Println(string(host), port, uname, pwd, e.BodyType, e.Subject, e.Body)
	d := gomail.NewDialer(host, port, uname, pwd)
	logs.Info("host:%s,%d,%s,%s", d.Host, d.Port, d.Username, d.Password)
	err = d.DialAndSend(m)
	if err != nil {
		logs.Info("---%s", err.Error())
	}
	return err
}
