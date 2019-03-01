package controller
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../models"
)

func Sign () {
	
}

// 微信登录，用户使用TOKEN登录系统
// 系统使用token去微信换取正确信息
func Login (ctx *gin.Context) {
	var req models.RequestToken
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// request token
}