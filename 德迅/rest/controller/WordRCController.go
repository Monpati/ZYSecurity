package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddWordRC(ctx *gin.Context) {
	var reqBody *form.WordRCInfo
	var domain model.Domain
	var order model.ScdnService
	var info form.WordsRc

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

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	reqBody.OrderId = order_id
	reqBody.OrderUuid = order_uuid
	reqBody.DomainId = domain_id
	reqBody.DomainUuid = domain_uuid

	if err := model.CreateWordRC(tx, reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.Active = reqBody.Active
	info.DDUUID = reqBody.OrderUuid
	info.DomainUUID = reqBody.DomainUuid
	info.Gzip = reqBody.Gzip
	info.Keywords = reqBody.KeyWords
	info.ProType = strconv.FormatInt(pro_type, 10)

	if err := d.UpdateWordsRC(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func UpdateWordRC(ctx *gin.Context) {
	var reqBody *form.WordRCInfo
	var cm model.WordRC

	id, _ := strconv.ParseInt(ctx.Param("cc"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}

	if err := cm.UpdateWordRC(config.GetDB(), id, reqBody); err != nil {
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

func WordRCLists(ctx *gin.Context) {
	var reqPages form.Filter
	var list model.WordRC
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetWordRCLists(config.GetDB(), &reqPages); err != nil {
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
