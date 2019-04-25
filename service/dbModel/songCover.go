package dbModel

import (
	"online-music/common/utils"
	"time"
)

const (
	//查询用户歌单数量，通过歌单名称
	QUERY_USER_COVER_COUNT_BY_SONG_COVER_NAME = `select count(*) 
                                                 from tb_song_cover as tbsc 
                                                 inner join tb_user_song_cover as tbusc 
                                                 on tbsc.song_cover_id = tbusc.song_cover_id
                                                 where tbusc.user_id = ? and tbsc.song_cover_name = ? and tbsc.type = ? `
	//查询用户歌单列表
	QUERY_USER_COVER_LIST = `select tbusc.song_cover_id,tbusc.user_song_cover_id,tbsc.song_cover_name,tbsc.cover_url,tbsc.type 
                             from tb_song_cover tbsc
	                         inner join tb_user_song_cover tbusc on tbsc.song_cover_id = tbusc.song_cover_id
	                         and tbusc.user_id = ? and tbusc.del_status = ?  
                             order by tbusc.create_time `
	//根据id查询歌单信息
	QUERY_SONG_COVER_BY_ID = `select song_cover_id,song_cover_name,cover_url
                              from tb_song_cover tsc
                              where song_cover_id = ?`
	//查询用户歌单数量根据歌单名
	QUERY_USER_SONG_COVER_COUNTS_BY_NAME = `select count(*) 
                                            from tb_song_cover as tbsc
                                            inner join tb_song_cover_song tbscs 
                                            on tbsc.song_cover_id = tbscs.song_cover_id
                                            where tbsc.song_cover_name = ? and tbscs.user_id = ?`
	//查询分页歌单列表
	QUERY_PAGE_SONG_COVER_LIST = `select tbsc.song_cover_id,tbsc.song_cover_name ,tbu.user_name,tbu.user_id,
                                  case tbsc.type
                                  when 1 then
                                      '个人'
                                  when 2 then 
                                      '第三方'
                                  end as sctype
                                  from tb_song_cover tbsc
                                  left join  tb_user_song_cover tbusc
                                  on tbsc.song_cover_id = tbusc.song_cover_id 
                                  left join  tb_user tbu
                                  on tbusc.user_id = tbu.user_id 
                                  where 1=1 `
	//查询分页歌单列表个数
	QUERY_PAGE_SONG_COVER_LIST_COUNTS = `select count(*)
                                         from tb_song_cover tbsc
                                         left join  tb_user_song_cover tbusc
                                         on tbsc.song_cover_id = tbusc.song_cover_id 
                                         left join  tb_user tbu
                                         on tbusc.user_id = tbu.user_id 
                                         where 1=1`
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

type QueryUserSongCover struct {
	//用户歌单id
	UserSongCoverId string `gorm:"column:user_song_cover_id"`
	//歌单id
	SongCoverId string `gorm:"column:song_cover_id"`
	//歌单名称
	SongCoverName string `gorm:"column:song_cover_name"`
	//歌单封面url
	SongCoverUrl string `gorm:"column:cover_url"`
	//歌单的类型 1自定义 2 其他
	Type int `gorm:"column:type"`
}

//创建歌单的返回值
type CreateSongCoverReply struct {
	//歌单id
	SongCoverId string
	//歌单名称
	SongCoverName string
}

type BSongCover struct {
	//用户Id
	UserId string `gorm:"column:user_id"`
	//歌单id
	SongCoverId string `gorm:"column:song_cover_id"`
	//歌单名称
	SongCoverName string `gorm:"column:song_cover_name"`
	//用户名称
	UserName string `gorm:"column:user_name"`
	//歌单的类型 1自定义 2 其他
	Type string `gorm:"column:sctype"`
}

type BSongCoverList struct {
	Page utils.Page
	List []BSongCover
}
