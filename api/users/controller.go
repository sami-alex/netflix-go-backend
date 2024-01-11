package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sami-alex/netflix-go-backend/utils"
)

func CreateUser(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error parsing user data",
		})
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error hashing password",
		})
		return
	}
	user.Password = hashedPassword
	userId, err := user.CreateUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating user data",
			"err":     err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "sucessfuly created user",
		"_id":     userId,
	})

}

func Login(ctx *gin.Context) {
	var user *User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = user.Login()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var userData = make(map[string]interface{})
	userData["email"] = user.Email
	userData["userName"] = user.UserName
	userData["_id"] = user.ID
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "sucessfully logged in",
		"user":    userData,
		"token":   token,
	})
}
func Check(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "active",
		"userId":  userId,
	})
}
