package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SwitchPreHeat(ctx *gin.Context) {
	var reqBody *form.ScdnHeatInfo
	var domain model.Domain
	var order model.ScdnService
	var heat model.ScdnHeat
	var info form.DomainHeat

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

	reqBody.DomainUUID = domain_uuid
	reqBody.DomainId = domain_id
	reqBody.ProType = pro_type
	reqBody.OrderUUID = order_uuid

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := heat.UpdateHeatStatus(tx, reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.DDUUID = order_uuid
	info.DomainUUID = domain_uuid
	info.ProType = pro_type

	if err := d.UpdatePreheat(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func UpdatePreHeat(ctx *gin.Context) {
	var tmp *form.ScdnHeatCon
	var reqBody form.ScdnHeatInfo
	var domain model.Domain
	var order model.ScdnService
	var info form.DomainHeatUpdate

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	domain_uuid := domain.GetUuidById(config.GetDB(), domain_id)
	order_id := domain.GetOrderIdById(config.GetDB(), domain_id)
	order_uuid := order.GetUuidById(config.GetDB(), order_id)
	pro_type := order.GetProTypeById(config.GetDB(), order_id)

	if err := ctx.BindJSON(&tmp); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	reqBody.DomainId = domain_id
	reqBody.DomainUUID = domain_uuid
	reqBody.OrderUUID = order_uuid
	reqBody.ProType = pro_type

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, _ := range tmp.CacheConfig {
		reqBody.Url = tmp.CacheConfig[i]

		if err := model.CreateHeat(tx, &reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	}

	info.DDUUID = order_uuid
	info.DomainUUID = domain_uuid
	info.ProType = pro_type
	info.CacheConfig = tmp.CacheConfig

	if err := d.EditPreheat(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func GetPreHeatLists(ctx *gin.Context) {
	var heat model.ScdnHeat
	var reqPages form.HeatFilterForm
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	reqPages.DomainId = domain_id
	if heats, err := heat.GetHeatList(config.GetDB(), &reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"lists": heats,
		})
	}

}
