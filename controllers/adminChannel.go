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

// @Title QueryChannelList
// @Description 平台分类列表
// @Failure exec error
// @router /queryChannelList [get]
func (receiver *ChannelController) QueryChannelList() error {
	receiver.BeforeStart("QueryChannelList")

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

	receiver.TplName = "admin/channel/channelModify.html"
	return nil
}
