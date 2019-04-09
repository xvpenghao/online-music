package service

import (
	"github.com/jinzhu/gorm"
	"online-music/models"
)

type (
	IBaseService interface {
		//得到数据库链接
		GetConn() (*gorm.DB, error)
		//前置日志打印
		BeforeLog(msg string)
		//设置初始化信息
		SetInitInfo(init IBaseServiceInit)
	}

	IBaseServiceInit interface {
		SetBaseRequest(request models.BaseRequest)
		GetBaseRequest() models.BaseRequest
	}
)

func NewBaseServiceInit(req models.BaseRequest) IBaseServiceInit {
	temp := allService[ServiceIBaseInit]
	result := temp.(IBaseServiceInit)
	result.SetBaseRequest(req)
	return result
}
