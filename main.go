package main

import (
	_ "BeeGoWeb/config"
	"BeeGoWeb/models"
	_ "BeeGoWeb/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

func init(){
	//初始化orm
	models.Init()
	// 生产环境不输出debug日志

	if beego.AppConfig.String("runmode") == "dev" {
		beego.SetLevel(beego.LevelDebug)
	}
	beego.AppConfig.Set("version", beego.AppConfig.String("AppVer"))
}
