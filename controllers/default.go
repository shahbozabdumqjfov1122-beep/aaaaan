package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	viewFile  = "views.txt"
	viewMutex sync.Mutex
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	viewMutex.Lock()
	defer viewMutex.Unlock()

	// eski sonni o‘qish
	data, _ := os.ReadFile(viewFile)
	views, _ := strconv.Atoi(strings.TrimSpace(string(data)))

	// +1
	views++

	// qayta yozish
	os.WriteFile(viewFile, []byte(strconv.Itoa(views)), 0644)

	c.Data["Views"] = views
	c.TplName = "1.html"
}
