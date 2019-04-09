package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetUUID() string {
	uid, _ := uuid.NewV4()
	return strings.Replace(uid.String(), "-", "", -1)
}

//获取字符串长度（中文一个字算一个）
func GetStringLen(str string) int {
	return len([]rune(str))
}
