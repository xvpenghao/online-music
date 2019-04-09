package models

//存入redis的session
type Session struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}
