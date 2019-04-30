package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"online-music/common/constants"
)

//过滤用户请求
var filterUser = func(ctx *context.Context) {
	//它是怎么判断出浏览器关闭，就不能从session取出值了
	value := ctx.Input.Session(constants.BACK_SESSION_KEY)
	if value == nil {
		ctx.ResponseWriter.Write([]byte("请你登录"))
		return
	}
	valueStr := value.(string)
	logs.Debug("filterUser-%v\n", valueStr)
}
