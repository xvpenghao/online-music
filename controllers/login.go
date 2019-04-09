package controllers

import (
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
		logs.Error("添加用户-service返回错误:%s", err.Error())
		return receiver.returnError("添加用户-(%v)", err.Error())
	}
	resp := models.LoginResp{
		Id:       result.Id,
		Name:     result.Name,
		Email:    result.Email,
		Gender:   result.Gender,
		Age:      result.Age,
		Birthday: result.Birthday.Format("2006-01-02"),
	}
	//获取用户的其他的资源信息
	receiver.SetSession(constants.SESSION_USER, resp)

	receiver.Redirect("/v1/index/indexUI", http.StatusFound)
	return nil
}

//@Title LoginOut
//@Description 退出登录
//@router /loginOut [get]
func (receiver *LoginController) LoginOut() error {
	receiver.DelSession(constants.SESSION_USER)

	receiver.Redirect("/v1/index/indexUI", http.StatusFound)
	return nil
}
