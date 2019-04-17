package redis

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"online-music/common/constants"
)

//利用map是线程安全的机制
var redisMap map[string]*redis.Client

func init() {
	logs.Debug("***************************************************")
	logs.Debug("***************redis初始化**************************")
	logs.Debug("***************************************************")
	//初始化redis
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.217.80:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//初始化map
	redisMap = make(map[string]*redis.Client)
	ping, err := client.Ping().Result()
	if err != nil {
		logs.Error("redis链接失败:(%v)", err.Error())
		panic("redis链接失败")
	} else {
		logs.Debug("redis链接成功：", ping)
	}

	redisMap[constants.REDIS_CLIENT] = client
}

func GetRedis() *redis.Client {

	if v, ok := redisMap[constants.REDIS_CLIENT]; ok {
		return v
	}
	return nil
}
