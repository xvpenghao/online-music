package models

//歌曲信息
type Song struct {
	//歌曲id
	SongId string
	//歌曲名称
	SongName string
	//歌手
	Singer string
	//歌曲专辑
	SongAlbum string
	//歌曲封面Url
	SongCoverUrl string
	//歌曲播放链接
	SongPlayUrl string
	//歌词
	SongLyric string
}

//查询歌曲详情
type QuerySongDetailReq struct {
	//歌曲id
	SongId string
}

type QuerySongDetailResp struct {
	Song
}
