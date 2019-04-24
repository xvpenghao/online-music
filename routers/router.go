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
		//歌单
		beego.NSNamespace("/songCover", beego.NSInclude(&controllers.SongCoverController{})),
	)

	//后台页面
	ns2 := beego.NewNamespace("admin",
		//首页
		beego.NSNamespace("/index", beego.NSInclude(&controllers.AdminIndexController{})),
		//数据
		beego.NSNamespace("/data", beego.NSInclude(&controllers.DataController{})),
		//渠道
		beego.NSNamespace("/channel", beego.NSInclude(&controllers.ChannelController{})),
		//后台普通用户
		beego.NSNamespace("/user", beego.NSInclude(&controllers.BUserController{})),
	)

	beego.AddNamespace(ns, ns2)
}
