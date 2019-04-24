package impl

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"online-music/common/constants"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
	"time"
)

type ChannelService struct {
	BaseService
}

/*
*@Title:添加渠道
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/23 0023
*@Param
*@Return
 */
func (receiver *ChannelService) CreateChannel(req models.CreateChannelReq) error {
	receiver.BeforeLog("CreateChannel")
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("添加渠道-数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	var counts int
	//渠道名称不能重复
	err = db.Table("tb_channel").Where("channel_name = ?", req.ChannelName).Count(&counts).Error
	if err != nil {
		logs.Error("添加渠道-根据名称查询渠道个数失败：(%v)", err.Error())
		return utils.NewDBErr("根据名称查询渠道个数失败", err)
	}
	if counts > 0 {
		logs.Error("添加渠道-渠道名称不唯一")
		return utils.NewDBErr("渠道名称不唯一")
	}
	nowTime := time.Now()
	channel := dbModel.ChannelTable{
		ChannelId:    utils.GetUUID(),
		ChannelName:  req.ChannelName,
		CreatTime:    nowTime,
		CreateUser:   receiver.BaseRequest.UserName,
		CreateUserId: receiver.BaseRequest.UserID,
		UpdateTime:   nowTime,
		UpdateUser:   receiver.BaseRequest.UserName,
		UpdateUserId: receiver.BaseRequest.UserID,
	}

	tx := db.Begin()
	err = tx.Table("tb_channel").Create(&channel).Error
	if err != nil {
		tx.Rollback()
		logs.Error("添加渠道错误：(%v)", err.Error())
		return utils.NewDBErr("添加渠道错误", err)
	}
	tx.Commit()
	return nil
}

/*
*@Title:查询渠道详情
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/23 0023
*@Param
*@Return
 */
func (receiver *ChannelService) QueryChannelDetail(req models.QueryChannelDetailReq) (dbModel.ChannelDetail, error) {
	receiver.BeforeLog("QueryChannelDetail")
	var result dbModel.ChannelDetail
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("查询渠道详情-数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	sql := dbModel.QUERY_CHANNEL_DETAIL
	sqlParam := []interface{}{req.ChannelId}

	err = db.Raw(sql, sqlParam).First(&result).Error
	if err != nil {
		logs.Error("查询渠道详情错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询渠道详情错误", err)
	}
	return result, nil
}

/*
*@Title: 修改渠道信息
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/23 0023
*@Param
*@Return
 */
func (receiver *ChannelService) ModifyChannel(req models.ModifyChannelReq) error {
	receiver.BeforeLog("ModifyChannel")
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("修改渠道信息-数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()
	var counts int
	//渠道名称不能重复
	err = db.Table("tb_channel").Where("channel_name = ?", req.ChannelName).Count(&counts).Error
	if err != nil {
		logs.Error("修改渠道信息-根据名称查询渠道个数失败：(%v)", err.Error())
		return utils.NewDBErr("根据名称查询渠道个数失败", err)
	}
	if counts > 0 {
		logs.Error("修改渠道信息-渠道名称不唯一")
		return utils.NewDBErr("渠道名称不唯一", err)
	}
	nowTime := time.Now()
	updateField := map[string]interface{}{
		"channel_name":   req.ChannelName,
		"update_time":    nowTime,
		"update_user_id": receiver.BaseRequest.UserID,
		"update_user":    receiver.BaseRequest.UserName,
	}

	tx := db.Begin()
	err = tx.Table("tb_channel").Where("channel_id = ?", req.ChannelId).Update(updateField).Error
	if err != nil {
		tx.Rollback()
		logs.Error("修改渠道信息失败：(%v)", err.Error())
		return utils.NewDBErr("修改渠道信息失败", err)
	}
	tx.Commit()

	return nil
}

/*
*@Title: 查询渠道列表
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/24 0024
*@Param
*@Return
 */
func (receiver *ChannelService) QueryChannelList(req models.QueryChannelListReq) (dbModel.ChannelInfoList, error) {
	receiver.BeforeLog("QueryChannelList")
	var result dbModel.ChannelInfoList
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("查询渠道列表-数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	page := utils.Page{
		CurPage: req.CurPage,
		Groups:  constants.PAGE_DEFAULT_GROUPS,
		Limit:   constants.PAGE_DEFAULT_LIMIT,
	}

	//查询总记录数
	sql2 := dbModel.QUERY_CHANNEL_COUNTS
	var counts int
	err = db.Raw(sql2, req.ChannelName).Count(&counts).Error
	if err != nil {
		logs.Error("查询渠道列表-查询列表个数错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询列表个数错误", err)
	}

	//设置总页数
	page.TotalCount = counts

	var channels []dbModel.ChannelInfo
	//数据库分页
	dbPage := utils.CalPageCount(req.CurPage, constants.PAGE_DEFAULT_LIMIT)
	sql := dbModel.QUERY_CHANNEL_LIST
	sqlParam := []interface{}{req.ChannelName, page.Limit, dbPage}
	err = db.Raw(sql, sqlParam...).Find(&channels).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error("查询渠道列表错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询渠道列表错误", err)
	}

	result.List = channels
	result.Page = page
	return result, nil
}
