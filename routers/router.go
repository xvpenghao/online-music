package routers

import (
	"github.com/astaxie/beego"
	"online-music/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
