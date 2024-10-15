package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddPreAcc(ctx *gin.Context) {
	var reqBody *form.PreAccInfo
	var domain model.Domain
	var order model.ScdnService
	var info form.PreAcc
	var reqInfo form.PreAccConInfo

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	domain_uuid := domain.GetUuidById(config.GetDB(), domain_id)
	order_id := domain.GetOrderIdById(config.GetDB(), domain_id)
	order_uuid := order.GetUuidById(config.GetDB(), order_id)
	pro_type := order.GetProTypeById(config.GetDB(), order_id)

	reqBody.DomainUUID = domain_uuid
	reqBody.DDUUID = order_uuid
	reqBody.ProType = pro_type

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, _ := range reqBody.Rule {
		reqInfo.DomainUuid = domain_uuid
		reqInfo.DomainId = domain_id
		reqInfo.OrderUuid = order_uuid
		reqInfo.OrderId = order_id
		reqInfo.Action = reqBody.Action
		reqInfo.Active = reqBody.Active
		reqInfo.CheckList = reqBody.Rule[i].MItem
		reqInfo.MItem = reqBody.Rule[i].MItem
		reqInfo.MValue = reqBody.Rule[i].MValue
		reqInfo.MOperate = reqBody.Rule[i].MOperate
		reqInfo.MValueXs = reqBody.Rule[i].MValueXs
	}

	if err := model.CreatePreAcc(tx, &reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.Action = reqBody.Action
	info.Active = reqBody.Active
	info.BlockTime = reqBody.BlockTime
	info.DDUUID = reqBody.DDUUID
	info.DomainUUID = reqBody.DomainUUID
	info.Location = reqBody.Location
	info.ProType = reqBody.ProType
	info.Rule.CheckList = reqBody.CheckList
	info.Rule.Rule = reqBody.Rule

	if err := d.UpdatePreAccCon(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func UpdatePreAcc(ctx *gin.Context) {
	var reqBody *form.PreAccInfo
	var cm model.PreAcc

	id, _ := strconv.ParseInt(ctx.Param("cc"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}

	if err := cm.UpdatePreAcc(config.GetDB(), id, reqBody); err != nil {
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

func PreAccLists(ctx *gin.Context) {
	var reqPages form.Filter
	var list model.PreAcc
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetPreAccLists(config.GetDB(), &reqPages); err != nil {
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
