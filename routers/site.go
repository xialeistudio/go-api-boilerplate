package routers

import (
	"github.com/astaxie/beego"
	"api-boilerplate/controllers"
)

func init() {
	beego.Router("/", &controllers.SiteController{}, "get:Index")
}
