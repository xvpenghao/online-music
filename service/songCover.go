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
		CreateSongCover(req models.CreateSongCoverReq) (dbModel.CreateSongCoverReply, error)
		//查询用户歌单列表
		QueryUserSongCoverList(req models.QueryUserSongCoverListReq) ([]dbModel.QueryUserSongCover, error)
		//创建收藏歌单
		CreateCollectSongCover(req models.CreateCollectSongCoverReq) error
		//根据歌单id查询信息
		QuerySongCoverById(req models.QueryCoverSongByIdReq) (dbModel.SongCoverInfo, error)
		//编辑歌单
		ModifySongCover(req models.ModifySongCoverReq) error
		//删除歌单
		DeleteSongCover(req models.DeleteSongCoverReq) error
	}
)

func NewSongCoverService(init IBaseServiceInit) ISongCoverService {
	temp := allService[ServiceISongCover]
	result := temp.(ISongCoverService)
	result.SetInitInfo(init)
	return result
}
