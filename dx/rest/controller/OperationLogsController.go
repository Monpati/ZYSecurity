package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"Dexun/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OperationLogAdd(ctx *gin.Context) error {
	username := ctx.GetHeader("Username")
	originUrl := ctx.GetHeader("Origin")
	reqUrl := ctx.GetHeader("Host")
	userAgent := ctx.GetHeader("User-Agent")
	var reqLog form.OperationLog
	var reqAccount model.Account
	var request utils.JSON
	if err := ctx.BindJSON(request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	reqLog.UserId = reqAccount.GetIdByAccountName(config.GetDB(), username)
	reqLog.OriginUrl = originUrl
	reqLog.UserAgent = userAgent
	reqLog.Request = request
	reqLog.ReqUrl = reqUrl

	if err := model.OperationLogCreate(config.GetDB(), &reqLog); err != nil {
		return err
	} else {
		return nil
	}
}
