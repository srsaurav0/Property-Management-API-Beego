package main

import (
	"beego-api-service/database"
	_ "beego-api-service/routers"

	"github.com/beego/beego"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/swagger"
)

func main() {
	database.Init()
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
