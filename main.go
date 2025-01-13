package main

import (
	"beego-api-service/database"
	_ "beego-api-service/routers"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/swagger"
)

func main() {
	database.Init()
	web.Run()
}
