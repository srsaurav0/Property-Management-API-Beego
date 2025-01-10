package main

import (
	"beego-api-service/database"
	_ "beego-api-service/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.Init()
	beego.Run()
}
