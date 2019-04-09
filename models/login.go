package models

//登录
type LoginReq struct {
	UserNameOrEmail string `form:"usernameOrEmail"`
	Password        string `form:"password"`
}

type LoginResp struct {
	Id string
	//用户姓名
	Name string
	//邮箱
	Email string
	//性别
	Gender string
	//年龄
	Age int
	//生日
	Birthday string
}
