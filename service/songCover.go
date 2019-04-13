package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	ISongCoverService interface {
		IBaseService
		//查询歌单列表
		QuerySongCoverList(req models.QuerySongCoverListReq) ([]dbModel.SongCover, error)
		//根据歌单id查询歌曲列表
		QuerySongList(req models.QuerySongListReq) ([]dbModel.Song, error)
	}
)

func NewSongCoverService(init IBaseServiceInit) ISongCoverService {
	temp := allService[ServiceISongCover]
	result := temp.(ISongCoverService)
	result.SetInitInfo(init)
	return result
}
