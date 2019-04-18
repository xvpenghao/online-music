package models

//歌曲信息
type Song struct {
	//歌曲id
	SongId string `json:"songId"`
	//歌曲名称
	SongName string `json:"songName"`
	//歌手
	Singer string `json:"singer"`
	//歌曲专辑
	SongAlbum string `json:"songAlbum"`
	//歌曲封面Url
	SongCoverUrl string `json:"songCoverUrl"`
	//歌曲播放链接
	SongPlayUrl string `json:"songPlayUrl"`
	//歌词
	SongLyric string `json:"songLyric"`
}

//查询歌曲详情
type QuerySongDetailReq struct {
	//歌曲id
	SongId string `json:"songId"`
}

type QuerySongDetailResp struct {
	Song
}

//查询自定义歌曲列表
type QueryUserSongListReq struct {
	//歌曲id
	SongCoverId string
}

type QueryUserSongListResp struct {
	//歌单id
	SongCoverId string
	//歌单url
	SongCoverImgUrl string
	//歌单名称
	SongCoverName string
	//用户歌曲列表
	UserSongList []Song
}

//添加歌曲
type CreateSongReq struct {
	//歌曲id
	SongId string `json:"songId"`
	//歌单id
	SongCoverId string `json:"songCoverId"`
	//歌曲信息 忽略
	SongInfo Song `json:"-"`
}

type CreateSongResp struct {
	baseResp
}

//删除歌曲
type DeleteSongReq struct {
	//歌曲id
	SongId string `json:"songId"`
	//歌单id
	SongCoverId string `json:"songCoverId"`
}

type DeleteSongResp struct {
	baseResp
}
