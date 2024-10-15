package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCC(ctx *gin.Context) {
	var reqBody *form.CC
	var domain model.Domain
	var order model.ScdnService

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	domain_uuid := domain.GetUuidById(config.GetDB(), domain_id)
	order_id := domain.GetOrderIdById(config.GetDB(), domain_id)
	order_uuid := order.GetUuidById(config.GetDB(), order_id)
	pro_type := order.GetProTypeById(config.GetDB(), order_id)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	reqBody.UseDefault = "1"
	reqBody.OrderId = order_id
	reqBody.DomainId = domain_id
	reqBody.DomainUUID = domain_uuid
	reqBody.DDUUID = order_uuid
	reqBody.ProType = strconv.FormatInt(pro_type, 10)

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := model.CreateCC(tx, reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	if err := d.UpdateCCInfo(reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Rollback()
}

func UpdateCC(ctx *gin.Context) {
	var reqBody *form.CCInfo
	var cm model.CC

	id, _ := strconv.ParseInt(ctx.Param("cc"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}

	if err := cm.UpdateCC(config.GetDB(), id, reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": nil,
		})
	}
}

func CCLists(ctx *gin.Context) {
	var reqPages form.CcFilterForm
	var list model.CC
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetCCLists(config.GetDB(), &reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "",
			"data": gin.H{
				"lists": lists,
				"total": total,
			},
		})
	}
}
