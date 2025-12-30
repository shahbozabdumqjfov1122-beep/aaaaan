package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BoshController struct {
	beego.Controller
}

func (c *BoshController) Get() {

	c.TplName = "1.html"
}
