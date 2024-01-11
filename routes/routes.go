package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sami-alex/netflix-go-backend/api/users"
)

func CreateRoutes(ctx *gin.Engine) {
	users.UserRoutes(ctx.Group("/user"))

}
