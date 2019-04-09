package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GetConn() (*gorm.DB, error) {
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
