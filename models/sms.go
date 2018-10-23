package models

import (
	"bytes"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Sms struct {
	//To      []string `json:"to" `
	To      string `json:"to" `
	Content string `json:"content"`
}

func (s *Sms) FalconSendSms() (err error) {

	//	for _, to := range s.To {
	//		logs.Info("to:%s", to)
	//	}
	logs.Info("send sms to;%s,content:%s", s.To, s.Content)
	var body string
	body = "phone=" + s.To + "&msg=" + s.Content
	apiurl := beego.AppConfig.String("smsHost")
	req, err := http.NewRequest("POST", apiurl, bytes.NewBuffer([]byte(body)))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/html")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
