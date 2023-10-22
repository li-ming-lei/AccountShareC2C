package model

type User struct {
	Code    string `json:"code"`
	UnionId string `json:"unionid"`
	OpenId  string `json:"openid"`
}
