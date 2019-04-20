package dbModel

import "time"

const (
	//根据歌单id查询用户歌曲列表
	QUERY_USER_SONG_LIST = `select tbs.song_id,tbs.song_name,tbs.song_cover_url,tbs.song_play_url,
                                   tbs.singer,tbs.song_lyric,song_album
                           from tb_song_cover_song as tbscs 
                           INNER JOIN tb_song  tbs on tbscs.song_id = tbs.song_id 
                           where song_cover_id = ? `
)

type SongTable struct {
	//歌曲ID
	SongId string `gorm:"column:song_id"`
	//歌曲名称
	SongName string `gorm:"column:song_name"`
	//歌手
	Singer string `gorm:"column:singer"`
	//歌曲专辑
	SongAlbum string `gorm:"column:song_album"`
	//歌曲封面
	SongCoverUrl string `gorm:"column:song_cover_url"`
	//歌曲播放url
	SongPlayUrl string `gorm:"column:song_play_url"`
	//歌曲歌词
	SongLyric string `gorm:"column:song_lyric"`
	//渠道id
	ChannelCd string `gorm:"column:channel_id"`
	//删除的状态 1删除 2不删除
	DelState int `gorm:"column:del_status"`
	//创建时间
	CreatTime time.Time `gorm:"column:create_time"`
	//创建人
	CreateUser string `gorm:"column:create_user"`
	//创建人ID
	CreateUserId string `gorm:"column:create_user_id"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time"`
	//更新人
	UpdateUser string `gorm:"column:update_user"`
	//更新人ID
	UpdateUserId string `gorm:"column:update_user_id"`
}

//歌单歌曲表
type SongCoverSongTable struct {
	//歌单歌曲id
	SongCoverSongId string `gorm:"column:song_cover_song_id"`
	//歌单id
	SongCoverId string `gorm:"column:song_cover_id"`
	//用户id
	UserId string `gorm:"column:user_id"`
	//歌曲id
	SongId string `gorm:"column:song_id"`
	//删除的状态 1删除 2不删除
	DelState int `gorm:"column:del_status"`
	//创建时间
	CreatTime time.Time `gorm:"column:create_time"`
	//创建人
	CreateUser string `gorm:"column:create_user"`
	//创建人ID
	CreateUserId string `gorm:"column:create_user_id"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time"`
	//更新人
	UpdateUser string `gorm:"column:update_user"`
	//更新人ID
	UpdateUserId string `gorm:"column:update_user_id"`
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

//用户歌曲
type UserSong struct {
	//歌曲ID
	SongId string `gorm:"column:song_id"`
	//歌曲名称
	SongName string `gorm:"column:song_name"`
	//歌曲专辑
	SongAlbum string `gorm:"column:song_album"`
	//歌手
	Singer string `gorm:"column:singer"`
	//歌曲封面
	SongCoverUrl string `gorm:"column:song_cover_url"`
	//歌曲播放url
	SongPlayUrl string `gorm:"column:song_play_url"`
	//歌曲歌词
	SongLyric string `gorm:"column:song_lyric"`
}

//歌曲播放历史
type SongPlayHistory struct {
	//歌曲ID
	SongId string `json:"songId"`
	//歌曲名称
	SongName string `json:"songName"`
	//歌曲播放Url
	SongPlayUrl string `json:"playUrl"`
	//歌曲封面URL
	SongCoverUrl string `json:"songCoverUrl"`
	//歌手
	Singer string `json:"singer"`
}

//搜索歌曲结构体
type SearchSong struct {
	SongId string `json:"id"`
	//歌曲名称
	SongName string `json:"name"`
	//歌手
	Singer string `json:"singer"`
	//歌曲封面Url
	SongCoverUrl string `json:"pic"`
	//歌曲播放链接
	SongPlayUrl string `json:"url"`
	//歌词
	SongLyric string `json:"lrc"`
	//播放时间
	PlayTime int `json:"time"`
}

type SongSearchResult struct {
	Songs  []SearchSong `json:"data"`
	Code   int          `json:"code"`
	Result string       `json:"result"`
}
