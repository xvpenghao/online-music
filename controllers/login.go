package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service"
	"online-music/verify"
)

type LoginController struct {
	BaseController
}

//@Title LoginUI
//@Description 登录UI
//@router /loginUI [get]
func (receiver *LoginController) LoginUI() error {

	receiver.TplName = "login.html"
	return nil
}

//@Title LoginIn
//@Description 登录
//@router / [post]
func (receiver *LoginController) LoginIn() error {
	receiver.BeforeStart("LoginIn")

	var req models.LoginReq
	err := receiver.ParseForm(&req)
	if err != nil {
		logs.Error("登录-解析表单错误：%s", err.Error())
		return receiver.returnError("解析表单错误:(%v)", err.Error())
	}

	err = verify.LoginReqVerify(req)
	if err != nil {
		logs.Error("登录-参数错误:(%v)", err.Error())
		return receiver.returnError("登录-(%v)", err.Error())
	}

	loginService := service.NewLoginService(receiver.GetServiceInit())
	result, err := loginService.DoLogin(req)
	if err != nil {
		logs.Error("登录-service返回错误:%s", err.Error())
		return receiver.returnError("添加用户-(%v)", err.Error())
	}
	session := models.Session{
		UserId:   result.Id,
		UserName: result.Name,
		Email:    result.Email,
		Gender:   result.Gender,
		Age:      result.Age,
		Birthday: result.Birthday.Format("2006-01-02"),
	}
	//获取用户的其他的资源信息

	sessionService := service.NewSessionService(receiver.GetServiceInit())
	//遵循阿里redis使用规范：业务名:表名:id ，中间用逗号分隔开
	key := fmt.Sprintf("%s:%s:%s", "login", "user", session.UserId)
	value, _ := json.Marshal(session)

	//session存放用户信息
	err = sessionService.SetSession(key, string(value), constants.COOKIE_EXPIRE)
	if err != nil {
		logs.Error("登录-service返回错误:%s", err.Error())
		return receiver.returnError("添加用户-(%v)", err.Error())
	}

	//cookie存放用户id，关闭浏览器，则自动情况cookie
	receiver.SetSecureCookie(constants.COOKIE_SECRET, constants.COOKIE_NAME, key)

	receiver.Redirect("/v1/index/indexUI", http.StatusFound)
	return nil
}

//@Title LoginOut
//@Description 退出登录
//@router /loginOut [get]
func (receiver *LoginController) LoginOut() error {
	sessionId, b := receiver.GetSecureCookie(constants.COOKIE_SECRET, constants.COOKIE_NAME)
	//说明cookie中有值
	if b {
		sessionService := service.NewSessionService(receiver.GetServiceInit())
		sessionService.DelSession(sessionId)
	}
	receiver.DelSession(constants.SESSION_USER)

	receiver.Redirect("/v1/index/indexUI", http.StatusFound)
	return nil
}
