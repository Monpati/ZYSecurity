package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddSafeAcc(ctx *gin.Context) {
	var tmp *form.SafeAccInfo
	var reqBody form.SafeAccCon
	var domain model.Domain
	var order model.ScdnService
	var info form.SafeCon

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	if err := ctx.BindJSON(&tmp); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	domain_uuid := domain.GetUuidById(config.GetDB(), domain_id)
	order_id := domain.GetOrderIdById(config.GetDB(), domain_id)
	order_uuid := order.GetUuidById(config.GetDB(), order_id)
	pro_type := order.GetProTypeById(config.GetDB(), order_id)

	reqBody.OrderId = order_id
	reqBody.OrderUuid = order_uuid
	reqBody.DomainId = domain_id
	reqBody.DomainUuid = domain_uuid
	reqBody.ProType = strconv.FormatInt(pro_type, 10)

	for i, _ := range tmp.URL {
		reqBody.URL = tmp.URL[i]
	}
	for i, _ := range tmp.Password {
		reqBody.Password = tmp.Password[i]
	}

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := model.CreateSafeAcc(tx, &reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.DDUUID = reqBody.OrderUuid
	info.DomainUUID = reqBody.DomainUuid
	info.ProType = reqBody.ProType
	info.Config.URL = reqBody.URL
	info.Config.Password = reqBody.Password

	if err := d.UpdateSafeCon(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func UpdateSafeAcc(ctx *gin.Context) {
	var reqBody *form.SafeAccInfo
	var cm model.SafeAcc

	id, _ := strconv.ParseInt(ctx.Param("cc"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}

	if err := cm.UpdateSafeAcc(config.GetDB(), id, reqBody); err != nil {
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

func SafeAccLists(ctx *gin.Context) {
	var reqPages form.Filter
	var list model.SafeAcc
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetSafeAccLists(config.GetDB(), &reqPages); err != nil {
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
