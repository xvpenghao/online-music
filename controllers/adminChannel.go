package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"online-music/models"
	"online-music/service"
)

type ChannelController struct {
	BaseController
}

// @Title ChannelAddUI
// @Description 添加平台分类UI
// @Failure exec error
// @router /createChannelUI [get]
func (receiver *ChannelController) CreateChannelUI() error {
	receiver.BeforeStart("ChannelAddUI")

	receiver.TplName = "admin/channel/channelAdd.html"
	return nil
}

// @Title QueryChannelListUI
// @Description 平台分类列表UI
// @Failure exec error
// @router /queryChannelListUI [get]
func (receiver *ChannelController) QueryChannelListUI() error {
	receiver.BeforeStart("QueryChannelListUI")

	receiver.TplName = "admin/channel/channelList.html"
	return nil
}

// @Title CreateChannel
// @Description 添加平台分类
// @Param info body models.CreateChannelReq true "req"
// @Success resp {object} models.CreateChannelResp true "resp"
// @Failure exec error
// @router /createChannel [post]
func (receiver *ChannelController) CreateChannel() error {
	receiver.BeforeStart("CreateChannel")
	var req models.CreateChannelReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("添加平台分类-参数解析错误(%v)", err.Error())
		return receiver.returnJSONError("参数解析错误")
	}

	channelService := service.NewChannelService(receiver.GetServiceInit())
	err = channelService.CreateChannel(req)
	if err != nil {
		logs.Error("添加平台分类-service返回错误(%v)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.CreateChannelResp
	return receiver.returnJSONSuccess(resp)
}

// @Title QueryChannelDetail
// @Description 查询渠道详情
// @Param channelId path string true "渠道id"
// @Success 200 {object} models.QueryChannelDetailResp "resp"
// @Failure exec error
// @router /queryChannelDetail/:channelId [get]
func (receiver *ChannelController) QueryChannelDetail() error {
	receiver.BeforeStart("QueryChannelDetail")

	req := models.QueryChannelDetailReq{
		ChannelId: receiver.GetString(":channelId"),
	}
	channelService := service.NewChannelService(receiver.GetServiceInit())
	result, err := channelService.QueryChannelDetail(req)
	if err != nil {
		logs.Error("查询渠道详情-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误")
	}
	resp := models.QueryChannelDetailResp{
		ChannelId:   result.ChannelId,
		ChannelName: result.ChannelName,
	}

	receiver.Data["channel"] = resp
	receiver.TplName = "admin/channel/channelModify.html"
	return nil
}

// @Title QueryChannelList
// @Description 查询渠道列表
// @Param req body models.QueryChannelListReq true "req"
// @Success resp {object} models.QueryChannelListResp true "resp"
// @Failure exec error
// @router /queryChannelList [post]
func (receiver *ChannelController) QueryChannelList() error {
	receiver.BeforeStart("QueryChannelList")
	var req models.QueryChannelListReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("查询渠道列表-参数解析错误(%v)", err.Error())
		return receiver.returnJSONError("参数解析错误")
	}

	channelService := service.NewChannelService(receiver.GetServiceInit())
	result, err := channelService.QueryChannelList(req)
	if err != nil {
		logs.Error("查询渠道列表-service返回错误(%v)", err.Error())
		return receiver.returnJSONError("service返回错误")
	}

	var resp models.QueryChannelListResp
	resp.Page = result.Page
	var c models.ChannelInfo
	formStr := "2006-01-02 15:04"
	for _, v := range result.List {
		c.ChannelId = v.ChannelId
		c.ChannelName = v.ChannelName
		c.CreateUser = v.CreateUser
		c.UpdateTime = v.UpdateTime.Format(formStr)
		resp.List = append(resp.List, c)
	}

	return receiver.returnJSONSuccess(resp)
}
