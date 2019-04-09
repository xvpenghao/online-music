package impl

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"online-music/common/constants"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
	"time"
)

type UserService struct {
	BaseService
}

//创建用户
func (receiver *UserService) CreateUser(param models.CreateUserReq) error {
	receiver.BeforeLog("CreateUser")

	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据错误", err)
	}
	defer db.Close()

	//用户名和邮箱不能重复
	var counts int
	err = db.Table("tb_login").Where("login_name = ?", param.UserName).Or("login_email = ?", param.Email).
		Count(&counts).Error
	if err != nil {
		logs.Error("创建用户-根据用户名和邮箱查询用户错误:%s", err.Error())
		return utils.NewDBErr("根据用户名和邮箱查询用户错误", err)
	}

	if counts > 0 {
		logs.Error("创建用户-根据用户名和邮箱查询，用户名和邮箱不唯一:%s")
		return utils.NewSysErr("用户名和邮箱不唯一")
	}

	formStr := "2006-01-02"
	birthday, _ := time.Parse(formStr, param.Birthday)
	nowTime := time.Now()

	//封装用户信息，用户登录信息
	user := dbModel.User{
		Id:         utils.GetUUID(),
		Name:       param.UserName,
		Email:      param.Email,
		Gender:     param.Gender,
		Age:        param.Age,
		Birthday:   birthday,
		DelState:   constants.USER_NO_DEL_STATUS,
		CreatTime:  nowTime,
		UpdateTime: nowTime,
	}

	loginUser := dbModel.UserLogin{
		LoginId:    utils.GetUUID(),
		UserId:     user.Id,
		LoginName:  param.UserName,
		LoginEmail: param.Email,
		Password:   param.Password,
		CreatTime:  nowTime,
		UpdateTime: nowTime,
	}

	tx := db.Begin()
	err = tx.Table("tb_user").Create(&user).Error
	if err != nil {
		tx.Rollback()
		logs.Error("创建用户错误:%s", err.Error())
		return utils.NewDBErr("数据错误", err)
	}
	err = tx.Table("tb_login").Create(&loginUser).Error
	if err != nil {
		tx.Rollback()
		logs.Error("创建用户-创建用户登录信息失败:%s", err.Error())
		return utils.NewDBErr("创建用户登录信息失败", err)
	}
	tx.Commit()
	return nil
}

//修改用户
func (receiver *UserService) ModifyUser(param models.ModifyUserReq) error {
	receiver.BeforeLog("ModifyUser")

	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据错误", err)
	}
	defer db.Close()

	//用户名和邮箱不能重复
	var counts int
	err = db.Raw(dbModel.QUERY_USER_COUNTS_BY_UID, receiver.BaseRequest.UserID, param.UserName, param.Email).Count(&counts).Error
	if err != nil {
		logs.Error("创建用户-根据用户名和邮箱查询用户错误:%s", err.Error())
		return utils.NewDBErr("根据用户名和邮箱查询用户错误", err)
	}

	if counts > 0 {
		logs.Error("创建用户-根据用户名和邮箱查询，用户名和邮箱不唯一:%s")
		return utils.NewSysErr("用户名和邮箱不唯一")
	}

	formStr := "2006-01-02"
	birthday, _ := time.Parse(formStr, param.Birthday)
	nowTime := time.Now()

	//更新用户字段
	updateUserField := map[string]interface{}{
		"user_name":      param.UserName,
		"email":          param.Email,
		"age":            param.Age,
		"gender":         param.Gender,
		"birthday":       birthday,
		"update_time":    nowTime,
		"update_user_id": receiver.BaseRequest.UserID,
		"update_user":    receiver.BaseRequest.UserName,
	}
	//更新登录用户信息
	updateUserLoginField := map[string]interface{}{
		"login_name":     param.UserName,
		"login_email":    param.Email,
		"update_time":    nowTime,
		"update_user_id": receiver.BaseRequest.UserID,
		"update_user":    receiver.BaseRequest.UserName,
	}

	tx := db.Begin()
	err = tx.Table("tb_user").Where("user_id = ?", receiver.BaseRequest.UserID).Update(updateUserField).Error
	if err != nil {
		tx.Rollback()
		logs.Error("修改用户-更新用户信息错误:%s", err.Error())
		return utils.NewDBErr("更新用户信息错误", err)
	}
	//登录用户信息
	err = tx.Table("tb_login").Where("user_id = ?", receiver.BaseRequest.UserID).Update(updateUserLoginField).Error
	if err != nil {
		tx.Rollback()
		logs.Error("修改用户-更新登录用户信息错误:%s", err.Error())
		return utils.NewDBErr("更新登录用户信息错误", err)
	}

	tx.Commit()
	return nil
}

/*
*@Title: 修改密码
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *UserService) ModifyPwd(param models.ModifyPwdReq) error {
	receiver.BeforeLog("ModifyPwd")

	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("数据库链接错误：(%v)", err.Error())
		return utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()
	//校验用户输出的旧密码是否正确
	var counts int
	err = db.Table("tb_login").Where("user_id =?", receiver.BaseRequest.UserID).
		Where("password = ?", param.OldPwd).Count(&counts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Error("修改密码-根据旧密码查询用户失败：(%v)", err.Error())
		return utils.NewDBErr("根据旧密码查询用户失败", err)
	}
	if counts == 0 {
		logs.Error("修改密码-输入的旧密码错误：(%v)", err.Error())
		return utils.NewSysErr("输入的旧密码错误，无法修改用户信息", err)
	}

	nowTime := time.Now()

	updateField := map[string]interface{}{
		"password":       param.NewPwd,
		"update_time":    nowTime,
		"update_user_id": receiver.BaseRequest.UserID,
		"update_user":    receiver.BaseRequest.UserName,
	}

	//修改用户信息
	tx := db.Begin()
	err = tx.Table("tb_login").Where("user_id = ?", receiver.BaseRequest.UserID).Update(updateField).Error
	if err != nil {
		tx.Rollback()
		logs.Error("修改密码错误：(%v)", err.Error())
		return utils.NewSysErr("修改密码错误", err)
	}

	tx.Commit()
	return nil
}
