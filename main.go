package main

import (
	"github.com/astaxie/beego/orm"
	"tianwei.pro/sam-agent"
	_ "tianwei.pro/sam/models"
	_ "tianwei.pro/sam/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	beego.BConfig.ServerName = "sam"
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.EnableErrorsRender = false
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sam"

	// set default database

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("MysqlUrl") , 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true

		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.EnableErrorsShow = false
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("/*", beego.BeforeRouter, sam_agent.SamFilter)

	//beego.InsertFilter("/*",beego.BeforeRouter, func(context *context.Context) {
	//	context.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	//	context.ResponseWriter.Write([]byte("请重新登录"))
	//})

	beego.Run()
}
