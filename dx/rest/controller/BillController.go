package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"Dexun/model/Dexun"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AllBillLists(ctx *gin.Context) {
	var info model.Bill
	var reqPages form.BillsFilterForm
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if lists, total, err := info.GetByParams(config.GetDB(), &reqPages); err != nil {
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

func BillLists(ctx *gin.Context) {
	var reqPages *form.BillsFilterForm
	var info model.Bill
	var originPage form.PageOrigin
	var user model.Account

	username := ctx.GetHeader("Username")
	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	userId := user.GetIdByAccountName(config.GetDB(), username)
	info.UserId = userId
	if lists, total, err := info.GetBills(config.GetDB(), reqPages); err != nil {
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

func AddScdnOrder(ctx *gin.Context) {
	var billInfo form.Purchase
	var account model.Account
	var results []Dexun.Data
	var scdn model.ScdnService
	var combos Dexun.ScdnCombos

	billInfo.Username = ctx.GetHeader("Username")
	if err := ctx.BindJSON(&billInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	id := account.GetIdByAccountName(config.GetDB(), billInfo.Username)
	billInfo.UserId = id
	months, _ := strconv.ParseInt(billInfo.Months, 10, 64)
	tmp, _ := strconv.ParseInt(billInfo.ComboId, 10, 64)
	combo_id := combos.GetIdByUuid(config.GetDB(), tmp)
	source := combos.GetSourceById(config.GetDB(), combo_id)
	balance := account.GetBalanceById(config.GetDB(), billInfo.UserId)

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if balance < combos.ZyksMoney {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "余额不足！",
		})
		tx.Rollback()
		return
	}

	orderid := d.CreateSCDNOrder(billInfo.ComboUuid, months)
	orderUuids := strings.Fields(orderid)
	results = d.GetBillCallBack(orderUuids)
	for i, _ := range results {
		if results[i].OrderStatus == 1 {
			//这里回调只会回调这一个订单，只需确认状态是否成功，然后根据返回的订单uuid，再查询该订单的相关信息，将信息写入ScdnService
			d.GetSCDNList("", "100", "desc", 1)
			for _, item := range d.Data {
				if item.UUID == orderid {
					if err := model.CreateScdnService(tx, item, combo_id, months, billInfo.UserId, results[i].ProductType, billInfo.Username, source); err != nil {
						ctx.JSON(http.StatusOK, gin.H{
							"code":    http.StatusBadRequest,
							"message": err,
						})
						tx.Rollback()
					} else {
						ctx.JSON(http.StatusOK, gin.H{
							"node":    0,
							"message": nil,
						})
					}
					service_id := scdn.GetIdByUuid(tx, orderid)
					billInfo.ServiceId = &service_id
					if err := model.CreateBill(tx, billInfo); err != nil {
						ctx.JSON(http.StatusOK, gin.H{
							"code":    http.StatusBadRequest,
							"message": err,
						})
						tx.Rollback()
					}
					if err := account.UpdateBalancePurchase(tx, id, balance, combos.ZyksMoney); err != nil {
						ctx.JSON(http.StatusOK, gin.H{
							"code":    http.StatusBadRequest,
							"message": err,
						})
						tx.Rollback()
					}
				}
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "failed",
			})
			tx.Rollback()
		}
	}

	tx.Commit()
}
