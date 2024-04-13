package main

import (
	_ "chat/routers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

const (
	APP_VER = "0.1.1.0227"
)

func init() {
	logs.Info(web.BConfig.AppName, APP_VER)
}

func main() {

	// Register template functions.
	web.AddFuncMap("i18n", i18n.Tr)

	web.Run()
}
