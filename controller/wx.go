package controller
import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"../models"
)
const (
	appid = ""
	secret = ""
)
func Sign () {
	
}

// 微信登录，用户使用TOKEN登录系统
// 系统使用token去微信换取正确信息
func WXLogin (ctx *gin.Context) {
	var req models.RequestToken
	var userLoginModel models.Jscode2session
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var getUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + secret + "&js_code=" + req.Token + "&grant_type=authorization_code"
	resp, err := http.Get(getUrl);
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "远程调用微信服务器失败"})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "读取数据失败"})
		return
	}
	if err := json.Unmarshal(body, &userLoginModel); err == nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "jscode2session转换结构失败"})
		return
	}
	// 持久化到数据库中
	ctx.JSON(http.StatusOK, gin.H{"code": 200})
	// 判断是否在库
	// 写库
	// 创建 jwt
	// resultBody := string(body)
	// request token
}