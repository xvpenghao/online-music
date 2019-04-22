package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	IDataService interface {
		IBaseService
		//查询网站使用群体
		QueryWebsiteUseGroup(req models.QueryWebsiteUseGroupReq) ([]dbModel.WebSiteUseGroup, error)
		//查询性别比例
		QueryGenderProportion(req models.QueryGenderProportionReq) ([]dbModel.GenderProportion, error)
	}
)

func NewDataService(init IBaseServiceInit) IDataService {
	temp := allService[ServiceIData]
	result := temp.(IDataService)
	result.SetInitInfo(init)
	return result
}
