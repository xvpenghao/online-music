package impl

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"online-music/models"
	"online-music/service"
)

type BaseServiceInit struct {
	//基础请求
	BaseRequest models.BaseRequest
}

func (receiver *BaseServiceInit) GetBaseRequest() models.BaseRequest {
	return receiver.BaseRequest
}
func (receiver *BaseServiceInit) SetBaseRequest(request models.BaseRequest) {
	receiver.BaseRequest = request
}

type BaseService struct {
	*BaseServiceInit
}

func (receiver *BaseService) SetInitInfo(init service.IBaseServiceInit) {
	initInfo := new(BaseServiceInit)
	initInfo.BaseRequest = init.GetBaseRequest()
	receiver.BaseServiceInit = initInfo
}

func (receiver *BaseService) GetConn() (*gorm.DB, error) {
	//从配置文件中获取
	username := "root"
	password := "123456"
	dbUrl := "(127.0.0.1:3306)/db_online_music?charset=utf8&parseTime=True&loc=Local"
	res := fmt.Sprintf("%s:%s@tcp%s", username, password, dbUrl)
	fmt.Println(res)
	//user:password@tcp(localhost:5555)/dbname?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", res)
	if err != nil {
		log.Println("数据库链接失败")
		return db, err
	}
	db.LogMode(true)
	db.DB().SetMaxOpenConns(100)
	return db, nil
}

func (receiver *BaseService) BeforeLog(msg string) {
	logs.Debug("%s", msg)
}
