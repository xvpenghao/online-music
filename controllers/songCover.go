package controllers

import (
	"encoding/json"
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
	//我知道歌曲的数量
	for i, v := range result {
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

// @Title CreateSongCover
// @Description 创建歌单
// @Param info body models.CreateSongCoverReq true "req"
// @Success res {object} models.CreateCollectSongCoverResp true "resp"
// @Failure exec error
// @router /createSongCover [post]
func (receiver *SongCoverController) CreateSongCover() error {
	receiver.BeforeStart("CreateSongCover")

	if receiver.Session.UserId == "" {
		logs.Error("创建歌单-用户未登录不能创建歌单")
		return receiver.returnJSONError("对不起，您未登录，不能创建歌单，请登录后操作")
	}
	var req models.CreateSongCoverReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("创建歌单-解析参数错误：(%v),请求参数(%+v)", err.Error(), req)
		return receiver.returnJSONError("请求参数错误:(%v)", err.Error())
	}

	err = verify.CreateSongCoverReqVerify(req)
	if err != nil {
		logs.Error("创建歌单-参数不合法：(%v)", err.Error())
		return receiver.returnJSONError("参数不合法:(%v)", err.Error())
	}

	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := songCoverService.CreateSongCover(req)
	if err != nil {
		logs.Error("创建歌单-service返回错误：(%v)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}
	resp := models.CreateSongCoverResp{
		SongCoverId: result.SongCoverId,
	}

	return receiver.returnJSONSuccess(resp)
}

//@Title QueryUserSongCoverList
//@Description 查询用户歌单列表
//@Param info body models.QueryUserSongCoverListReq true "req"
//@Failure exec error
//@router /queryUserSongCoverList [get]
func (receiver *SongCoverController) QueryUserSongCoverList() error {
	receiver.BeforeStart("QueryUserSongCoverList")
	var resp models.QueryUserSongCoverListResp

	//用户没有登录，则直接退出
	if receiver.Session.UserId == "" {
		receiver.Data["json"] = resp
		receiver.ServeJSON()
		return nil
	}

	req := models.QueryUserSongCoverListReq{
		UserId: receiver.Session.UserId,
		Type:   constants.SONG_COVER_TYPE_CUSTOMER,
	}
	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := songCoverService.QueryUserSongCoverList(req)
	if err != nil {
		logs.Error("查询用户歌单列表-service返回错误：(%v)", err.Error())
		return receiver.returnJSONError("service返回错误：(%v)", err.Error())
	}

	var userSongCover models.UserSongCover
	for _, v := range result {
		//根据type来区分自定义歌单和收藏歌单
		userSongCover.UserSongCoverId = v.UserSongCoverId
		userSongCover.SongCoverName = v.SongCoverName
		userSongCover.SongCoverId = v.SongCoverId
		userSongCover.SongCoverUrl = v.SongCoverUrl
		if v.Type == constants.SONG_COVER_TYPE_CUSTOMER {
			resp.UserSongCoverList = append(resp.UserSongCoverList, userSongCover)
		} else {
			resp.CollectSongCoverList = append(resp.CollectSongCoverList, userSongCover)
		}
	}

	return receiver.returnJSONSuccess(resp)
}

//@Title CreateCollectSongCover
//@Description 创建收藏歌单
//@Param req body models.CreateCollectSongCoverReq true "req"
//@Success resp {object} models.CreateCollectSongCoverResp true "resp"
//@Failure exec error
//@router /createCollectSongCover [post]
func (receiver *SongCoverController) CreateCollectSongCover() error {
	receiver.BeforeStart("CreateCollectSongCover")

	var req models.CreateCollectSongCoverReq
	var resp models.CreateCollectSongCoverResp
	logs.Debug("%s", receiver.Ctx.Input.RequestBody)

	if receiver.Session.UserId == "" {
		receiver.Ctx.Output.Status = http.StatusBadRequest
		resp.Msg = "对不起，你还没有登录，请登录之后，再收藏"
		receiver.Data["json"] = resp
		receiver.ServeJSON()
		return nil
	}

	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("创建收藏歌单-解析参数错误：(%v)", err.Error())
		receiver.Ctx.Output.Status = http.StatusBadRequest
		resp.Msg = err.Error()
		receiver.Data["json"] = resp
		receiver.ServeJSON()
		return nil
	}

	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	err = songCoverService.CreateCollectSongCover(req)
	if err != nil {
		logs.Error("创建收藏歌单-service返回错误：(%v)", err.Error())
		receiver.Ctx.Output.Status = http.StatusBadRequest
		resp.Msg = err.Error()
		receiver.Data["json"] = resp
		receiver.ServeJSON()
	}

	receiver.Data["json"] = resp
	receiver.ServeJSON()
	return nil
}

// @Title UserSongCoverListUI
// @Description 用户歌单列表UI
// @Param singId path string true "歌曲id"
// @Failure exec error
// @router /userSongCoverListUI/:songId [get]
func (receiver *SongCoverController) UserSongCoverListUI() error {
	receiver.BeforeStart("UserSongCoverListUI")

	receiver.Data["songId"] = receiver.GetString(":songId")

	receiver.TplName = "song/userSongCoverList.html"
	return nil
}

// @Title ModifySongCoverUI
// @Description 修改歌单UI
// @Param songCoverId query string true "歌单id"
// @Param songCoverName query string true "歌单名称"
// @Failure exec error
// @router /modifySongCoverUI [get]
func (receiver *SongCoverController) ModifySongCoverUI() error {
	receiver.BeforeStart("ModifySongCoverUI")

	receiver.Data["songCoverId"] = receiver.GetString("songCoverId")
	receiver.Data["songCoverName"] = receiver.GetString("songCoverName")

	receiver.TplName = "song/modifySongCover.html"
	return nil
}

// @Title ModifySongCover
// @Description 编辑歌单
// @Param info body models.ModifySongCoverReq true "req"
// @Success 200 {object} models.ModifySongCoverResp "resp"
// @Failure exec error
// @router /modifySongCover [post]
func (receiver *SongCoverController) ModifySongCover() error {
	receiver.BeforeStart("ModifySongCover")

	if receiver.Session.UserId == "" {
		logs.Error("编辑歌单-用户未登录不能创建歌单")
		return receiver.returnJSONError("对不起，您未登录，不能创建歌单，请登录后操作")
	}

	var req models.ModifySongCoverReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("编辑歌单-解析参数错误：(%v),请求参数(%s)", err.Error(), receiver.Ctx.Input.RequestBody)
		return receiver.returnJSONError("请求参数错误:(%v)", err.Error())
	}

	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	err = songCoverService.ModifySongCover(req)
	if err != nil {
		logs.Error("编辑歌单-service返回错误:(%v)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.ModifySongCoverResp

	return receiver.returnJSONSuccess(resp)
}

// @Title DeleteSongCover
// @Description 删除歌单
// @Param songCoverId path string true "req"
// @Success 200 {object} models.DeleteSongCoverResp "resp"
// @Failure exec error
// @router /deleteSongCover/:songCoverId [delete]
func (receiver *SongCoverController) DeleteSongCover() error {
	receiver.BeforeStart("DeleteSongCover")

	req := models.DeleteSongCoverReq{
		SongCoverId: receiver.GetString(":songCoverId"),
	}

	songCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	err := songCoverService.DeleteSongCover(req)
	if err != nil {
		logs.Error("删除歌单-service返回错误:(%v)", err.Error())
		return receiver.returnJSONError("service返回错误:(%v)", err.Error())
	}

	var resp models.DeleteSongCoverResp
	return receiver.returnJSONSuccess(resp)
}
