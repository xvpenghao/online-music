package models

import (
	"online-music/common/utils"
)

//修改用户信息
type ModifyBUserReq struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}

type ModifyBUserResp struct {
	baseResp
}

//查询后台用户列表
type QueryBUserListReq struct {
	//当前页
	CurPage int `json:"curPage"`
	//用户名
	UserName string `json:"userName"`
	//邮箱
	Email string `json:"email"`
	//年龄
	Age int `json:"age"`
	//生日
	Birthday string `json:"birthday"`
	//性别
	Gender string `json:"gender"`
}

type BUserInfo struct {
	UserId string `json:"userId"`
	//用户名
	UserName string `json:"userName"`
	//邮箱
	Email string `json:"email"`
	//年龄
	Age int `json:"age"`
	//生日
	Birthday string `json:"birthday"`
	//性别
	Gender string `json:"gender"`
}

type QueryBUserListResp struct {
	//分页
	Page utils.Page `json:"page"`
	//用户列表
	List []BUserInfo `json:"list"`
}

//查询用户根据ID
type QueryBUserByIDReq struct {
	//用户ID
	UserId string
}

type QueryBUserByIDResp struct {
	UserId string
	//用户名
	UserName string
	//邮箱
	Email string
	//年龄
	Age int
	//生日
	Birthday string
	//性别
	Gender string
}
