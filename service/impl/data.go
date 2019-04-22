package impl

import (
	"github.com/astaxie/beego/logs"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
)

type DataService struct {
	BaseService
}

/*
*@Title: 查询网站使用群体
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/22 0022
*@Param
*@Return
 */
func (receiver *DataService) QueryWebsiteUseGroup(req models.QueryWebsiteUseGroupReq) ([]dbModel.WebSiteUseGroup, error) {
	receiver.BeforeLog("QueryWebsiteUseGroup")
	var result []dbModel.WebSiteUseGroup
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("查询网站使用群体-数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()
	sql := dbModel.QUERY_WEBSITE_USE_GROUP
	err = db.Raw(sql).Find(&result).Error
	if err != nil {
		logs.Error("查询网站使用群体错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询网站使用群体错误：(%v)", err)
	}

	return result, nil
}

/*
*@Title: 查询性别比例
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/22 0022
*@Param
*@Return
 */
func (receiver *DataService) QueryGenderProportion(req models.QueryGenderProportionReq) ([]dbModel.GenderProportion, error) {
	receiver.BeforeLog("QueryGenderProportion")
	db, err := receiver.GetConn()
	var result []dbModel.GenderProportion
	if err != nil {
		logs.Error("数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()
	sql := dbModel.QUERY_GENDER_PROPORTION
	db.Raw(sql).Find(&result)

	return result, nil
}
