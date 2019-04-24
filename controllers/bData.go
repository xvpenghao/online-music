package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"online-music/models"
	"online-music/service"
)

type DataController struct {
	BaseController
}

//@Title QueryWebsiteUseGroup
//@Description 查询网站使用群体
//@Failure exec error
//@router /queryWebsiteUseGroup [get]
func (receiver *DataController) QueryWebsiteUseGroup() error {
	receiver.BeforeStart("QueryWebsiteUserGroup")
	var req models.QueryWebsiteUseGroupReq

	dataService := service.NewDataService(receiver.GetServiceInit())
	result, err := dataService.QueryWebsiteUseGroup(req)
	if err != nil {
		logs.Error("查询网站使用群体-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误：(%v)", err.Error())
	}

	var resp models.QueryWebsiteUseGroupResp
	var websiteUseGroup models.WebsiteUseGroup
	for _, v := range result {
		websiteUseGroup.GroupName = v.GroupName
		websiteUseGroup.GroupCounts = v.GroupCounts
		resp.List = append(resp.List, websiteUseGroup)
	}

	bytes, _ := json.Marshal(resp)
	receiver.Data["resp"] = string(bytes)

	receiver.TplName = "admin/data/websiteUseGroup.html"
	return nil
}

//@Title QueryGenderProportion
//@Description 查询性别比例
//@Failure exec error
//@router /queryGenderProportion [get]
func (receiver *DataController) QueryGenderProportion() error {
	receiver.BeforeStart("QueryGenderCount")
	var req models.QueryGenderProportionReq

	dataService := service.NewDataService(receiver.GetServiceInit())
	result, err := dataService.QueryGenderProportion(req)
	if err != nil {
		logs.Error("查询性别比例-service返回错误：(%v)", err.Error())
		return receiver.returnError("service返回错误：(%v)", err.Error())
	}
	var resp models.QueryGenderProportionResp
	var genderInfo models.GenderProportion

	for _, v := range result {
		genderInfo.Gender = v.Gender
		genderInfo.GenderCounts = v.GenderCounts
		resp.List = append(resp.List, genderInfo)
	}

	bytes, _ := json.Marshal(resp)
	logs.Debug("%v", string(bytes))

	receiver.Data["resp"] = string(bytes)
	receiver.TplName = "admin/data/genderProportion.html"
	return nil
}
