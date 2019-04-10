package utils

import (
	"fmt"
	"strings"
)

//自定错误
type CustomerErr struct {
	Code string
	Msg  string
}

func (e CustomerErr) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Msg)
}

// NewErr : 返回新错误
func NewErr(code, msg string, err ...error) error {
	var result CustomerErr
	result.Code = code

	errStr := ""
	for _, e := range err {
		errStr = fmt.Sprintf("%s : %s ", errStr, e.Error())
	}
	//可能会出现 "空串":"错误信息"
	errStr = strings.TrimPrefix(errStr, " :")

	if len(errStr) > 0 {
		result.Msg = fmt.Sprintf("%s : %s", msg, errStr)
	} else {
		result.Msg = fmt.Sprintf("%s", msg)
	}

	return result
}

func NewDetailErr(format string, a ...interface{}) error {
	msg := fmt.Sprintf(format, a...)
	return NewSysErr(msg)
}

func NewSysErr(msg string, err ...error) error {
	return NewErr(E_MUSIC_999_CODE, msg, err...)
}

func NewDBErr(msg string, err ...error) error {
	return NewErr(E_MUSIC_998_CODE, msg, err...)
}
