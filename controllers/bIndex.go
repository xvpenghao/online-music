package controllers

import "time"

type AdminIndexController struct {
	BaseController
}

//@Title IndexUI
//@Description 后台首页UI
//@Failure exec error
//@router / [get]
func (receiver *AdminIndexController) IndexUI() error {
	receiver.TplName = "admin/index.html"
	return nil
}

//@Title WelcomeUI
//@Description 欢迎页
//@Failure exec error
//@router /welcomeUI [get]
func (receiver *AdminIndexController) WelcomeUI() error {

	formStr := "2006-01-02 15:04"
	receiver.Data["nowTime"] = time.Now().Format(formStr)

	receiver.TplName = "admin/welcome.html"
	return nil
}
