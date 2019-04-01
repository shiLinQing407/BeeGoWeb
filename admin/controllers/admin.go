/*
@Time : 2019/3/28 16:48 
@Author : shilinqing
@File : Admin
*/
package controllers

import (
	"BeeGoWeb/models/system"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"BeeGoWeb/libs"

)

type AdminController struct {
	BaseController
}

/**
 入口方法
 */
func (this *AdminController) Enter(){
	this.Layout = "layout/admin/layout.html"
	this.TplName = "admin/admin/index.html"
}

func (this *AdminController) Index() {
	this.Data["activeSidebarUrl"] = this.URLFor(this.controllerName + "." + this.actionName)
	this.display()
}

func (this *AdminController) Error() {
	this.Data["error"] = this.GetString(":error")
	this.setTpl("admin/error.html", "layout/admin/layout_pullbox.html")
}

// 个人信息
func (this *AdminController) Profile() {
	beego.ReadFromRequest(&this.Controller)
	user, _ := system.UserGetById(this.userId)

	if this.isPost() {
		flash := beego.NewFlash()
		user.Email = this.GetString("email")
		user.Update()
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")
		if password1 != "" {
			if len(password1) < 6 {
				flash.Error("密码长度必须大于6位")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else if password2 != password1 {
				flash.Error("两次输入的密码不一致")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else {
				user.Salt = string(utils.RandomCreateBytes(10))
				user.Password = libs.Md5([]byte(password1 + user.Salt))
				user.Update()
			}
		}
		flash.Success("修改成功！")
		flash.Store(&this.Controller)
		this.redirect(beego.URLFor(".Profile"))
	}

	this.Data["pageTitle"] = "个人信息"
	this.Data["user"] = user
	this.display()
}