package controllers

import (
	//"fmt"
	"net/http"
	"smsEmailServices/models"

	//"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/toolkits/web/param"
)

type SmsController struct {
	beego.Controller
}

func (this *SmsController) FalconSendSms() {

	body := this.Ctx.Input.RequestBody
	logs.Info("FalconSendSms:%s", string(body))
	var s models.Sms
	r := this.Ctx.Request
	var w http.ResponseWriter
	w = this.Ctx.ResponseWriter

	tos := param.MustString(r, "tos")
	s.Content = param.MustString(r, "content")
	//tos = strings.Replace(tos, ",", ";", -1)
	s.To = tos //strings.Split(tos, ";")

	err := s.FalconSendSms()
	//	var ret string
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//ret = fmt.Sprintf("%s,%d", err.Error(), http.StatusInternalServerError)
		logs.Info(err.Error())
	} else {
		http.Error(w, "success", http.StatusOK)
		//ret = fmt.Sprintf("%s,%d", "success", http.StatusOK)
		logs.Info("ok...")
	}
	//this.Controller.Ctx.WriteString(ret)
}
