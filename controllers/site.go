package controllers

import (
	"api-boilerplate/common"
	"github.com/astaxie/beego/config"
)

type SiteController struct {
	common.BaseController
}

var PkgInfo = make(map[string]interface{})

func init() {
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	PkgInfo["name"] = cfg.String("appname")
	PkgInfo["version"] = cfg.String("version")
	PkgInfo["build"] = cfg.String("build")
}
func (c *SiteController) Index() {
	c.Data["json"] = PkgInfo
	c.ServeJSON()
}
