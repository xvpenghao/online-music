package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service"
	"online-music/service/impl"
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
	//这里的返回的错误可以忽略
	sessionService := new(impl.SessionService)
	key, _ := receiver.GetSecureCookie(constants.COOKIE_SECRET, constants.COOKIE_NAME)
	user, _ := sessionService.GetSession(key)
	var session models.Session
	json.Unmarshal([]byte(user), &session)
	receiver.Data[constants.SESSION_USER] = session
	receiver.Session.UserId = session.UserId
	receiver.Session.UserName = session.UserName
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

//返回400错误
func (receiver *BaseController) returnJSONError(format string, a ...interface{}) (e error) {
	res := fmt.Sprintf(format, a...)
	receiver.Ctx.Output.Status = http.StatusBadRequest
	//返回的状态码
	errMap := map[string]string{
		"resultMsg": res,
	}
	receiver.Data["json"] = errMap
	receiver.ServeJSON()
	return
}

//返回成功
func (receiver *BaseController) returnJSONSuccess(data interface{}) (e error) {
	//返回的状态码
	receiver.Data["json"] = data
	receiver.ServeJSON()
	return
}
