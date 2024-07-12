package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-rest-api/utils"
)

func Authenticated(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you are not authenticated"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you are not authenticated"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
