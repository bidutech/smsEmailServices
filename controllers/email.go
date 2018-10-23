package controllers

import (
	//"fmt"
	"net/http"
	"smsEmailServices/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/toolkits/web/param"
)

type EmailController struct {
	beego.Controller
}

func (this *EmailController) FalconSendEmail() {
	logs.Info("SendEmail")
	var e models.Email
	r := this.Ctx.Request
	var w http.ResponseWriter
	w = this.Ctx.ResponseWriter

	token := param.String(r, "token", "")
	if beego.AppConfig.String("emailToken") != token {
		http.Error(w, "no privilege", http.StatusForbidden)
		return
	}

	tos := param.MustString(r, "tos")
	subject := param.MustString(r, "subject")
	content := param.MustString(r, "content")
	tos = strings.Replace(tos, ",", ";", -1)
	from := beego.AppConfig.String("emailUser")
	e.To = strings.Split(tos, ";")
	e.From = from
	e.Subject = subject
	e.BodyType = "text/html"
	e.Body = content
	err := e.FalconSendEmail()
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
