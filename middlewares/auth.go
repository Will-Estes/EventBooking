package middlewares

import (
	"eventbooking/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	userId, ok := utils.IsTokenValid(token)
	if !ok {
		fmt.Println("Error parsing auth token")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
