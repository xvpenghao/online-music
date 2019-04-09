package main

import (
	"github.com/astaxie/beego"
	_ "online-music/common/redis"
	_ "online-music/routers"
	_ "online-music/service/impl"
)

func main() {
	beego.Run()
}
