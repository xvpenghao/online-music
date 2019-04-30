package controllers

import (
	"github.com/astaxie/beego/logs"
	"online-music/common/constants"
	"online-music/models"
)

type BLoginController struct {
	BaseController
}

// @Title BLoginUI
// @Description 登录UI
// @Failure exec error
// @router / [get]
func (receiver *BLoginController) BLoginUI() error {
	receiver.BeforeStart("BLoginUI")
	sessionIDStr := receiver.Ctx.Input.CruSession.SessionID()
	logs.Debug("BLoginUI-sessionIDStr", sessionIDStr)
	receiver.TplName = "admin/login.html"
	return nil
}

// @Title BLogin
// @Description 后台用户登录
// @Param info body models.BLoginReq true "req"
// @Success 200 {object} models.BLoginResp "resp"
// @Failure exec error
// @router / [post]
func (receiver *BLoginController) BLogin() error {
	receiver.BeforeStart("BLogin")

	session := receiver.StartSession()
	sessionId := session.SessionID()
	logs.Debug("sessionID", sessionId)
	//session走的内存session
	_ = session.Set(constants.BACK_SESSION_KEY, "123456")
	var resp models.BLoginResp
	return receiver.returnJSONSuccess(resp)
}
