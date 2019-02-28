package main
import (
    "fmt"
    "html"
    "log"
    "github.com/gin-gonic/gin"
    "net/http"
)
import (
    "./models",
    "./controller"
)

func main() {
    var router = gin.Default()
    router.Run("8088")
}