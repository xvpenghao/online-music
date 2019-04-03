package impl

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BaseService struct {
}

func (receiver *BaseService) GetConn() (*gorm.DB, error) {
	//从配置文件中获取
	username := ""
	password := ""
	dbUrl := ""
	res := username + ":" + password + dbUrl
	db, err := gorm.Open("mysql", res)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	return db, nil
}
