package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"online-music/models"
	"online-music/service"
)

type BSongCoverController struct {
	BaseController
}

// @Title BSongCoverListUI
// @Description 歌单列表UI
// @Failure exec error
// @router /bSongCoverListUI [get]
func (receiver *BSongCoverController) BSongCoverListUI() error {
	receiver.BeforeStart("BSongCoverListUI")

	receiver.TplName = "admin/songCover/songCoverList.html"
	return nil
}

// @Title QueryBSongCoverByID
// @Description 查询歌单根据id
// @Param songCoverId path string true "歌单ID"
// @Param userId path string true "用户ID"
// @Failure exec error
// @router /queryBSongCoverByID/:songCoverId/:userId [get]
func (receiver *BSongCoverController) QueryBSongCoverByID() error {
	receiver.BeforeStart("QueryBSongCoverByID")
	req := models.QueryCoverSongByIdReq{
		SongCoverId: receiver.GetString(":songCoverId"),
		UserId:      receiver.GetString(":userId"),
	}
	bSongCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := bSongCoverService.QuerySongCoverById(req)
	if err != nil {
		logs.Error("查询歌单根据id-service返回错误:(%v)", err.Error())
		return receiver.returnError("service返回错误")
	}

	resp := models.QueQueryBSongCoverByIDResp{
		SongCoverId:   result.ID,
		SongCoverName: result.SongCoverName,
		UserId:        req.UserId,
	}

	receiver.Data["songCover"] = resp
	receiver.TplName = "admin/songCover/songCoverModify.html"
	return nil
}

// @Title QueryPageSongCoverList
// @Description 查询分页歌单列表
// @Param info body models.QueryPageSongCoverListReq true "req"
// @Success 200 {object} models.QueryPageSongCoverListResp "resp"
// @Failure exec error
// @router /queryPageSongCoverList [post]
func (receiver *BSongCoverController) QueryPageSongCoverList() error {
	receiver.BeforeStart("QueryPageSongCoverList")
	var req models.QueryPageSongCoverListReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("查询分页歌单列表-参数错误:(%v)", err.Error())
		return receiver.returnError("参数错误")
	}
	bSongCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	result, err := bSongCoverService.QueryPageSongCoverList(req)
	if err != nil {
		logs.Error("查询分页歌单列表-service返回错误:(%v)", err.Error())
		return receiver.returnError("service返回参数错误")
	}
	var resp models.QueryPageSongCoverListResp
	var songCover models.SongCoverInfo
	for _, v := range result.List {
		songCover.UserId = v.UserId
		songCover.SongCoverId = v.SongCoverId
		songCover.SongCoverName = v.SongCoverName
		songCover.UserName = v.UserName
		songCover.Type = v.Type
		resp.List = append(resp.List, songCover)
	}

	resp.Page = result.Page
	return receiver.returnJSONSuccess(resp)
}

// @Title ModifyBSongCover
// @Description 修改歌单信息
// @Param info body models.ModifyBSongCoverReq true "req"
// @Success 200 {object} models.resp "resp"
// @Failure exec error
// @router /modifyBSongCover [put]
func (receiver *BSongCoverController) ModifyBSongCover() error {
	receiver.BeforeStart("ModifyBSongCover")
	var req models.ModifyBSongCoverReq
	err := json.Unmarshal(receiver.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("修改歌单信息-参数错误:(%v)", err.Error())
		return receiver.returnError("参数错误")
	}

	bSongCoverService := service.NewSongCoverService(receiver.GetServiceInit())
	modifyReq := models.ModifySongCoverReq{
		SongCoverId:   req.SongCoverId,
		SongCoverName: req.SongCoverName,
		UserId:        req.UserId,
	}
	err = bSongCoverService.ModifySongCover(modifyReq)
	if err != nil {
		logs.Error("修改歌单信息-service发挥错误:(%v)", err.Error())
		return receiver.returnError("service返回错误")
	}

	var resp models.ModifyBSongCoverResp

	return receiver.returnJSONSuccess(resp)
}
