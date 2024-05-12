package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers は全ユーザの情報をJSON形式で返します
func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "users": []string{"Alice", "Bob", "Charlie"},
    })
}
