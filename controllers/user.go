package controllers

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service"
	"online-music/verify"
)

type UserController struct {
	BaseController
}

//@Title RegisterUI
//@Description 注册页面
//@router /registerUI [get]
func (receiver *UserController) RegisterUI() {

	receiver.TplName = "register.html"
}

//@Title UserDetailUI
//@Description 用户基本信息
//@router /userDetailUI [get]
func (receiver *UserController) UserDetailUI() {
	receiver.BeforeStart("UserDetailUI")
	sessionUser := receiver.GetSession(constants.SESSION_USER)
	result, _ := sessionUser.(models.LoginResp)
	receiver.Data[constants.SESSION_USER] = result

	receiver.TplName = "user/userInfo.html"
}

//@Title ModifyPwdUI
//@Description 修改密码UI
//@router /modifyPwdUI [get]
func (receiver *UserController) ModifyPwdUI() {
	receiver.BeforeStart("ModifyPwdUI")
	receiver.TplName = "user/modifyPwd.html"
}

//@Title CreateUser
//@Description 添加用户
//@router /createUser [post]
func (receiver *UserController) CreateUser() error {
	receiver.BeforeStart("CreateUser")
	var req models.CreateUserReq
	err := receiver.ParseForm(&req)
	if err != nil {
		logs.Error("添加用户-解析表单错误：%s", err.Error())
		return receiver.returnError("解析表单错误:(%v)", err.Error())
	}

	//参数的校验
	err = verify.CreateUserVerify(req)
	if err != nil {
		logs.Error("添加用户-参数错误:%s", err.Error())
		return receiver.returnError("添加用户-(%v)", err.Error())
	}

	userService := service.NewUserService(receiver.GetServiceInit())
	err = userService.CreateUser(req)
	if err != nil {
		logs.Error("添加用户-service返回错误:%s", err.Error())
		return receiver.returnError("添加用户-(%v)", err.Error())
	}

	//重定向到首页
	receiver.Redirect("/v1/index/indexUI", http.StatusFound)

	return nil
}

//@Title ModifyUser
//@Description 修改用户
//@router /modifyUser [post]
func (receiver *UserController) ModifyUser() error {
	receiver.BeforeStart("ModifyUser")
	var req models.ModifyUserReq
	err := receiver.ParseForm(&req)
	if err != nil {
		logs.Error("修改用户-解析表单错误：%s", err.Error())
		return receiver.returnError("解析表单错误:(%v)", err.Error())
	}

	err = verify.ModifyUserReqVerify(req)
	if err != nil {
		logs.Error("修改用户-参数错误:%s", err.Error())
		return receiver.returnError("修改用户-参数错误:(%v)", err.Error())
	}

	userService := service.NewUserService(receiver.GetServiceInit())
	err = userService.ModifyUser(req)
	if err != nil {
		logs.Error("修改用户-service返回错误:%s", err.Error())
		return receiver.returnError("修改用户-service错误:(%v)", err.Error())
	}

	//清除session，重定向到首页
	receiver.DelSession(constants.SESSION_USER)
	receiver.Redirect("/v1/index/indexUI", http.StatusFound)

	return nil
}

//@Title ModifyPwd
//@Description 修改密码
//@router /modifyPwd [post]
func (receiver *UserController) ModifyPwd() error {
	receiver.BeforeStart("ModifyPwd")
	var req models.ModifyPwdReq
	err := receiver.ParseForm(&req)
	if err != nil {
		logs.Error("修改密码-解析表单错误：(%v)", err.Error())
		return receiver.returnError("解析表单错误:(%v)", err.Error())
	}

	err = verify.ModifyPwdReqVerify(req)
	if err != nil {
		logs.Error("修改密码-参数错误:(%v)", err.Error())
		return receiver.returnError("修改密码-参数错误:(%v)", err.Error())
	}

	userService := service.NewUserService(receiver.GetServiceInit())
	err = userService.ModifyPwd(req)
	if err != nil {
		logs.Error("修改密码-service返回错误:%s", err.Error())
		return receiver.returnError("修改密码-service错误:(%v)", err.Error())
	}

	//清空session
	receiver.DelSession(constants.SESSION_USER)
	//调整首页
	receiver.Redirect("/v1/index/indexUI", http.StatusFound)

	return nil
}
