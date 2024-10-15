package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"Dexun/model/Dexun"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetCloudEyeComboLists() {
	var results []Dexun.Datum
	var combo *Dexun.DexunCloudEyeList
	results = d.GetCloudEyeComboList()

	for i, _ := range results {
		if combo = combo.FindComboByUuid(config.GetDB(), results[i].UUID); combo.UUID != "" {
			if err := combo.UpdateCloudEyeByUUID(config.GetDB(), combo.UUID, &results[i]); err != nil {
				continue
			}
		} else if combo.UUID == "" {
			if err := Dexun.CreateDxCloudEye(config.GetDB(), &results[i]); err != nil {
				continue
			}
		}
	}
}

func CopyDexunCloudEyeList() {
	if err := Dexun.CopyDxCloudEyeList(config.GetDB()); err != nil {
		fmt.Println(err)
	}
}

func CloudEyeLists(ctx *gin.Context) {
	var info Dexun.CloudEyeList
	var reqPages form.Filter
	var originPage form.PageOrigin

	role := ctx.GetHeader("Role")
	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if role == "user" {
		if lists, total, err := info.UserGetCloudEyeByParams(config.GetDB(), &reqPages); err != nil {
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
	} else if role == "admin" {
		if lists, total, err := info.AdminGetCloudEyeByParams(config.GetDB(), &reqPages); err != nil {
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
}

func AddCloudEyeOrders(ctx *gin.Context) {
	var billInfo form.Purchase
	var account model.Account
	var results []Dexun.Data
	var ddos model.DDoSService
	var combos Dexun.DDoSCombos
	var reqInfo form.DDoSServiceInfo
	var info form.CloudEyeOrderInfo

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
	info.TcUUID = billInfo.ComboUuid
	info.Months = months

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	orderid := d.CreateCloudEyeOrder(&info)
	orderUuids := strings.Fields(orderid)
	results = d.GetBillCallBack(orderUuids)
	for i, _ := range results {
		if results[i].OrderStatus == 1 {
			//这里回调只会回调这一个订单，只需确认状态是否成功，然后根据返回的订单uuid，再查询该订单的相关信息，将信息写入DDoSService
			d.GetDDoSList("", "100", "desc", 1)
			for _, item := range d.Data {
				if item.UUID == orderid {
					reqInfo.ComboId = combo_id
					reqInfo.Uuid = item.UUID
					reqInfo.UUserId = item.UUserID
					reqInfo.ServerIp = item.ServerIP
					reqInfo.TcName = item.TcName
					reqInfo.KsMoney = item.KsMoney
					reqInfo.ProductSitename = item.ProductSitename
					reqInfo.StatTime = item.StatTime
					reqInfo.EndTime = item.EndTime
					reqInfo.SiteStart = item.SiteStart
					reqInfo.DdosHh = item.DdosHh
					reqInfo.KsStart = item.KsStart
					reqInfo.DomainNum = item.DomainNum
					reqInfo.RechargeDomain = *item.RechargeDomain
					reqInfo.PortNum = *item.PortNum
					reqInfo.RechargePort = *item.RechargePort
					reqInfo.UserId = id
					reqInfo.Agent = source

					if err := model.CreateDDoSService(tx, &reqInfo); err != nil {
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
					*billInfo.DDoSId = ddos.GetIdByUuid(tx, orderid)
					if err := model.CreateBill(tx, billInfo); err != nil {
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
