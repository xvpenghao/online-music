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

//根据歌单id查询歌曲列表
type QuerySongListReq struct {
	//保留字段，可能会需要
	ChannelId string
	//歌单id
	SongCoverId string
	//歌单封面
	SongCoverImgUrl string
	//歌单描述
	Description string
}

type QuerySongListResp struct {
	//歌单id
	SongCoverId string
	//歌单url
	SongCoverImgUrl string
	//歌单描述
	Description string
	//歌曲列表
	List []Song
}
