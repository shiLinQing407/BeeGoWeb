package routers

import (
	admin "BeeGoWeb/admin/controllers"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

func init() {
	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})
	/*路由设置*/
	//固定路由
	beego.Router("/", &admin.AdminController{}, "*:Enter")
	//beego.Router("/login", &web.LoginController{}, "*:Login")
	//beego.Router("/logout", &web.LoginController{}, "*:Logout")
	/******************************		后台路由	******************************/
	beego.Router("/admin/index", &admin.AdminController{}, "*:Index")
	beego.Router("/admin/news", &admin.NewsController{}, "*:Index")
	beego.Router("/admin/news/loadList", &admin.NewsController{}, "GET:LoadList")
	beego.Router("/admin/news/edit", &admin.NewsController{}, "*:Edit")
	beego.Router("/admin/news/loadData", &admin.NewsController{}, "GET:LoadData")

	beego.Router("/admin/newsclass", &admin.NewsClassController{},"GET:Index")
	beego.Router("/admin/newsclass/loadList", &admin.NewsClassController{}, "GET:LoadList")
	beego.Router("/admin/newsclass/edit", &admin.NewsClassController{}, "*:Edit")
	beego.Router("/admin/newsclass/loadData", &admin.NewsClassController{}, "GET:LoadData")
	//beego.Router("/admin/login", &admin.LoginController{}, "*:Login")
	//beego.Router("/admin/logout", &admin.LoginController{}, "*:Logout")
	//beego.Router("/admin/config", &admin.ConfigController{}, "*:Config") //设置网站配置
	//beego.Router("/admin/config/edit", &admin.ConfigController{}, "*:EditConfig") //编辑网站配置
	//beego.Router("/admin/config/delete", &admin.ConfigController{}, "*:DeleteConfig") //删除网站配置
	//beego.Router("/login/gettime", &admin.LoginController{}, "*:GetTime") //获取系统时间
	//
	////用户有权管理的菜单列表（包括区域）
	beego.Router("/admin/menu/menutree", &admin.MenuController{}, "POST:UserMenuTree")
	//
	////菜单管理
	beego.Router("/admin/menu/index", &admin.MenuController{}, "*:Index")
	//
	////权限管理
	//beego.Router("/admin/auth/index", &admin.AuthController{}, "*:Index")
	//
	////角色管理
	//beego.Router("/admin/role/index", &admin.RoleController{}, "*:Index")
	//beego.Router("/admin/role/datagrid", &admin.RoleController{}, "POST:RoleDataGrid")
	//beego.Router("/admin/role/rolelist", &admin.RoleController{}, "POST:RoleList")
	//beego.Router("/admin/role/edit/?:id", &admin.RoleController{}, "*:RoleEdit")
	//beego.Router("/admin/role/delete", &admin.RoleController{}, "POST:RoleDelete")
	//
	////用户管理
	//beego.Router("/admin/user/index", &admin.UserController{}, "*:Index")
	//beego.Router("/admin/user/datagrid", &admin.UserController{}, "POST:UserDataGrid")
	//beego.Router("/admin/user/edit/?:id", &admin.UserController{}, "*:UserEdit")
	//beego.Router("/admin/user/delete", &admin.UserController{}, "POST:UserDelete")
	//
	//
	////图片管理
	//beego.Router("/admin/picture/index", &admin.PictureController{}, "*:Index")
	//beego.Router("/admin/picture/pictureupload", &admin.PictureController{}, "*:PictureUpload")
	//
	////日志管理
	//beego.Router("/admin/log/index", &admin.LogController{}, "*:Index")
	//beego.Router("/admin/log/system", &admin.LogController{}, "*:System")
	//beego.Router("/admin/log/opera", &admin.LogController{}, "*:Opera")


}
