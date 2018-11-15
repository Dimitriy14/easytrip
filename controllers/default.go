package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

<<<<<<< HEAD
func (this *MainController) Get() {
	this.Layout = "main_layout.tpl"
	this.TplName = "index.tpl"
=======
func (c *MainController) Get() {
	c.Layout = "main_layout.tpl"
	c.TplName = "index.tpl"
>>>>>>> dbd7f9e20e2df2bb8786b76473fcf66688bd2e77
}
