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
	//歌单名称
	Description string
	//歌曲列表
	List []Song
}

//创建歌单
type CreateSongCoverReq struct {
	//歌单名称
	SongCoverName string `json:"songCoverName"`
}

type CreateSongCoverResp struct {
	//歌单id
	SongCoverId string `json:"songCoverId"`
	//返回的消息
	baseResp
}

//查询用户歌单列表
type QueryUserSongCoverListReq struct {
	//1 自定义歌单 2other歌单
	Type int
	//用户id
	UserId string
}

type UserSongCover struct {
	//用户歌单Id
	UserSongCoverId string `json:"userSongCoverId"`
	//歌曲id
	SongCoverId string `json:"songCoverId"`
	//歌单封面url
	SongCoverUrl string `json:"songCoverUrl"`
	//用户歌单名称
	SongCoverName string `json:"songCoverName"`
}

type QueryUserSongCoverListResp struct {
	//用户歌单列表
	UserSongCoverList []UserSongCover `json:"songCoverList"`
	//用户收藏歌单列表
	CollectSongCoverList []UserSongCover `json:"collectList"`
}

//创建收藏歌单
type CreateCollectSongCoverReq struct {
	//歌单ID
	SongCoverId string `json:"songCoverId"`
	//歌单名称
	SongCoverName string `json:"songCoverName"`
	//歌单封面的url
	SongCoverUrl string `json:"songCoverUrl"`
}

type CreateCollectSongCoverResp struct {
	Msg string `json:"msg"`
}

//修改歌单
type ModifySongCoverReq struct {
	//歌单Id
	SongCoverId string `json:"songCoverId"`
	//歌单名称
	SongCoverName string `json:"songCoverName"`
	//歌单用户
	UserId string `json:"-"`
}

type ModifySongCoverResp struct {
	baseResp
}

//删除歌单
type DeleteSongCoverReq struct {
	//歌单id
	SongCoverId string `json:"songCoverId"`
}

type DeleteSongCoverResp struct {
	baseResp
}

//根据id查询歌单信息
type QueryCoverSongByIdReq struct {
	//歌单id
	SongCoverId string
	//用户id
	UserId string
}
