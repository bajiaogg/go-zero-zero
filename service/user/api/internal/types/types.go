// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Uid      string `json:"uid"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type WxMiniAuthReq struct {
	Code string `json:"code"`
}

type WxMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoReq struct {
	Uid string `path:"uid"`
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}