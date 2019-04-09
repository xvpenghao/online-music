package routers

import (
	"github.com/astaxie/beego"
	"online-music/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("v1",
		//用户
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		//首页
		beego.NSNamespace("/index", beego.NSInclude(&controllers.IndexController{})),
		//登录
		beego.NSNamespace("/login", beego.NSInclude(&controllers.LoginController{})),
		//歌曲
		beego.NSNamespace("/song", beego.NSInclude(&controllers.SongController{})),
	)

	beego.AddNamespace(ns)
}
