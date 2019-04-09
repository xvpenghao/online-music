package test

import (
	"fmt"
	"online-music/common/db"
	"testing"
)

func Test_DBConnect(t *testing.T) {
	db, err := db.GetConn()
	if err != nil {
		fmt.Printf("错误:%s\n", err.Error())
	} else {
		fmt.Println("链接成功", db)
	}
}
