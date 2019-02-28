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
)

func main() {
    var router = gin.Default()
    var a models.Article = new (models.Article)
    router.Run("8088")
}