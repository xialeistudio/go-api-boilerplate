package common

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) MakeResponse(message string, code int) {
	data := make(map[string]interface{})
	data["errcode"] = code
	data["errmsg"] = message
	c.Data["json"] = data
	c.ServeJSON()
}
