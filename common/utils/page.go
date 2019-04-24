package utils

type Page struct {
	//总数
	TotalCount int `json:"count"`
	//组数
	Groups int `json:"groups"`
	//每页的个数
	Limit int `json:"limit"`
	//当前页
	CurPage int `json:"curPage"`
}

func CalPageCount(curPage, limit int) int {
	if curPage < 1 {
		curPage = 0
	}
	return (curPage - 1) * limit
}
