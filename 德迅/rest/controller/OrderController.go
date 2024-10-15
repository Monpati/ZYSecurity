package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddScdnService(ctx *gin.Context) {
	var reqOrder form.ScdnServiceInfo

	if err := ctx.BindJSON(&reqOrder); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
}

func OrdersLists(ctx *gin.Context) {
	var reqPages *form.Filter
	var info model.ScdnService
	var originPage form.PageOrigin
	var user model.Account

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	username := ctx.GetHeader("Username")
	userId := user.GetIdByAccountName(config.GetDB(), username)
	info.UserId = userId
	if lists, total, err := info.GetOrdersByUser(config.GetDB(), reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "",
			"data": gin.H{
				"lists": lists,
				"total": total,
			},
		})
	}
}
