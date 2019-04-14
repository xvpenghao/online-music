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
		//创建歌单
		CreateSongCover(req models.CreateSongCoverReq) error
		//查询用户歌单列表
		QueryUserSongCoverList(req models.QueryUserSongCoverListReq) ([]dbModel.QueryUserSongCover, error)
	}
)

func NewSongCoverService(init IBaseServiceInit) ISongCoverService {
	temp := allService[ServiceISongCover]
	result := temp.(ISongCoverService)
	result.SetInitInfo(init)
	return result
}
