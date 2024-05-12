package routes

import (
	"hellowApi/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes はすべてのルートを登録します
func RegisterRoutes(router *gin.Engine) {
    router.GET("/users", controllers.GetUsers)
}
