package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service"
)

type BaseController struct {
	beego.Controller
	BaseControllerInit
}

type BaseControllerInit struct {
	Session models.Session
}

// Prepare : 处理完路由后调用
func (receiver *BaseController) Prepare() {
	sessionUser := receiver.GetSession(constants.SESSION_USER)
	result, _ := sessionUser.(models.LoginResp)
	receiver.Session.UserId = result.Id
	receiver.Session.UserName = result.Name
}

func (receiver *BaseController) BeforeStart(methodName string) {
	logs.Debug("%s", methodName)
}

func (receiver *BaseController) GetServiceInit() service.IBaseServiceInit {
	init := receiver.BaseControllerInit
	req := models.BaseRequest{
		UserID:   init.Session.UserId,
		UserName: init.Session.UserName,
	}
	return service.NewBaseServiceInit(req)
}

//返回参数错误
func (receiver *BaseController) returnError(format string, a ...interface{}) (e error) {
	res := fmt.Sprintf(format, a...)
	//返回的状态码
	receiver.Data["errorMsg"] = res
	receiver.TplName = "errors.html"
	return
}
