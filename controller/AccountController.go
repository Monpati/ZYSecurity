package controller

import (
	"Dexun/config"
	"Dexun/config/rest"
	"Dexun/form"
	"Dexun/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(ctx *gin.Context) {
	var reqAccount form.AccountInfo
	if err := ctx.BindJSON(&reqAccount); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := model.AccountCreate(rest.EngineCfg.MysqlDB, &reqAccount); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  200,
			"error": "",
			"data":  gin.H{},
		})
	}
}

func Login(ctx *gin.Context) {
	var reqAccount model.Account
	if err := ctx.BindJSON(&reqAccount); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	fmt.Println(reqAccount.FindSaltByAccountName(config.GetDB(), reqAccount.Username))
	fmt.Println(reqAccount.Password)
	//hashedPassword, _ := model.HashSalt(reqAccount.Password, reqAccount.FindSaltByAccountName(config.GetDB(), reqAccount.Username))
	if err := reqAccount.FindByAccountName(rest.EngineCfg.MysqlDB, reqAccount.Username, reqAccount.Password); err == nil {
		token, _ := config.ReleaseToken(reqAccount)
		if err := config.StoreTokenInRedis(config.GetRedis(), token, reqAccount.Username); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "系统异常",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "登录成功",
				"data": gin.H{
					"username": reqAccount.Username,
					"password": reqAccount.Password,
					"token":    token,
				}})
		}
	} else {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": err,
		})
		return
	}
}

func Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	username := ctx.GetHeader("username")
	if token != "" && username != "" {
		storedToken, err := config.GetTokenFromRedis(config.GetRedis(), username)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
		if token == storedToken {
			err := config.DeleteTokenFromRedis(config.GetRedis(), username)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    http.StatusBadRequest,
					"message": err,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 0,
					"data": nil,
				})
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
	}
}

func Info(ctx *gin.Context) {
	var usr model.Account

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if accountInfo, err := usr.GetInfo(rest.EngineCfg.MysqlDB, id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"id":       1,
				"username": accountInfo.Username,
				//"realName": 'admin',
				"phoneNumber": accountInfo.TelNum,
				"email":       accountInfo.Email,
				//"avatarUrl": '',
				//"roleList": ['超级管理员']
			},
		})
	}
}
