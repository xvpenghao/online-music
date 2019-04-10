package models

//查询歌单列表
type QuerySongCoverListReq struct {
	//歌单来源渠道id
	ChannelId string `json:"channelId"`
	//当前页
	CurPage int `json:"curPage"`
}

type SongCover struct {
	//封面图片路径
	CoverImgUrl string
	//歌单描述
	Description string
	//歌单id
	SongCoverId string
}

type QuerySongCoverListResp struct {
	//歌单列表
	List []SongCover
}
