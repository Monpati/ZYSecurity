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

func DDoSLists(ctx *gin.Context) {
	var info Dexun.DDoSCombos
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
		if lists, total, err := info.UserGetDDoSByParams(config.GetDB(), &reqPages); err != nil {
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
		if lists, total, err := info.AdminGetDDoSByParams(config.GetDB(), &reqPages); err != nil {
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

func UpdateDDoSCombosStatus(ctx *gin.Context) {
	var ddos Dexun.DDoSCombos
	var reqBody struct {
		Status int `json:"status"`
	}

	id, _ := strconv.ParseInt(ctx.Param("ddos"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := ddos.ChangeDDoSCombosStatus(config.GetDB(), id, reqBody.Status); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}
}

func GetDDoSLists(ctx *gin.Context) {
	var services model.DDoSService
	var reqPages form.Filter
	var originPage form.PageOrigin
	var user model.Account

	username := ctx.GetHeader("Username")
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
		services.UserId = user.GetIdByAccountName(config.GetDB(), username)
		if lists, total, err := services.GetOrdersByUser(config.GetDB(), &reqPages); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": gin.H{
					"lists": lists,
					"total": total,
				},
			})
		}
	} else if role == "agent" {
		services.Agent = username
		if lists, total, err := services.GetOrdersByAgent(config.GetDB(), &reqPages); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": gin.H{
					"lists": lists,
					"total": total,
				},
			})
		}
	} else if role == "admin" {
		if lists, total, err := services.GetByParams(config.GetDB(), &reqPages); err != nil {
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

func AddDDoSOrders(ctx *gin.Context) {
	var billInfo form.Purchase
	var account model.Account
	var results []Dexun.Data
	var ddos model.DDoSService
	var combos Dexun.DDoSCombos
	var reqInfo form.DDoSServiceInfo
	var info form.OrderConsume

	billInfo.Username = ctx.GetHeader("Username")
	if err := ctx.BindJSON(&billInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	id := account.GetIdByAccountName(config.GetDB(), billInfo.Username)
	billInfo.UserId = id
	balance := account.GetBalanceById(config.GetDB(), billInfo.UserId)
	months, _ := strconv.ParseInt(billInfo.Months, 10, 64)
	tmp, _ := strconv.ParseInt(billInfo.ComboId, 10, 64)
	combo_id := combos.GetIdByUuid(config.GetDB(), tmp)
	source := combos.GetSourceById(config.GetDB(), combo_id)
	info.UUID = billInfo.ComboUuid
	info.Months = months

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if balance < combos.ZyksMoney {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "余额不足!",
		})
		tx.Rollback()
	}

	orderid := d.CreateDDoSOrder(&info)
	//orderid := "cd1cc971-a414-46ce-8d7a-0d5aba45a7fa"
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
					reqInfo.ProType = results[i].ProductType

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

					ddos_id := ddos.GetIdByUuid(tx, orderid)
					billInfo.DDoSId = &ddos_id
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

func GetDDoSComboLists() {
	var results []Dexun.Data
	var combo *Dexun.DexunDDoSCombos
	results = d.GetDDoSComboList()

	for i, _ := range results {
		if combo = combo.FindComboByUuid(config.GetDB(), results[i].UUID); combo.Uuid != "" {
			if err := combo.UpdateComboByUuid(config.GetDB(), combo.Uuid, results[i]); err != nil {
				continue
			}
		} else if combo.Uuid == "" {
			if err := Dexun.CreateDxDDoSCombos(config.GetDB(), results[i]); err != nil {
				continue
			}
		}
	}
}

func AllowDexunDDoSSell() {
	if err := Dexun.CopyDxDDoSCombos(config.GetDB()); err != nil {
		fmt.Println(err)
	}
}

func DDoSLoginPanel(ctx *gin.Context) {
	var reqInfo *form.LoginPanel

	if err := ctx.BindJSON(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := d.LoginDDoSPanel(reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Success",
		})
	}
}

func DDoSAddDomain(ctx *gin.Context) {
	var domains form.DomainInfo
	var domain model.Domain
	var reqInfo form.DomainsInfo
	var user model.Account
	var order model.DDoSService
	var info model.ConfigList
	var clInfo form.ConfigListInfo
	var saInfo form.SourceAddressInfo

	username := ctx.GetHeader("Username")
	if err := ctx.BindJSON(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	id := user.GetIdByAccountName(config.GetDB(), username)
	domains.UserId = id
	order_id, _ := strconv.ParseInt(reqInfo.OrderId, 10, 64)
	domains.Domain = reqInfo.Domain
	domains.DDoSDDId = &order_id
	order_uuid := order.GetUuidById(config.GetDB(), order_id)

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := model.CreateDomains(tx, &domains); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "",
		})
	}
	domain_id := domain.GetIdByOIAndDomain(tx, order_id, domains.Domain, "ddos")

	if err := d.AddDomains(domains.Domain, order_uuid, order.ProType); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": nil,
		})
	}

	if err := d.GetDomainsLists(order_uuid, order.ProType); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": nil,
		})
	}

	for i, _ := range d.List {
		if d.List[i].Domain == domains.Domain {
			if err := d.GetDomainInfo(order_uuid, d.List[i].DomainUUID, order.ProType); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    http.StatusBadRequest,
					"message": err,
				})
				tx.Rollback()
			} else {
				if err := domain.UpdateDomain(tx, &d, domains.Domain, "ddos", order_id); err != nil {
					ctx.JSON(http.StatusOK, gin.H{
						"code":    http.StatusBadRequest,
						"message": err,
					})
					tx.Rollback()
				}
				for i, _ := range d.ConfigList {
					clInfo.LoadBalancing = d.ConfigList[i].LoadBalancing
					clInfo.OverloadRedirectUrl = d.ConfigList[i].OverloadRedirectURL
					clInfo.OverloadStatusCode = *d.ConfigList[i].OverloadStatusCode
					clInfo.OverloadType = d.ConfigList[i].OverloadType
					clInfo.Port = d.ConfigList[i].Port
					clInfo.Protocol = d.ConfigList[i].Protocol
					clInfo.Redirect = d.ConfigList[i].Redirect
					clInfo.Server = d.ConfigList[i].Server
					clInfo.UriForward = *d.ConfigList[i].URIForward
					if err, _ := model.CreateCL(tx, &clInfo, domain_id); err != nil {
						ctx.JSON(http.StatusOK, gin.H{
							"code":    http.StatusBadRequest,
							"message": err,
						})
						tx.Rollback()
					}
					config_list_id := info.GetIdByDomainId(tx, domain_id)
					for j, _ := range d.ConfigList[i].SourceAddresses {
						saInfo.ConfigListId = config_list_id
						saInfo.Address = d.ConfigList[i].SourceAddresses[j].Address
						saInfo.Concurrent = d.ConfigList[i].SourceAddresses[j].Concurrent
						saInfo.Port = d.ConfigList[i].SourceAddresses[j].Port
						saInfo.Protocol = d.ConfigList[i].SourceAddresses[j].Protocol
						saInfo.Sni = d.ConfigList[i].SourceAddresses[j].Sni
						saInfo.Weight = d.ConfigList[i].SourceAddresses[j].Weight
					}
					if err := model.CreateSA(tx, &saInfo, config_list_id); err != nil {
						ctx.JSON(http.StatusOK, gin.H{
							"code":    http.StatusBadRequest,
							"message": err,
						})
						tx.Rollback()
					}
				}
			}
		}
	}

	tx.Commit()
}

func GetDDoSDomains(ctx *gin.Context) {
	var reqPages form.Filter
	var info model.Domain
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
	id := user.GetIdByAccountName(config.GetDB(), username)
	order_id, _ := strconv.ParseInt(originPage.OrderId, 10, 64)

	if lists, total, err := info.GetDomainsByUser(config.GetDB(), &reqPages, id, order_id, "ddos"); err != nil {
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
