package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var RegisterByPwd int64 = 1     // 账号密码注册
var RegisterBySmsCode int64 = 2 // 短信验证码快速注册

// 登录方式

var LoginWithPwd int64 = 1
var LoginWithSmsCode int64 = 2

var UserAuthTypeSystem string = "system"       // 平台
var UserAuthTypeWxMiniApp string = "wxMiniApp" // 微信小程序
