package models

import "online-music/common/utils"

//添加平台分类
type CreateChannelReq struct {
	//平台名称
	ChannelName string `json:"channelName"`
}

type CreateChannelResp struct {
	baseResp
}

//查询平台分类详情
type QueryChannelDetailReq struct {
	//渠道ID
	ChannelId string
}

type QueryChannelDetailResp struct {
	//渠道Id
	ChannelId string
	//渠道名称
	ChannelName string
}

//修改渠道信息
type ModifyChannelReq struct {
	//渠道Id
	ChannelId string `json:"channelId"`
	//渠道名称
	ChannelName string `json:"channelName"`
}

type ModifyChannelResp struct {
	baseResp
}

type QueryChannelListReq struct {
	//当前页 默认为1
	CurPage     int    `json:"curPage"`
	ChannelName string `json:"channelName"`
}

//渠道信息
type ChannelInfo struct {
	//歌曲来源渠道id
	ChannelId string `json:"channelId"`
	//渠道名称
	ChannelName string `json:"channelName"`
	//创建人
	CreateUser string `json:"createUser"`
	//更新时间
	UpdateTime string `json:"updateTime"`
}

type QueryChannelListResp struct {
	Page utils.Page    `json:"page"`
	List []ChannelInfo `json:"list"`
}
