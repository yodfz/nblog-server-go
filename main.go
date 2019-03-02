package main
import (
    "fmt"
    "html"
    "log"
    "github.com/gin-gonic/gin"
    "net/http"
)
import (
    "./models"
    "./controller"
)

func main() {
    var router = gin.Default()
    router.POST("/wx/login",controller.WXLogin)
    router.Run("8088")
}