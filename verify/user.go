package verify

import (
	"fmt"
	"online-music/common/utils"
	"online-music/models"
)

func CreateUserVerify(req models.CreateUserReq) error {
	if !utils.CheckLegal(&utils.StrChecker{Value: req.UserName, MinLen: 1, MaxLen: 16}) {
		return fmt.Errorf("用户名(%v)参数错误，取值(%v ~ %v)", req.UserName, 1, 16)
	}
	if !utils.CheckLegal(&utils.EmailCheck{Value: req.Email}) {
		return fmt.Errorf("邮箱(%v)参数错误", req.Email)
	}
	if !utils.CheckLegal(&utils.IntChecker{Value: req.Age, Min: 1, Max: 150}) {
		return fmt.Errorf("年龄(%v)参数错误，取值(%v ! %v)", req.Age, 1, 150)
	}
	if !utils.CheckLegal(&utils.StrChecker{Value: req.Password, MinLen: 6, MaxLen: 16}) {
		return fmt.Errorf("密码(%v)参数错误，取值(%v ~ %v)", req.UserName, 6, 16)
	}
	//定义密码格式的校验
	if !utils.CheckLegal(&utils.PwdCheck{Value: req.Password}) {
		return fmt.Errorf("密码(%v)格式不正确错误,密码必须包含数字+大小写字母", req.Password)
	}

	return nil
}

//修改用户验证
func ModifyUserReqVerify(req models.ModifyUserReq) error {
	if !utils.CheckLegal(&utils.StrChecker{Value: req.UserName, MinLen: 1, MaxLen: 16}) {
		return fmt.Errorf("用户名(%v)参数错误，取值(%v ~ %v)", req.UserName, 1, 16)
	}
	if !utils.CheckLegal(&utils.EmailCheck{Value: req.Email}) {
		return fmt.Errorf("邮箱(%v)参数错误", req.Email)
	}
	if !utils.CheckLegal(&utils.IntChecker{Value: req.Age, Min: 1, Max: 150}) {
		return fmt.Errorf("年龄(%v)参数错误，取值(%v ! %v)", req.Age, 1, 150)
	}
	return nil
}

//修改密码参数验证
func ModifyPwdReqVerify(req models.ModifyPwdReq) error {
	pwds := []string{req.OldPwd, req.NewPwd, req.OldPwd}

	for _, v := range pwds {
		if !utils.CheckLegal(&utils.StrChecker{Value: v, MinLen: 6, MaxLen: 16}) {
			return fmt.Errorf("密码(%v)参数错误，取值(%v ~ %v)", v, 6, 16)
		}
	}

	if !utils.CheckLegal(&utils.StrEqualsChecker{Value: req.NewPwd, SecondValue: req.BeSurePwd}) {
		return fmt.Errorf("新密码(%v)和确认码(%v)不相同", req.NewPwd, req.BeSurePwd)
	}

	//校验格式
	if !utils.CheckLegal(&utils.PwdCheck{Value: req.OldPwd}) {
		return fmt.Errorf("旧(%v)格式不正确错误,密码必须包含数字+大小写字母", req.OldPwd)
	}
	if !utils.CheckLegal(&utils.PwdCheck{Value: req.NewPwd}) {
		return fmt.Errorf("新密码(%v)格式不正确错误,密码必须包含数字+大小写字母", req.NewPwd)
	}
	if !utils.CheckLegal(&utils.PwdCheck{Value: req.BeSurePwd}) {
		return fmt.Errorf("确认码(%v)格式不正确错误,密码必须包含数字+大小写字母", req.BeSurePwd)
	}

	return nil

}
