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

type SongDb struct {
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
