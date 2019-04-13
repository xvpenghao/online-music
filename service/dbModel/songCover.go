package dbModel

//歌单封面
type SongCover struct {
	//封面图片路径
	CoverImgUrl string
	//歌单描述
	Description string
	//歌单id
	SongCoverId string
}

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
