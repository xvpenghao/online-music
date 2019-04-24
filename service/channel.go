package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	IChannelService interface {
		IBaseService
		//添加渠道
		CreateChannel(req models.CreateChannelReq) error
		//查询渠道详情
		QueryChannelDetail(req models.QueryChannelDetailReq) (dbModel.ChannelDetail, error)
		//修改渠道信息
		ModifyChannel(req models.ModifyChannelReq) error
		//查询渠道列表
		QueryChannelList(req models.QueryChannelListReq) (dbModel.ChannelInfoList, error)
	}
)

func NewChannelService(init IBaseServiceInit) IChannelService {
	temp := allService[ServiceIChannel]
	result := temp.(IChannelService)
	result.SetInitInfo(init)
	return result
}
