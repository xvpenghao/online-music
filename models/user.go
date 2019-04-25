package models

//创建用户
type CreateUserReq struct {
	UserName string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Age      int    `form:"age"`
	Birthday string `form:"birthday"`
	Gender   string `form:"gender"`
}

//修改用户信息
type ModifyUserReq struct {
	UserName string `form:"username"`
	Email    string `form:"email"`
	Age      int    `form:"age"`
	Birthday string `form:"birthday"`
	Gender   string `form:"gender"`
	UserId   string `form:"-"`
}

//修改用户密码
type ModifyPwdReq struct {
	OldPwd    string `form:"oldPwd"`
	NewPwd    string `form:"newPwd"`
	BeSurePwd string `form:"beSurePwd"`
}
