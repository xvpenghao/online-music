package impl

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"log"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
)

type LoginService struct {
	BaseService
}

//登录
func (receiver *LoginService) DoLogin(param models.LoginReq) (dbModel.User, error) {
	var result dbModel.User
	receiver.BeforeLog("DoLogin")

	db, err := receiver.GetConn()
	if err != nil {
		log.Println("数据库链接错误")
		return result, utils.NewDBErr("数据错误", err)
	}
	defer db.Close()

	err = db.Raw(dbModel.QUERY_LOGIN_USER_INFO, param.UserNameOrEmail, param.UserNameOrEmail, param.Password).
		First(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error("登录-根据用户名或邮箱和密码查询用户错误:%s", err.Error())
		return result, utils.NewDBErr("根据用户名或邮箱和密码查询用户错误", err)
	}

	if err == gorm.ErrRecordNotFound {
		logs.Error("登录-登录账号或密码不正确")
		return result, utils.NewDBErr("登录账号或密码不正确")
	}

	return result, nil

}
