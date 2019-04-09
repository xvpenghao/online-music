package utils

// RPC 错误 及 提示信息
var RPCErrMap = map[string]string{
	E_MUSIC_999_CODE: "系统错误",
	E_MUSIC_998_CODE: "数据库错误",
}

func GetMsg(code string) string {
	if msg, ok := RPCErrMap[code]; ok {
		return msg
	}

	return RPCErrMap[E_MUSIC_999_CODE]
}
