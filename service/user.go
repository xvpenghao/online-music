package service

import (
	"online-music/models"
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
		//查询用户信息根据用户id
		/*QueryUserByUID()dbModel.User*/
	}
)

func NewUserService(init IBaseServiceInit) IUserService {
	temp := allService[ServiceIUser]
	result := temp.(IUserService)
	result.SetInitInfo(init)
	return result
}
