package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"net/http"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service"
	"online-music/service/dbModel"
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
	//查询库中歌曲是否存在，若存在则不需要爬虫
	dbSong, err := songService.QuerySongInfoById(models.QuerySongDetailReq{SongId: req.SongId})
	if err != nil {
		logs.Error("添加歌曲-根据歌曲id查询歌曲详情service返回错误：(%v)", err.Error())
		return receiver.returnJSONError("根据歌曲id查询歌曲详情service返回错误:(%v)", err.Error())
	}

	if dbSong.SongId == "" {
		//爬虫获取歌曲详情
		song, err := songService.QuerySongDetail(models.QuerySongDetailReq{SongId: req.SongId})
		if err != nil {
			logs.Error("添加歌曲-根据歌曲id爬取歌曲详情service返回错误：(%v)", err.Error())
			return receiver.returnJSONError("根据歌曲id爬取歌曲详情service返回错误:(%v)", err.Error())
		}
		dbSong.SongId = song.SongId
		dbSong.SongName = song.SongName
		dbSong.Singer = song.Singer
		dbSong.SongAlbum = song.SongAlbum
		dbSong.SongCoverUrl = song.SongCoverUrl
		dbSong.SongPlayUrl = song.SongPlayUrl
		dbSong.SongLyric = song.SongLyric
	}

	//爬虫到的歌曲进行赋值
	req.SongInfo.SongId = dbSong.SongId
	req.SongInfo.SongName = dbSong.SongName
	req.SongInfo.Singer = dbSong.Singer
	req.SongInfo.SongAlbum = dbSong.SongAlbum
	req.SongInfo.SongCoverUrl = dbSong.SongCoverUrl
	req.SongInfo.SongPlayUrl = dbSong.SongPlayUrl
	req.SongInfo.SongLyric = dbSong.SongLyric
	err = songService.CreateSong(req)
	if err != nil {
		logs.Error("添加歌曲-service返回错误：(%v)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	return receiver.returnJSONSuccess(resp)
}

// @Title ModifySongCoverList
// @Description 根据收藏歌单id得到歌曲列表信息
// @Param songCoverId path string true "歌单id"
// @Failure exec error
// @router /queryCollectSCoverSongList/:songCoverId [get]
func (receiver *SongController) QueryCollectSCoverSongList() error {
	receiver.BeforeStart("CustomerSongCoverList")

	req := models.QueryUserSongListReq{
		SongCoverId: receiver.GetString(":songCoverId"),
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	songCover, err := songCoverService.QuerySongCoverById(models.QueryCoverSongByIdReq{SongCoverId: req.SongCoverId})
	if err != nil {
		logs.Error("根据歌单id得到歌曲列表-查询歌单详情service返回错误：(%v)", err.Error())
		return receiver.returnError("查询歌单详情service返回错误：(%v)", err.Error())
	}

	collectSongList, err := songCoverService.QuerySongList(models.QuerySongListReq{SongCoverId: req.SongCoverId})
	if err != nil {
		logs.Error("根据歌单id得到歌曲列表-爬取歌曲信息service返回错误：(%v)", err.Error())
		return receiver.returnError("爬取歌曲信息service返回错误：(%v)", err.Error())
	}
	var resp models.QueryUserSongListResp
	var querySongDetailReq models.QuerySongDetailReq
	var dbSong dbModel.Song
	var song models.Song
	for i, v := range collectSongList {
		//FIXME 这里需要优化，目前先定位3首歌曲
		if i == constants.SPIDER_SONG_COUNT && constants.SPIDER_SONG_COUNT > 0 {
			break
		}
		querySongDetailReq.SongId = v.SongId
		//FIXME 这里需要优化,考虑使用携程
		dbSong, err = songService.QuerySongBaseInfo(querySongDetailReq)
		if err != nil {
			logs.Error("根据歌单id获取歌曲列表-查询歌单详情service返回错误：(%v)", err.Error())
			return receiver.returnError("查询歌单详情service返回错误：(%v)", err.Error())
		}
		song.SongId = dbSong.SongId
		song.Singer = dbSong.Singer
		song.SongName = dbSong.SongName
		song.SongAlbum = dbSong.SongAlbum
		song.SongLyric = dbSong.SongLyric
		song.SongPlayUrl = dbSong.SongPlayUrl
		song.SongCoverUrl = dbSong.SongCoverUrl
		resp.UserSongList = append(resp.UserSongList, song)
	}

	resp.SongCoverId = songCover.ID
	resp.SongCoverName = songCover.SongCoverName
	resp.SongCoverImgUrl = songCover.CoverUrl

	receiver.Data["resp"] = resp
	receiver.TplName = "song/collectSongList.html"

	return nil
}

// @Title DeleteSong
// @Description 删除歌曲
// @Param info body models.DeleteSongReq true "req"
// @Success 200 {object} models.DeleteSongResp "resp"
// @Failure exec error
// @router /deleteSong [delete]
func (receiver *SongController) DeleteSong() error {
	receiver.BeforeStart("DeleteSong")
	var req models.DeleteSongReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("删除歌曲-解析参数错误：(%v),请求参数(%s)", err.Error(), receiver.Ctx.Input.RequestBody)
		return receiver.returnJSONError("请求参数错误:(%v)", err.Error())
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	err = songService.DeleteSong(req)
	if err != nil {
		logs.Error("删除歌曲-service返回错误(%s)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.DeleteSongResp
	return receiver.returnJSONSuccess(resp)
}

// @Title CreateSongPlayHistory
// @Description 添加歌曲播放历史
// @Param info body models.CreateSongPlayHistoryReq true "req"
// @Success 200 {object} models.CreateSongPlayHistoryResp "resp"
// @Failure exec error
// @router /createSongPlayHistory [post]
func (receiver *SongController) CreateSongPlayHistory() error {
	receiver.BeforeStart("CreateSongPlayHistory")

	var req models.CreateSongPlayHistoryReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("添加歌曲播放历史-解析参数错误(%s)", err.Error())
		return receiver.returnJSONError("解析参数错误:(%v)", err.Error())
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	err = songService.CreateSongPlayHistory(req)
	if err != nil {
		logs.Error("添加歌曲播放历史-service返回错误(%s)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.CreateSongPlayHistoryResp

	return receiver.returnJSONSuccess(resp)
}

// @Title QuerySongPlayHistoryList
// @Description 查询歌曲播放历史列表
// @Param info body models.QuerySongPlayHistoryListReq true "req"
// @Success 200 {object} models.QuerySongPlayHistoryListResp "resp"
// @Failure exec error
// @router /querySongPlayHistoryList [get]
func (receiver *SongController) QuerySongPlayHistoryList() error {
	receiver.BeforeStart("QuerySongPlayHistoryList")

	var req models.QuerySongPlayHistoryListReq

	songService := service.NewSongService(receiver.GetServiceInit())
	result, err := songService.QuerySongPlayHistoryList(req)
	if err != nil {
		logs.Error("查询歌曲播放历史列表-service返回错误(%s)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.QuerySongPlayHistoryListResp
	var song models.SongPlayHistory
	for _, v := range result {
		song.Singer = v.Singer
		song.SongId = v.SongId
		song.SongName = v.SongName
		song.SongPlayUrl = v.SongPlayUrl
		song.SongCoverUrl = v.SongCoverUrl
		resp.List = append(resp.List, song)
	}

	return receiver.returnJSONSuccess(resp)
}

// @Title DeleteAllSongPlayHistory
// @Description 删除所有歌曲播放历史
// @Success 200 {object} models.DeleteAllSongPlayHistoryResp "resp"
// @Failure exec error
// @router /deleteAllSongPlayHistory [delete]
func (receiver *SongController) DeleteAllSongPlayHistory() error {
	receiver.BeforeStart("DeleteAllSongPlayHistory")

	var req models.DeleteAllSongPlayHistoryReq
	songService := service.NewSongService(receiver.GetServiceInit())
	err := songService.DeleteAllSongPlayHistory(req)
	if err != nil {
		logs.Error("删除所有歌曲播放历史-service返回错误(%s)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.DeleteAllSongPlayHistoryResp

	return receiver.returnJSONSuccess(resp)
}

// @Title DeleteSongPlayHistory
// @Description 删除播放历史歌曲
// @Param songId path string true "歌曲Id"
// @Success 200 {object} models.DeleteSongPlayHistoryResp "resp"
// @Failure exec error
// @router /deleteSongPlayHistory/:songId [delete]
func (receiver *SongController) DeleteSongPlayHistory() error {
	receiver.BeforeStart("DeleteSongPlayHistory")

	req := models.DeleteSongPlayHistoryReq{
		SongId: receiver.GetString(":songId"),
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	err := songService.DeleteSongPlayHistory(req)
	if err != nil {
		logs.Error("删除播放历史歌曲-service返回错误(%s)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.DeleteSongPlayHistoryResp

	return receiver.returnJSONSuccess(resp)
}
