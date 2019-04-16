package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
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

// @Title ModifySongCoverList
// @Description 根据歌单id得到歌曲列表
// @Param songCoverId path string true "歌单id"
// @Failure exec error
// @router /queryUserSongList/:songCoverId [get]
func (receiver *SongController) QueryUserSongList() error {
	receiver.BeforeStart("CustomerSongCoverList")

	req := models.QueryUserSongListReq{
		SongCoverId: receiver.GetString(":songCoverId"),
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	result, err := songService.QueryUserSongList(req)
	if err != nil {
		logs.Error("根据歌单id得到歌曲列表-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误：(%v)", err.Error())
	}
	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	songCover, err := songCoverService.QuerySongCoverById(models.QueryCoverSongByIdReq{SongCoverId: req.SongCoverId})
	if err != nil {
		logs.Error("根据歌单id得到歌曲列表-查询歌单详情service返回错误：(%v)", err.Error())
		return receiver.returnError("查询歌单详情service返回错误：(%v)", err.Error())
	}
	var resp models.QueryUserSongListResp
	var song models.Song
	for _, v := range result {
		song.SongId = v.SongId
		song.Singer = v.Singer
		song.SongName = v.SongName
		song.SongAlbum = v.SongAlbum
		song.SongLyric = v.SongLyric
		song.SongPlayUrl = v.SongPlayUrl
		song.SongCoverUrl = v.SongCoverUrl
		resp.UserSongList = append(resp.UserSongList, song)
	}

	resp.SongCoverId = songCover.ID
	resp.SongCoverName = songCover.SongCoverName
	resp.SongCoverImgUrl = songCover.CoverUrl

	receiver.Data["resp"] = resp
	receiver.TplName = "song/customerSongList.html"

	return nil
}

// @Title CreateSong
// @Description 添加歌曲，添加歌曲到歌单
// @Param info body models.CreateSongReq true "req"
// @Success 200 {object} models.CreateSongResp "resp"
// @Failure exec error
// @router /createSong [post]
func (receiver *SongController) CreateSong() error {
	receiver.BeforeStart("CreateSong")

	if receiver.Session.UserId == "" {
		logs.Error("添加歌曲-用户未登录不能创建歌单")
		return receiver.returnJSONError("对不起，您未登录，不能创建歌单，请登录后操作")
	}

	var req models.CreateSongReq
	var resp models.CreateSongResp
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("添加歌曲-解析参数失败(%v)", err.Error())
		return receiver.returnJSONError("解析参数失败")
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	//爬虫获取歌曲详情
	song, err := songService.QuerySongDetail(models.QuerySongDetailReq{SongId: req.SongId})
	if err != nil {
		logs.Error("添加歌曲-根据歌曲id爬取歌曲详情service返回错误：(%v)", err.Error())
		return receiver.returnJSONError("根据歌曲id爬取歌曲详情service返回错误:(%v)", err.Error())
	}

	//爬虫到的歌曲进行赋值
	req.SongInfo.SongId = song.SongId
	req.SongInfo.SongName = song.SongName
	req.SongInfo.Singer = song.Singer
	req.SongInfo.SongAlbum = song.SongAlbum
	req.SongInfo.SongCoverUrl = song.SongCoverUrl
	req.SongInfo.SongPlayUrl = song.SongPlayUrl
	req.SongInfo.SongLyric = song.SongLyric
	err = songService.CreateSong(req)
	if err != nil {
		logs.Error("添加歌曲-service返回错误：(%v)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	return receiver.returnJSONSuccess(resp)
}
