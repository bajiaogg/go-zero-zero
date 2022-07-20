package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[DefaultCode] = "SUCCESS"
	message[DefaultErrorCode] = "服务器开小差啦,稍后再来试一试"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[MissParameter] = "缺少参数"
	message[InvalidParameter] = "非法参数"
	message[LimitedError] = "请求太快，休息一下吧"
	message[TokenGenerateError] = "生成token失败"
	message[TokenExpireError] = "token已失效，请重新登录"
	message[OperationError] = "操作有误"
}

func MapErrMsg(codeNo uint32) string {
	if msg, ok := message[codeNo]; ok {
		return msg
	} else {
		return ""
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
