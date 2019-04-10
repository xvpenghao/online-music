package service

type (
	IChannelService interface {
		IBaseService
		//查询歌单列表
		QueryAllChannelInfo() error
	}
)

func NewChannelService(init IBaseServiceInit) IChannelService {
	temp := allService[ServiceIChannel]
	result := temp.(IChannelService)
	result.SetInitInfo(init)
	return result
}
