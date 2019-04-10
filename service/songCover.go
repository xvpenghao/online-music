package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	ISongCoverService interface {
		IBaseService
		//查询歌单列表
		QuerySongCoverList(req models.QuerySongCoverListReq) ([]dbModel.SongConver, error)
	}
)

func NewSongCoverService(init IBaseServiceInit) ISongCoverService {
	temp := allService[ServiceISongCover]
	result := temp.(ISongCoverService)
	result.SetInitInfo(init)
	return result
}
