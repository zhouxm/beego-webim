package main

import (
	_ "WebIM/routers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

const (
	APP_VER = "0.1.1.0227"
)

func main() {
	logs.Info(web.BConfig.AppName, APP_VER)

	// Register template functions.
	web.AddFuncMap("i18n", i18n.Tr)

	web.Run()
}
