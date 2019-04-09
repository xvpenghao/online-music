package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	ILoginService interface {
		IBaseService
		//登录
		DoLogin(param models.LoginReq) (dbModel.User, error)
		//退出登录
	}
)

func NewLoginService(init IBaseServiceInit) ILoginService {
	temp := allService[ServiceILogin]
	result := temp.(ILoginService)
	result.SetInitInfo(init)
	return result
}
