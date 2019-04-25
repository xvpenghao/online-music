package models

import "online-music/common/utils"

type QueryPageSongCoverListReq struct {
	//当前页
	CurPage int `json:"curPage"`
	//用户名
	UserName string `json:"userName"`
	//歌单名
	SongCoverName string `json:"songCoverName"`
	//类型
	Type int `json:"type"`
}

type SongCoverInfo struct {
	//用户Id
	UserId string `json:"userId"`
	//歌单ID
	SongCoverId string `json:"songCoverId"`
	//用户名
	UserName string `json:"userName"`
	//歌单名
	SongCoverName string `json:"songCoverName"`
	//类型
	Type string `json:"type"`
}

type QueryPageSongCoverListResp struct {
	//分页信息
	Page utils.Page `json:"page"`
	//歌单列表
	List []SongCoverInfo `json:"list"`
}

type QueQueryBSongCoverByIDResp struct {
	//歌单ID
	SongCoverId string
	//歌单名称
	SongCoverName string
	//用户ID
	UserId string
}

type ModifyBSongCoverReq struct {
	//歌单ID
	SongCoverId string `json:"songCoverId"`
	//用户Id
	UserId string `json:"userId"`
	//歌单名称
	SongCoverName string `json:"songCoverName"`
}

type ModifyBSongCoverResp struct {
	baseResp
}
