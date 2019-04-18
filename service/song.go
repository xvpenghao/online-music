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
		//歌曲歌单id查询歌曲列表
		QueryUserSongList(req models.QueryUserSongListReq) ([]dbModel.UserSong, error)
		//添加歌曲
		CreateSong(req models.CreateSongReq) error
		//通过歌曲id查询歌曲信息
		QuerySongInfoById(req models.QuerySongDetailReq) (dbModel.SongTable, error)
		//协程版查询歌曲信息
		QuerySongBaseInfoChan(req models.QuerySongDetailReq, song chan dbModel.Song)
		//删除歌曲
		DeleteSong(req models.DeleteSongReq) error
	}
)

func NewSongService(init IBaseServiceInit) ISongService {
	temp := allService[ServiceISong]
	result := temp.(ISongService)
	result.SetInitInfo(init)
	return result
}
