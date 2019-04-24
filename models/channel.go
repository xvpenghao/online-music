package models

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
