package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"online-music/models"
	"online-music/service"
	"online-music/verify"
)

type BUserController struct {
	BaseController
}

// @Title BUserListUI
// @Description 后台用户列表UI
// @Failure exec error
// @router /bUserListUI [get]
func (receiver *BUserController) BUserListUI() error {
	receiver.BeforeStart("BUserListUI")

	receiver.TplName = "admin/user/userList.html"
	return nil
}

// @Title ModifyBUserUI
// @Description 查询用户根据ID
// @Param userId path string true "用户id"
// @Failure exec error
// @router /queryBUserByID/:userId [get]
func (receiver *BUserController) QueryBUserByID() error {
	receiver.BeforeStart("QueryBUserByID")

	req := models.QueryBUserByIDReq{
		UserId: receiver.GetString(":userId"),
	}
	bUserService := service.NewUserService(receiver.GetServiceInit())
	result, err := bUserService.QueryBUserByID(req)
	if err != nil {
		logs.Error("查询用户根据ID-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误")
	}
	formStr := "2006-01-02"
	resp := models.QueryBUserByIDResp{
		UserId:   result.UserId,
		UserName: result.UserName,
		Email:    result.Email,
		Age:      result.Age,
		Birthday: result.Birthday.Format(formStr),
		Gender:   result.Gender,
	}

	receiver.Data["user"] = resp

	receiver.TplName = "admin/user/userModify.html"
	return nil
}

// @Title ModifyBUser
// @Description 后台修改用户信息
// @Param req body models.ModifyBUserReq true "req"
// @Success resp {object} models.ModifyBUserResp true "resp"
// @Failure exec error
// @router /modifyBUser [put]
func (receiver *BUserController) ModifyBUser() error {
	receiver.BeforeStart("ModifyBUser")
	var req models.ModifyBUserReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("查询用户列表-参数解析错误(%v)", err.Error())
		return receiver.returnJSONError("参数解析错误")
	}
	reqU := models.ModifyUserReq{
		UserId:   req.UserId,
		UserName: req.UserName,
		Age:      req.Age,
		Gender:   req.Gender,
		Birthday: req.Birthday,
		Email:    req.Email,
	}
	err = verify.ModifyUserReqVerify(reqU)
	if err != nil {
		logs.Error("查询用户列表-参数错误(%v)", err.Error())
		return receiver.returnJSONError("参数错误")
	}

	bUService := service.NewUserService(receiver.GetServiceInit())
	err = bUService.ModifyUser(reqU)
	if err != nil {
		logs.Error("查询用户列表-service返回错误(%v)", err.Error())
		return receiver.returnJSONError("service返回错误")
	}

	var resp models.ModifyBUserResp
	return receiver.returnJSONSuccess(resp)
}

// @Title QueryBUserList
// @Description 查询用户列表
// @Param info body models.QueryBUserListReq true "req"
// @Success 200 {object} models.QueryBUserListResp "resp"
// @Failure exec error
// @router /queryBUserList [post]
func (receiver *BUserController) QueryBUserList() error {
	receiver.BeforeStart("QueryBUserList")
	var req models.QueryBUserListReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("查询用户列表-参数解析错误(%v)", err.Error())
		return receiver.returnJSONError("参数解析错误")
	}

	logs.Debug("%+v", req)

	userService := service.NewUserService(receiver.GetServiceInit())
	result, err := userService.QueryBUserList(req)
	if err != nil {
		logs.Error("查询用户列表-service返回错误(%v)", err.Error())
		return receiver.returnJSONError("service返回错误")
	}

	var resp models.QueryBUserListResp
	formStr := "2006-01-02"
	var bUser models.BUserInfo
	for _, v := range result.List {
		bUser.UserId = v.UserId
		bUser.UserName = v.UserName
		bUser.Age = v.Age
		bUser.Gender = v.Gender
		bUser.Email = v.Email
		bUser.Birthday = v.Birthday.Format(formStr)
		resp.List = append(resp.List, bUser)
	}

	resp.Page = result.Page

	return receiver.returnJSONSuccess(resp)
}
