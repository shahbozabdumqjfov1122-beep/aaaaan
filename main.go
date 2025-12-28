package main

import (
	_ "aaaaan/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
	beego.Run("0.0.0.0:1")

}
