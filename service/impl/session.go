package impl

import (
	"github.com/astaxie/beego/logs"
	"online-music/common/redis"
	"online-music/common/utils"
	"time"
)

type SessionService struct {
	BaseService
}

/*
*@Title: 设置session
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *SessionService) SetSession(key, value, expire string) error {
	receiver.BeforeLog("SetSession")
	var d time.Duration
	var err error
	if len(expire) != 0 {
		d, err = time.ParseDuration(expire)
		if err != nil {
			logs.Error("设置session-转换过期时间(%s),失败:(%s)", expire, err.Error())
			return utils.NewDetailErr("转换过期时间(%s),失败:(%s)", expire, err.Error())
		}
	}
	client := redis.GetRedis()
	if client == nil {
		logs.Error("设置session-redis得到的链接是nil")
		return utils.NewDetailErr("redis得到的链接是nil")
	}

	cmd := client.Set(key, value, d)
	if cmd == nil {
		logs.Error("设置session-向redis设置key出错，得到cmd为空")
		return utils.NewDetailErr("向redis设置key出错，得到cmd为空")
	}

	if cmd.Err() != nil {
		logs.Error("设置session-向redis设置值是错误:(%v),key:(%v),value:(%v),expire:(%v)", cmd.Err(), key, value, expire)
		return utils.NewDetailErr("向redis设置值是错误:(%v),key:(%v),value:(%v),expire:(%v)", cmd.Err(), key, value, expire)
	}

	return nil
}

/*
*@Title: 得到session
*@Description: 返回的value是一个json的字符串,根据key得到value时，出错，(没有这个key的错)
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *SessionService) GetSession(key string) (string, error) {
	receiver.BeforeLog("GetSession")
	var result string
	client := redis.GetRedis()
	if client == nil {
		logs.Error("得到session-redis得到的链接是nil")
		return result, utils.NewDetailErr("redis得到的链接是nil")
	}

	cmd := client.Get(key)

	//返回一个空串不算错误
	result, _ = cmd.Result()

	return result, nil

}

/*
*@Title: 删除session
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *SessionService) DelSession(key string) error {
	receiver.BeforeLog("DelSession")

	client := redis.GetRedis()
	if client == nil {
		logs.Error("删除session-redis得到的链接是nil")
		return utils.NewDetailErr("redis得到的链接是nil")
	}

	cmd := client.Del(key)
	if cmd.Err() != nil {
		logs.Error("删除session-根据key删除value时错误:(%v)", cmd.Err())
		return utils.NewDetailErr("根据key删除value时错误:(%v)", cmd.Err())
	}
	return nil
}
