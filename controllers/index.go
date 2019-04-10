package controllers

type IndexController struct {
	BaseController
}

//@Title IndexUI
//@Description 首页页面
//@Failure exec error
//@router /indexUI [get]
func (receiver *IndexController) IndexUI() {
	receiver.TplName = "index.html"
}

//@Title ErrorUI
//@Description 错误页面
//@Failure exec error
//@router /errorUI [get]
func (receiver *IndexController) ErrorUI() {

	receiver.TplName = "errors.html"
}
