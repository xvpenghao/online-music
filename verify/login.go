package verify

import (
	"fmt"
	"online-music/common/utils"
	"online-music/models"
)

//登录校验
func LoginReqVerify(req models.LoginReq) error {

	//校验登录用户和邮箱
	if !utils.CheckLegal(&utils.StrChecker{Value: req.UserNameOrEmail, MinLen: 1, MaxLen: 30}) {
		return fmt.Errorf("用户名或者邮箱(%v)参数错误，取值(%v ~ %v)", req.UserNameOrEmail, 1, 30)
	}
	//定义密码格式的校验
	if !utils.CheckLegal(&utils.PwdCheck{Value: req.Password}) {
		return fmt.Errorf("登录密码(%v)格式不正确错误,密码必须包含数字+大小写字母", req.Password)
	}

	return nil
}
