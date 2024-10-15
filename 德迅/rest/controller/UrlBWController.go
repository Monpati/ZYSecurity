package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUrlBW(ctx *gin.Context) {
	var tmp *form.UrlBWInfo
	var reqBody form.UrlBWCon
	var domain model.Domain
	var order model.ScdnService
	var info form.UrlBW

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

	for i, _ := range tmp.Path {
		reqBody.Path = tmp.Path[i]
	}
	for i, _ := range tmp.Method {
		reqBody.Method = tmp.Method[i]
	}

	reqBody.Type = tmp.Type
	reqBody.OrderId = order_id
	reqBody.DomainId = domain_id
	reqBody.DomainUuid = domain_uuid
	reqBody.OrderUuid = order_uuid
	reqBody.ProType = strconv.FormatInt(pro_type, 10)

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := model.CreateUrlBW(tx, &reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.DDUUID = reqBody.OrderUuid
	info.DomainUUID = reqBody.DomainUuid
	info.ProType = reqBody.ProType
	info.Type = reqBody.Type
	info.URLList.Path = reqBody.Path
	info.URLList.Method = reqBody.Method

	if err := d.UrlBWUpdate(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func UpdateUrlBW(ctx *gin.Context) {
	var reqBody *form.UrlBWCon
	var cm model.UrlBW

	id, _ := strconv.ParseInt(ctx.Param("urlbw"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}

	if err := cm.UpdateUrlBW(config.GetDB(), id, reqBody); err != nil {
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

func UrlBWLists(ctx *gin.Context) {
	var reqPages form.Filter
	var list model.UrlBW
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetUrlBWLists(config.GetDB(), &reqPages); err != nil {
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
