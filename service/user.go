package service

import (
	"online-music/models"
	"online-music/service/dbModel"
)

type (
	IUserService interface {
		IBaseService
		//创建用户
		CreateUser(param models.CreateUserReq) error
		//修改用户
		ModifyUser(param models.ModifyUserReq) error
		//修改密码
		ModifyPwd(param models.ModifyPwdReq) error
		//查询用户列表
		QueryBUserList(req models.QueryBUserListReq) (dbModel.BUserList, error)
		//查询用户根据ID
		QueryBUserByID(req models.QueryBUserByIDReq) (dbModel.BUserInfo, error)
	}
)

func NewUserService(init IBaseServiceInit) IUserService {
	temp := allService[ServiceIUser]
	result := temp.(IUserService)
	result.SetInitInfo(init)
	return result
}
