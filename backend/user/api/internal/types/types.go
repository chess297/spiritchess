// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type GetUserInfoResp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	ExpireAt int    `json:"expireAt"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	ExpireAt int    `json:"expireAt"`
}
