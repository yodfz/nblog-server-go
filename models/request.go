package models

/**
* 用于解构URL
**/
type RequestDetail struct {
	id string `uri:"id" binding:"required"`
}

// 微信换取用户信息TOKEN
type RequestToken struct {
	token string `form:"token" json:"token" xml:"token"  binding:"required"`
}