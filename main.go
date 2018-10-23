package main

import (
	"fmt"
	"os"

	_ "smsEmailServices/models"
	_ "smsEmailServices/routers"

	"gopkg.in/gomail.v2"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	rm := beego.BConfig.RunMode
	logfile := beego.AppConfig.String("logfile")
	logs.SetLogger(logs.AdapterFile, logfile)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	if rm != "dev" && rm != "test" && rm != "prod" {
		fmt.Println("ERR ENV  DIBUS_RUNMOD: ", beego.BConfig.RunMode)
		os.Exit(1)
	}
	fmt.Println("init")
}
func main() {
	fmt.Println("beego.BConfig.RunMode:", beego.BConfig.RunMode)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	logs.Info("beego.BConfig.RunMode:", beego.BConfig.RunMode)
	beego.Run()
}

func testemailmain() {

	m := gomail.NewMessage()
	m.SetHeader("From", "adigwizsupport@qq.com")
	m.SetHeader("To", "duo.@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "adigwizsupport@qq.com", "ZHkHUssssRAUoqYm31B")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		logs.Info("send ok..")
	}
}
