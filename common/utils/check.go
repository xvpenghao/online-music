package utils

import (
	"math"
	"regexp"
)

// STR_MAX_LEN :字符串最长长度
const STR_MAX_LEN = 20000

// INT_MAX :INT 最大值
const INT_MAX = math.MaxInt64

//参数检测用
type Checker interface {
	isLegal() bool
}

type StrChecker struct {
	Value  string
	MaxLen int
	MinLen int
}

func (c *StrChecker) isLegal() bool {
	return c.MinLen <= GetStringLen(c.Value) && GetStringLen(c.Value) <= c.MaxLen
}

// IntChecker : 数字检查
type IntChecker struct {
	Value int
	Max   int64
	Min   int
}

func (c *IntChecker) isLegal() bool {
	return c.Min <= c.Value && int64(c.Value) <= c.Max
}

//检测所有的checker是否正确
func CheckLegal(checkers ...Checker) bool {
	for _, checker := range checkers {
		if !checker.isLegal() {
			return false
		}
	}
	return true
}

// EmailCheck : 邮箱正则匹配检测
type EmailCheck struct {
	Value string
}

func (c *EmailCheck) isLegal() bool {

	b, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", c.Value)
	if false == b {
		return false
	}
	return true
}

type PwdCheck struct {
	Value string
}

func (p *PwdCheck) isLegal() bool {
	b1, _ := regexp.MatchString("[a-z0-9A-Z]{6,12}", p.Value)
	b2, _ := regexp.MatchString("[a-z]", p.Value)
	b3, _ := regexp.MatchString("[A-Z]", p.Value)
	b4, _ := regexp.MatchString("[0-9]", p.Value)
	if b1 && b2 && b3 && b4 {
		//匹配成功
		return true
	}
	return false
}

//StrEqualCher:字符串是否相等
type StrEqualsChecker struct {
	Value       string
	SecondValue string
}

func (c *StrEqualsChecker) isLegal() bool {
	return c.Value == c.SecondValue
}
