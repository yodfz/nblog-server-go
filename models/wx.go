package models

type WX struct {

}

type WXPay struct {

}

// code2Session 接口返回值
type Jscode2session struct {
	Openid string	`json:"openid"`
	Session_key string	`json:"session_key"`
	Unionid string	`json:"unionid"`
	Errcode	int	`json:"errcode"`
	Errmsg	string	`json:"errmsg"`
}