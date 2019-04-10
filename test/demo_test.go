package test

import (
	"encoding/json"
	"fmt"
	"online-music/models"
	"testing"
)

func Test_JsonUnmarshal(t *testing.T) {
	var session models.Session
	var str string
	json.Unmarshal([]byte(str), &session)
	fmt.Println(session.UserName)
}
