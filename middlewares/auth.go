package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sami-alex/netflix-go-backend/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}
	userId, err := utils.VerifyToken(strings.TrimPrefix(token, "Bearer "))
	if err != nil {
		fmt.Print(err, "2")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
