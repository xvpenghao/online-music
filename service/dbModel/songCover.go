package dbModel

import "time"

const (
	//查询用户歌单数量，通过歌单名称
	QUERY_USER_COVER_COUNT_BY_SONG_COVER_NAME = `SELECT COUNT(*) 
                                                 FROM TB_SONG_COVER AS TBSC 
                                                 INNER JOIN TB_USER_SONG_COVER AS TBUSC 
                                                 ON TBSC.SONG_COVER_ID = TBUSC.SONG_COVER_ID
                                                 WHERE TBUSC.USER_ID = ? AND TBSC.SONG_COVER_NAME = ? AND TBSC.TYPE = ? `
	//查询用户歌单列表
	QUERY_USER_COVER_LIST = `SELECT TBUSC.USER_SONG_COVER_ID,TBSC.SONG_COVER_NAME
                             FROM TB_SONG_COVER TBSC 
                             INNER JOIN TB_USER_SONG_COVER TBUSC 
                             ON TBSC.SONG_COVER_ID = TBUSC.SONG_COVER_ID
                             AND TBUSC.USER_ID = ? AND TBUSC.DEL_STATUS = ? AND TBSC.TYPE = ? `
)

type SongCoverInfo struct {
	//歌单ID
	ID string `gorm:"column:song_cover_id"`
	//歌单的类型 1自定义 2other
	Type int `gorm:"column:type"`
	//渠道ID
	ChannelID string `gorm:"column:channel_id"`
	//歌单名称
	SongCoverName string `gorm:"column:song_cover_name"`
	//封面Url
	CoverUrl string `gorm:"column:cover_url"`
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

type UserSongCover struct {
	//用户歌单ID
	ID string `gorm:"column:user_song_cover_id"`
	//用户ID
	UserId string `gorm:"column:user_id"`
	//歌单ID
	SongCoverId string `gorm:"column:song_cover_id"`
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

type QueryUserSongCover struct {
	//用户id
	ID string `gorm:"column:USER_SONG_COVER_ID"`
	//歌单名称
	SongCoverName string `gorm:"column:SONG_COVER_NAME"`
}

//创建歌单的返回值
type CreateSongCoverReply struct {
	//歌单id
	SongCoverId string
	//歌单名称
	SongCoverName string
}
