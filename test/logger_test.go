package test

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

func Test_Logger(t *testing.T) {
	logs.Debug("%s", "123")

}
