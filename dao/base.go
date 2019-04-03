package dao

import "github.com/jinzhu/gorm"

type IBaseService interface {
	//得到数据库链接
	GetConn() (*gorm.DB, error)
}
