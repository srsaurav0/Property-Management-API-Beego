package main

import (
	_ "beego-api-service/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

