package controllers

import (
	"online-music/common/constants"
	"online-music/models"
)

type IndexController struct {
	BaseController
}

//@Title IndexUI
//@Description 首页页面
//@Failure exec error
//@router /indexUI [get]
func (receiver *IndexController) IndexUI() {
	sessionUser := receiver.GetSession(constants.SESSION_USER)
	result, _ := sessionUser.(models.LoginResp)
	receiver.Data[constants.SESSION_USER] = result
	receiver.TplName = "index.html"
}

//@Title ErrorUI
//@Description 错误页面
//@Failure exec error
//@router /errorUI [get]
func (receiver *IndexController) ErrorUI() {

	/*user := models.CreateUserReq{
		UserName:"张三",
	}

	receiver.Data["errorMsg"] = "参数错误"
	receiver.Data["user"] = "参数错误"*/

	receiver.TplName = "errors.html"
}
