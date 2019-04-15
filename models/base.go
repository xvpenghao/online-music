package models

type BaseRequest struct {
	UserID   string
	UserName string
}

type baseResp struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}
