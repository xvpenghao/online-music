package models

//查询网站使用群体
type QueryWebsiteUseGroupReq struct {
}

type WebsiteUseGroup struct {
	//群体名称
	GroupName string `json:"name"`
	//群体数量
	GroupCounts int `json:"value"`
}

type QueryWebsiteUseGroupResp struct {
	List []WebsiteUseGroup `json:"list"`
}

//查询性别比例
type QueryGenderProportionReq struct {
}

//查询性别比例
type GenderProportion struct {
	//性别
	Gender string `json:"name"`
	//性别分类数量
	GenderCounts float64 `json:"value"`
}
type QueryGenderProportionResp struct {
	List []GenderProportion `json:"list"`
}
