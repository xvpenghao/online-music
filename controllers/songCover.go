package controllers

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"online-music/common/constants"
	"online-music/models"
	"online-music/service"
	"online-music/service/dbModel"
	"online-music/verify"
)

type SongCoverController struct {
	BaseController
}

//@Title SongBestCoverUI
//@Description 查询歌单列表
//@Param channelId query string true "来源渠道id 如网易，qq"
//@Param curPage query int true "当前页"
//@Failure exec error
//@router /querySongCoverList [get]
func (receiver *SongCoverController) QuerySongCoverList() error {
	receiver.BeforeStart("querySongCoverList")

	channelId := receiver.GetString("channelId")
	curPage, err := receiver.GetInt("curPage")
	if err != nil {
		logs.Error("查询歌单列表-获取当前页错误：(%v)", err.Error())
		return receiver.returnError("获取当前页错误：(%v)", err.Error())
	}

	req := models.QuerySongCoverListReq{ChannelId: channelId, CurPage: curPage}
	err = verify.QuerySongCoverListReqVerify(req)
	if err != nil {
		logs.Error("查询歌单列表-参数错误：(%v)", err.Error())
		return receiver.returnError("参数错误：(%v)", err.Error())
	}

	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := songCoverService.QuerySongCoverList(req)
	if err != nil {
		logs.Error("查询歌单列表-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误：(%v)", err.Error())
	}

	var songCover models.SongCover
	var resp models.QuerySongCoverListResp
	for _, v := range result {
		songCover.SongCoverId = v.SongCoverId
		songCover.CoverImgUrl = v.CoverImgUrl
		songCover.Description = v.Description
		resp.List = append(resp.List, songCover)
	}

	//歌单列表
	receiver.Data["coverList"] = resp

	//显示歌单列表
	receiver.TplName = "song/songBestCover.html"
	return nil
}

//@Title QuerySongList
//@Description 根据歌单id获取歌曲列表
//@Param channelId query string true "渠道id"
//@Param songCoverId query string true "歌单id"
//@Param coverImgUrl query string true "歌单图片id"
//@Param description query string true "歌单描述"
//@Failure exec error
//@router /querySongList [get]
func (receiver *SongCoverController) QuerySongList() error {
	receiver.BeforeStart("SongListUI")

	req := models.QuerySongListReq{
		ChannelId:       receiver.GetString("channelId"),
		SongCoverId:     receiver.GetString("songCoverId"),
		SongCoverImgUrl: receiver.GetString("coverImgUrl"),
		Description:     receiver.GetString("description"),
	}
	err := verify.QuerySongListReqVerify(req)
	if err != nil {
		logs.Error("根据歌单id获取歌曲列表-参数错误：(%v)", err.Error())
		return receiver.returnError("参数错误：(%v)", err.Error())
	}
	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := songCoverService.QuerySongList(req)
	if err != nil {
		logs.Error("根据歌单id获取歌曲列表-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误：(%v)", err.Error())
	}

	songService := service.NewSongService(receiver.GetServiceInit())
	var querySongDetailReq models.QuerySongDetailReq
	var song models.Song
	var dbSong dbModel.Song
	var resp models.QuerySongListResp
	logs.Debug(len(result))
	for i, v := range result {
		//FIXME 这里需要优化，目前先定位3首歌曲
		if i == 3 {
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
		song.SongName = dbSong.SongName
		song.Singer = dbSong.Singer
		song.SongAlbum = dbSong.SongAlbum
		song.SongPlayUrl = dbSong.SongPlayUrl
		song.SongCoverUrl = dbSong.SongCoverUrl
		resp.List = append(resp.List, song)
	}

	resp.SongCoverId = req.SongCoverId
	resp.Description = req.Description
	resp.SongCoverImgUrl = req.SongCoverImgUrl
	receiver.Data["songList"] = resp

	receiver.TplName = "song/songList.html"
	return nil
}

//@Title CreateSongCover
//@Description 创建歌单
//@Param info body models.CreateSongCoverReq true "req"
//@Failure exec error
//@router /createSongCover [post]
func (receiver *SongCoverController) CreateSongCover() error {
	receiver.BeforeStart("CreateSongCover")

	if receiver.Session.UserId == "" {
		logs.Error("创建歌单-用户未登录不能创建歌单")
		return receiver.returnError("对不起，您未登录，不能创建歌单，请登录后操作")
	}
	var req models.CreateSongCoverReq
	err := receiver.ParseForm(&req)
	if err != nil {
		logs.Error("创建歌单-解析表单参数错误：(%v),请求参数(%+v)", err.Error(), req)
		return receiver.returnError("解析表单参数错误:(%v)", err.Error())
	}
	logs.Debug("%+v", req)

	err = verify.CreateSongCoverReqVerify(req)
	if err != nil {
		logs.Error("创建歌单-参数错误：(%v)", err.Error())
		return receiver.returnError("参数错误：(%v)", err.Error())
	}

	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	err = songCoverService.CreateSongCover(req)
	if err != nil {
		logs.Error("创建歌单-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误：(%v)", err.Error())
	}

	receiver.Redirect("/v1/index/indexUI", http.StatusFound)

	return nil
}

//@Title QueryUserSongCoverList
//@Description 查询用户歌单列表
//@Param info body models.QueryUserSongCoverListReq true "req"
//@Failure exec error
//@router /queryUserSongCoverList [get]
func (receiver *SongCoverController) QueryUserSongCoverList() error {
	receiver.BeforeStart("QueryUserSongCoverList")

	req := models.QueryUserSongCoverListReq{
		UserId: receiver.Session.UserId,
		Type:   constants.SONG_COVER_TYPE_CUSTOMER,
	}
	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := songCoverService.QueryUserSongCoverList(req)
	if err != nil {
		logs.Error("查询用户歌单列表-service返回错误：(%v)", err.Error())
		msg := map[string]string{
			"msg": err.Error(),
		}
		receiver.Data["json"] = msg
		receiver.ServeJSON()
		return nil
	}

	var resp models.QueryUserSongCoverListResp
	var userSongCover models.UserSongCover
	for _, v := range result {
		userSongCover.UserSongCoverId = v.ID
		userSongCover.SongCoverName = v.SongCoverName
		resp.UserSongCoverList = append(resp.UserSongCoverList, userSongCover)
	}

	logs.Debug("%+v", resp)

	receiver.Data["json"] = resp
	receiver.ServeJSON()

	return nil
}
