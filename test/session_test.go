package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"online-music/service/impl"
	"testing"
)

func Test_Session_SetSession(t *testing.T) {
	sessionService := new(impl.SessionService)
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "张三", Age: 23}
	value, _ := json.Marshal(&user)
	key := "user"
	sessionService.SetSession(key, string(value), "30m")
}

func Test_Session_GetSession(t *testing.T) {
	sessionService := new(impl.SessionService)
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	key := "user"
	var user User
	resutl, err := sessionService.GetSession(key)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(bytes.NewBufferString(resutl).Bytes(), &user)
		fmt.Printf("%+v", user)
	}
}

func Test_Session_DelSession(t *testing.T) {
	sessionService := new(impl.SessionService)
	key := "user"
	err := sessionService.DelSession(key)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("删除成功")
	}
}
