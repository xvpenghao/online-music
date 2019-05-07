package constants

//用户删除的状态
const (
	// 1 删除用户
	USER_DEL_STATUS = 1
	// 2 不删除用户
	USER_NO_DEL_STATUS = 2
)

//session的key
const (
	SESSION_USER = "session_user"
)

//redis
const (
	//redis客户端
	REDIS_CLIENT = "redis_client"
)

//cookie
const (
	//设置cookie的过期使，或者rediskey的过期时间
	COOKIE_EXPIRE = "2h"
	//设置cookie名字
	COOKIE_NAME = "xph"
	//设置cookie秘钥
	COOKIE_SECRET = "xvpenghao"
)

//音乐渠道
const (
	//网易云音乐
	WYY_MUSCI_TYPE = "1"
	//QQ音乐
	QQ_MUSIC_TYPE = "2"
)

//设置userAgent
const (
	//userAgent header
	USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36"
)

//歌单或歌曲
const (
	//歌单封面列表url
	SONG_COVER_LIST_URL = "https://music.163.com/discover/playlist"

	//分页获取数据
	SONG_PAGE_LIST_URL = "https://music.163.com/discover/playlist/?order=hot&cat=全部&limit=%s&offset=%s"

	//歌单详情
	SONG_COVER_DETAIL_URL = "https://music.163.com/playlist?id=%s"

	//歌曲url
	SONG_URL = "https://music.163.com/song?id=%s"

	//歌曲播放url
	SONG_PLAY_URL = "http://music.163.com/song/media/outer/url?id=%s.mp3"

	//歌词链接
	SONG_LYRIC_URL = "http://music.163.com/api/song/lyric?id=%s&lv=1&kv=1&tv=-1"

	//默认分页大小，35一页
	DEFAULT_PAGE_SIZE = "35"
)

//歌单类型
const (
	//歌单类型自定义
	SONG_COVER_TYPE_CUSTOMER = 1
	//歌单类型其他（其他的平台的歌单）
	SONG_COVER_TYPE_OTHER = 2
)

//删除的状态
const (
	//删除1
	DEL_STATUS = 1
	//不删除
	NOT_DEL_STATUS = 2
)

//创建用户歌单的默认值
const (
	//歌单默认详情
	SONG_COVER_DEFAULT_DESCRIPTION = "无"
	//歌单默认封面图片
	SONG_COVER_DEFAULT_COVER_URL = "/static/me/temp/default_cover.png"
)

//爬取歌曲的数量
const (
	//爬虫歌曲的数量
	SPIDER_SONG_COUNT = 3
)

//播放历史
const (
	//%s代表的是当前登录用户id
	CREATE_SONG_PLAY_HISTORY = "create_song_play_history:ph:%s"
	//播放历史存储的最大数量
	SONG_PLAY_HISTORY_MAX_COUNT = 20
)

//歌曲搜索
const (
	//请求秘钥 579621905(默认值) =>key
	SEARCH_REQUEST_SECRET = "579621905"
	//搜索类型 =>type  song,video,.....
	SEARCH_TYPE_STATUS = "song"
	//请求的数量默认 100  =>limit
	SEARCH_LIMIT = 5
	//分页 默认第一页  => offset
	SEARCH_OFFSET = 0
)

//来源渠道
const (
	//网易平台
	CHANNEL_WANGYI_PLATFORM = "1"
	//QQ平台
	CHANNEL_QQ_PLATFORM = "2"
)

//音乐平台url
const (
	//网易云音乐搜索url
	SEARCH_WANGYI_PLATFORM_URL = "https://api.itooi.cn/music/netease/search?key=%s&s=%s&type=%s&limit=%d&offset=%d"
	//QQ音乐搜索的url
	SEARCH_QQ_PLATFORM_URL = "https://api.itooi.cn/music/tencent/search?key=%s&s=%s&type=%s&limit=%d&offset=%d"
)

//分页
const (
	//默认页
	PAGE_DEFAULT_CUR_PAGE = 1
	//每一页的大小
	PAGE_DEFAULT_LIMIT = 3
	//分页的分组
	PAGE_DEFAULT_GROUPS = 5
)

//后台session的key
const (
	BACK_SESSION_KEY = "admin"
)
