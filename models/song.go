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
