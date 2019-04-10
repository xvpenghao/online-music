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
	COOKIE_EXPIRE = "24h"
	//设置cookie名字
	COOKIE_NAME = "xph"
	//设置cookie秘钥
	COOKIE_SECRET = "xvpenghao"
)
