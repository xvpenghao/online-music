package impl

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"online-music/common/constants"
	"online-music/common/utils"
	"online-music/models"
	"online-music/service/dbModel"
	"strings"
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
	err = db.Raw(dbModel.QUERY_USER_COUNTS_BY_UID, param.UserId, param.UserName, param.Email).Count(&counts).Error
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
	err = tx.Table("tb_user").Where("user_id = ?", param.UserId).Update(updateUserField).Error
	if err != nil {
		tx.Rollback()
		logs.Error("修改用户-更新用户信息错误:%s", err.Error())
		return utils.NewDBErr("更新用户信息错误", err)
	}
	//登录用户信息
	err = tx.Table("tb_login").Where("user_id = ?", param.UserId).Update(updateUserLoginField).Error
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

/*
*@Title: 查询用户列表
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/24 0024
*@Param
*@Return
 */
func (receiver *UserService) QueryBUserList(req models.QueryBUserListReq) (dbModel.BUserList, error) {
	receiver.BeforeLog("QueryBUserList")

	var result dbModel.BUserList
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("查询用户列表-数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	var whereSql strings.Builder
	sqlParam := []interface{}{}

	if req.UserName != "" {
		whereSql.WriteString(" and instr(user_name,?) ")
		sqlParam = append(sqlParam, req.UserName)
	}
	if req.Email != "" {
		whereSql.WriteString(" and instr(email,?) ")
		sqlParam = append(sqlParam, req.Email)
	}
	if req.Gender != "" {
		whereSql.WriteString(" and gender = ? ")
		sqlParam = append(sqlParam, req.Gender)
	}
	if req.Age > 0 {
		whereSql.WriteString(" age = ? ")
		sqlParam = append(sqlParam, req.Age)
	}
	if req.Birthday != "" {
		strs := strings.Split(req.Birthday, " - ")
		whereSql.WriteString(" and birthday between ? and ? ")
		sqlParam = append(sqlParam, strs[0], strs[1])
	}

	//查询总记录数
	var totalCounts int
	var queryListSqlCounts strings.Builder
	queryListSqlCounts.WriteString(dbModel.QUERY_USER_LIST_COUNTS)
	queryListSqlCounts.WriteString(whereSql.String())
	err = db.Raw(queryListSqlCounts.String(), sqlParam...).Count(&totalCounts).Error
	if err != nil {
		logs.Error("查询用户列表-查询用户列表总数错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询用户列表总数错误", err)
	}

	page := utils.Page{
		CurPage:    req.CurPage,
		TotalCount: totalCounts,
		Groups:     constants.PAGE_DEFAULT_GROUPS,
		Limit:      constants.PAGE_DEFAULT_LIMIT,
	}
	curPage := utils.CalPageCount(req.CurPage, page.Limit)

	//查询用户列表
	var queryListSql strings.Builder
	queryListSql.WriteString(dbModel.QUERY_USER_LIST)
	queryListSql.WriteString(whereSql.String())

	queryListSql.WriteString(" order by update_time ")
	queryListSql.WriteString(" limit ? offset ? ")
	sqlParam = append(sqlParam, page.Limit, curPage)

	var bUsers []dbModel.BUserInfo
	err = db.Raw(queryListSql.String(), sqlParam...).Find(&bUsers).Error
	if err != nil {
		logs.Error("查询用户列表错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询用户列表错误", err)
	}

	result.List = bUsers
	result.Page = page
	return result, nil
}

/*
*@Title:查询用户根据ID
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/25 0025
*@Param
*@Return
 */
func (receiver *UserService) QueryBUserByID(req models.QueryBUserByIDReq) (dbModel.BUserInfo, error) {
	receiver.BeforeLog("QueryBUserByID")

	var result dbModel.BUserInfo
	db, err := receiver.GetConn()
	if err != nil {
		logs.Error("查询用户根据ID-数据库链接错误：(%v)", err.Error())
		return result, utils.NewDBErr("数据库链接错误", err)
	}
	defer db.Close()

	err = db.Table("tb_user").Where("user_id = ?", req.UserId).First(&result).Error
	if err != nil {
		logs.Error("查询用户根据ID错误：(%v)", err.Error())
		return result, utils.NewDBErr("查询用户根据ID错误", err)
	}

	return result, nil
}
