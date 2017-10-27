package main

import (
	_ "api-boilerplate/routers"

	"github.com/astaxie/beego"
	"api-boilerplate/controllers"
	"runtime"
	"github.com/astaxie/beego/logs"
	"os"
)

func init() {
	logs.SetLogger("console")
	logs.Async(1e3)
	beego.ErrorController(&controllers.ErrorController{})
}
func main() {
	instances := runtime.NumCPU()
	runtime.GOMAXPROCS(instances)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logs.Info("%d running in %d instances", os.Getpid(), instances)
	beego.Run()
}
