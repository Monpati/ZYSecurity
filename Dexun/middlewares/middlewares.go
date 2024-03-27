package middlewares

import (
	"Dexun/config"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRequired(ctx *gin.Context) {
	existAdmin(ctx, ctx.GetHeader("Role"))
}

func UserRequired(ctx *gin.Context) {
	existUser(ctx, ctx.GetHeader("Role"))
}

func existAdmin(ctx *gin.Context, role string) {
	uToken := ctx.GetHeader("Authorization")
	username := ctx.GetHeader("Username")
	var reqAccount model.Account
	token, err := config.GetTokenFromRedis(config.GetRedis(), username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if uToken == token {
		if err := reqAccount.GetRole(config.GetDB(), username); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
		if reqAccount.Role == role && role == "admin" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": err,
			})
		}
	}
}

func existUser(ctx *gin.Context, role string) {
	uToken := ctx.GetHeader("Authorization")
	username := ctx.GetHeader("Username")
	var reqAccount model.Account
	token, err := config.GetTokenFromRedis(config.GetRedis(), username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if uToken == token {
		if err := reqAccount.GetRole(config.GetDB(), username); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
		if reqAccount.Role == role && role == "user" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": err,
			})
		}
	}
}
