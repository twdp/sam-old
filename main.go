package main

import (
	_ "tianwei.pro/sam/routers"

	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.ServerName = "sam"
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.EnableErrorsRender = false
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sam"
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.EnableErrorsShow = false
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {

	}
	beego.Run()
}
