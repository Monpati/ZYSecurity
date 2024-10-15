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

func GetDexunSSLList() {
	var results []Dexun.List
	var combo *Dexun.DexunSSLList
	results = d.GetSSLList()

	for i, _ := range results {
		if combo = combo.FindComboByUuid(config.GetDB(), results[i].UUID); combo.UUID != "" {
			if err := combo.UpdateSSLByUUID(config.GetDB(), combo.UUID, results[i]); err != nil {
				continue
			}
		} else if combo.UUID == "" {
			if err := Dexun.CreateDexunSSLList(config.GetDB(), results[i]); err != nil {
				continue
			}
		}
	}
}

func CopyDexunSSLList() {
	if err := Dexun.CopyDxSSLList(config.GetDB()); err != nil {
		fmt.Println(err)
	}
}

func SSLList(ctx *gin.Context) {
	var info Dexun.SSLList
	var reqPages form.Filter
	var originPage form.SSLPageOrigin

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
		if lists, total, err := info.UserGetSSLByParams(config.GetDB(), originPage.SSLType, &reqPages); err != nil {
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
		if lists, total, err := info.AdminGetSSLByParams(config.GetDB(), originPage.SSLType, &reqPages); err != nil {
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

func GetSSLCSR(ctx *gin.Context) {
	var reqInfo form.SSLCsr

	if err := ctx.BindJSON(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := d.GetSSLCSRandKey(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"loadcsr":    d.Loadcsr,
				"loadkey":    d.Loadkey,
				"loaddomain": d.Loaddomain,
			},
		})
	}
}

func UpdateSSLDoaminConfirm(ctx *gin.Context) {
	var reqInfo form.SSLDomain

	if err := ctx.BindJSON(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := d.GetSSLDomainConfirm(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
}

func CreateSSLOrders(ctx *gin.Context) {
	var reqDB form.SSLServiceInfo
	var reqInfo form.SSLOrderInfo
	var reqBody *form.SSLPurchaseInfo
	var reqId form.SSLListInfo
	var info form.SimpleInfo
	var ssl Dexun.SSLList
	var account model.Account

	username := ctx.GetHeader("Username")
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	user_id := account.GetIdByAccountName(config.GetDB(), username)
	id, err := strconv.ParseInt(reqBody.Id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	reqInfo.PUUID = ssl.GetUUIDById(config.GetDB(), id)
	reqInfo.DomainList = reqBody.DomainList
	reqInfo.Loadcsr = reqBody.LoadCsr
	reqInfo.Loaddomain = reqBody.LoadDomain
	reqInfo.Loadkey = reqBody.LoadKey
	reqInfo.Provemethod = reqBody.ProveMethod
	reqInfo.Admininfo = reqBody.LastName + ";" + reqBody.FirstName + ";" + reqBody.TelNum + ";" + reqBody.Email + ";" + reqBody.Job
	reqInfo.Orginfo = reqBody.CsrOrg + ";" + reqBody.Department + ";" + reqBody.CsrState + ";" + reqBody.CsrLocality
	balance := account.GetBalanceById(config.GetDB(), user_id)
	money := ssl.GetMoneyById(config.GetDB(), id)
	if balance < money {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "余额不足!",
		})
	}

	order_uuid := d.CreateSSLOrder(&reqInfo)
	//order_uuid := "60a179b6-1d22-48e5-9d6c-cba5fc18251c"
	orderUUIDs := strings.Fields(order_uuid)
	results := d.GetBillCallBack(orderUUIDs)

	for i, _ := range results {
		if results[i].OrderStatus == 1 {
			info.UUID = order_uuid

			if err := d.GetSSLOrderInfo(&info); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    http.StatusBadRequest,
					"message": err,
				})
			}

			for _, list := range d.List {
				if list.Name == "证书类型" {
					reqDB.SSLTypeDetail = list.Value
				}
				if list.Name == "购买日期" {
					reqDB.Create = list.Value
				}
				if list.Name == "CA证书编号" {
					reqDB.CaNum = list.Value
				}
				if list.Name == "证书安装服务" {
					reqDB.SetUpServiceDetail = list.Value
				}
			}

			for _, list := range d.AdminList {
				if list.Name == "申请人姓名" {
					reqDB.AdminName = list.Value
				}
				if list.Name == "申请人手机" {
					reqDB.AdminTel = list.Value
				}
				if list.Name == "申请人邮箱" {
					reqDB.AdminEmail = list.Value
				}
				if list.Name == "申请人职务" {
					reqDB.AdminJob = list.Value
				}
			}

			reqId.Keywords = reqBody.LoadDomain

			if err := d.GetSSLOrderList(&reqId); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    http.StatusBadRequest,
					"message": err,
				})
			}

			for _, list := range d.List {
				reqDB.UserId = user_id
				reqDB.SSLId = id
				reqDB.AdminInfo = list.AdminInfo
				reqDB.COMStatus = list.COMStatus
				reqDB.DNSHost = list.DNSHost
				reqDB.DNSType = list.DNSType
				reqDB.DNSValue = list.DNSValue
				reqDB.DomainList = list.DomainList
				reqDB.DomainNum = list.DomainNum
				reqDB.DomainType = list.DomainType
				reqDB.FileName = list.FileName
				reqDB.FileValue = list.FileValue
				reqDB.OrderStart = list.OrderStart
				reqDB.OrgInfo = list.OrgInfo
				reqDB.PMethod = list.PMethod
				reqDB.PTypeName = list.PTypeName
				reqDB.SetupServer = list.SetupServer
				reqDB.SSLCode = list.SSLCode
				reqDB.SSLCsr = list.SSLCsr
				reqDB.SSLKey = list.SSLKey
				reqDB.SSLPem = list.SSLPem
				reqDB.SSLType = list.SSLType
				reqDB.TechInfo = list.TechInfo
				reqDB.UUID = list.UUID
				reqDB.XufeiOrderid = list.XufeiOrderid
				reqDB.YnProve = list.YnProve
				reqDB.YnReplace = list.YnReplace
				reqDB.YnXufei = list.YnXufei
				reqDB.ZDomain = list.ZDomain
				reqDB.OrderId = list.OrderID
				reqDB.OrderName = list.OrderName
				reqDB.SSLName = list.SSLName
				reqDB.ZDomainList = list.ZDomainList
				reqDB.KsMoney = list.KsMoney
				reqDB.EdTime = list.EdTime
				reqDB.DateDiff = list.DateDiff
				reqDB.Img = list.Img

				if err := model.CreateSSLService(config.GetDB(), &reqDB); err != nil {
					ctx.JSON(http.StatusOK, gin.H{
						"code":    http.StatusBadRequest,
						"message": err,
					})
				} else {
					ctx.JSON(http.StatusOK, gin.H{
						"code":    0,
						"message": nil,
					})
				}

				if err := account.UpdateBalancePurchase(config.GetDB(), user_id, balance, money); err != nil {
					ctx.JSON(http.StatusOK, gin.H{
						"code":    http.StatusBadRequest,
						"message": err,
					})
				}
			}
		}
	}
}

func GetSSLLists(ctx *gin.Context) {
	var services model.SSLService
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

func GetSSLDetail(ctx *gin.Context) {
	var ssl model.SSLService
	var tmp struct {
		Id string `json:"id"`
	}

	if err := ctx.BindJSON(&tmp); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	order_id, _ := strconv.ParseInt(tmp.Id, 10, 64)
	if err := ssl.GetSSLById(config.GetDB(), order_id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	list := struct {
		SSLTypeDetail      string `json:"ssl_type_detail"`
		SSLName            string `json:"ssl_name"`
		OrderId            string `json:"order_id"`
		Create             string `json:"create"`
		CaNum              string `json:"ca_num"`
		Domain             string `json:"z_domain"`
		SetUpServiceDetail string `json:"setup_service_detail"`
		ZyksMoney          int64  `json:"zyks_money"`
		OrderName          string `json:"order_name"`
		AdminName          string `json:"admin_name"`
		AdminTel           string `json:"admin_tel"`
		AdminEmail         string `json:"admin_email"`
		AdminJob           string `json:"admin_job"`
	}{
		SSLTypeDetail:      ssl.SSLTypeDetail,
		SSLName:            ssl.SSLName,
		OrderId:            ssl.OrderId,
		Create:             ssl.Create,
		CaNum:              ssl.CaNum,
		Domain:             ssl.ZDomain,
		SetUpServiceDetail: ssl.SetUpServiceDetail,
		ZyksMoney:          ssl.ZyksMoney,
		OrderName:          ssl.OrderName,
		AdminName:          ssl.AdminName,
		AdminTel:           ssl.AdminTel,
		AdminEmail:         ssl.AdminEmail,
		AdminJob:           ssl.AdminJob,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"list": list,
	})

}
