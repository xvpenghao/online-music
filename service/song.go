package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	ISongService interface {
		IBaseService
		//查询歌曲详情
		QuerySongDetail(req models.QuerySongDetailReq) (dbModel.Song, error)
		//查询歌曲的基本信息
		QuerySongBaseInfo(req models.QuerySongDetailReq) (dbModel.Song, error)
	}
)

func NewSongService(init IBaseServiceInit) ISongService {
	temp := allService[ServiceISong]
	result := temp.(ISongService)
	result.SetInitInfo(init)
	return result
}
