package dbModel

import "time"

const (
	QUERY_CHANNEL_DETAIL = `select channel_id,channel_name
                            from tb_channel
                            where channel_id = ?
`
)

//来源渠道
type ChannelTable struct {
	//歌曲来源渠道id
	ChannelId string `gorm:"column:channel_id"`
	//渠道名称
	ChannelName string `gorm:"column:channel_name"`
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

type ChannelDetail struct {
	//歌曲来源渠道id
	ChannelId string `gorm:"column:channel_id"`
	//渠道名称
	ChannelName string `gorm:"column:channel_name"`
}

//渠道信息
type ChannelInfo struct {
	//歌曲来源渠道id
	ChannelId string `gorm:"column:channel_id"`
	//渠道名称
	ChannelName string `gorm:"column:channel_name"`
	//创建人
	CreateUser string `gorm:"column:create_user"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time"`
}

type ChannelInfoList struct {
}
