package main

import (
	"hellowApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    routes.RegisterRoutes(router)

    router.Run(":8080") // 8080ポートでサーバを起動
}