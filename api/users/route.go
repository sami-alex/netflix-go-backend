package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sami-alex/netflix-go-backend/middlewares"
)

func UserRoutes(ctx *gin.RouterGroup) {
	ctx.POST("/signup", CreateUser)
	ctx.POST("/login", Login)
	ctx.Use(middlewares.Authenticate)
	ctx.GET("/", Check)
}
