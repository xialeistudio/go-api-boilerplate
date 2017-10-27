package controllers

import "api-boilerplate/common"

type ErrorController struct {
	common.BaseController
}

func (c *ErrorController) Error404() {
	c.MakeResponse("Not Found", 404)
}
func (c *ErrorController) Error500() {
	c.MakeResponse("Server Internal Error", 500)
}
func (c *ErrorController) Error400() {
	c.MakeResponse("Bad Request", 400)
}
