package controllers

import (
	"net/http"
	"online-music/models"
	"online-music/service"
)

type SongController struct {
	BaseController
}

//@Title QuerySongDetail
//@Description 查询歌曲详情
//@Param songId path string true "歌曲ID"
//@Success info {object} models.QuerySongDetailResp true "返回歌曲相应信息"
//@Failure exec error
//@router /querySongDetail/:songID [get]
func (receiver *SongController) QuerySongDetail() error {
	receiver.BeforeStart("QuerySongDetail")

	req := models.QuerySongDetailReq{
		SongId: receiver.GetString(":songID"),
	}

	var resp models.QuerySongDetailResp
	songService := service.NewSongService(receiver.GetServiceInit())
	result, err := songService.QuerySongDetail(req)
	if err != nil {
		receiver.Ctx.ResponseWriter.Status = http.StatusBadRequest
		reply := map[string]string{
			"resultMsg": err.Error(),
		}
		receiver.Data["json"] = reply
		receiver.ServeJSON()
		return nil
	}

	resp.SongId = result.SongId
	resp.SongName = result.SongName
	resp.Singer = result.Singer
	resp.SongAlbum = result.SongAlbum
	resp.SongCoverUrl = result.SongCoverUrl
	resp.SongPlayUrl = result.SongPlayUrl
	resp.SongLyric = result.SongLyric

	receiver.Data["json"] = resp
	receiver.ServeJSON()

	return nil
}
