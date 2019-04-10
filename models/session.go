package models

//存入redis的session
type Session struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	//邮箱
	Email string `json:"email"`
	//性别
	Gender string `json:"gender"`
	//年龄
	Age int `json:"age"`
	//生日
	Birthday string `json:"birthday"`
}
