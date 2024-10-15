package controller

import (
	"Dexun/config"
	"Dexun/config/dexun"
	"Dexun/form"
	"Dexun/model"
	"Dexun/model/Dexun"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	d Dexun.DeXunBody
)

func OpenCert() {
	d.SwitchCert()
}

func GetToken() {
	d.Token = dexun.Api()
	fmt.Println(d.Token)
}

func SweepCache(ctx *gin.Context) {
	var reqInfo form.Domain
	var domain model.Domain
	var order model.ScdnService

	if err := ctx.BindJSON(&reqInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	reqInfo.DomainUUID = domain.GetUuidById(config.GetDB(), domain_id)
	order_id := domain.GetOrderIdById(config.GetDB(), domain_id)
	reqInfo.DDUUID = order.GetUuidById(config.GetDB(), order_id)
	reqInfo.ProType = order.GetProTypeById(config.GetDB(), order_id)

	if err := d.CleanCache(reqInfo.DomainUUID, reqInfo.DDUUID, reqInfo.ProType); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "缓存已清除",
		})
	}
}

func GetSCDNLists(ctx *gin.Context) {
	var services model.ScdnService
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

func GetSCDNComboLists() {
	var results []Dexun.Data
	var combo *Dexun.DexunCombos
	results = d.GetComboList()

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, _ := range results {
		if combo = combo.FindComboByUuid(tx, results[i].UUID); combo.UUID != "" {
			if combo.Status == 1 {
				if err := combo.UpdateComboByUuid(tx, combo.UUID, results[i]); err != nil {
					tx.Rollback()
					continue
				}
			}
		} else if combo.UUID == "" {
			if err := Dexun.CreateDexunCombos(tx, results[i]); err != nil {
				tx.Rollback()
				continue
			}
		}
	}

	tx.Commit()
}

func GetPackageLists() {
	var result, result2 []Dexun.Data
	var packages *Dexun.ScdnPackage
	result = d.GetFlowPackageLists()
	result2 = d.GetDDPackageLists()

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, _ := range result {
		if packages = packages.FindPackageByUuid(tx, result[i].UUID); packages.UUID != "" {
			if packages.Status == 1 {
				if err := packages.UpdatePackageByUuid(tx, packages.UUID, result[i]); err != nil {
					tx.Rollback()
					continue
				}
			}
		} else if packages.UUID == "" {
			if err := Dexun.CreateScdnPackage(tx, result[i]); err != nil {
				tx.Rollback()
				continue
			}
		}
	}

	for i, _ := range result2 {
		if packages = packages.FindPackageByUuid(tx, result2[i].UUID); packages.UUID != "" {
			if packages.Status == 1 {
				if err := packages.UpdatePackageByUuid(tx, packages.UUID, result2[i]); err != nil {
					tx.Rollback()
					continue
				}
			}
		} else if packages.UUID == "" {
			if packages.Status == 1 {
				if err := Dexun.CreateScdnPackage(tx, result2[i]); err != nil {
					tx.Rollback()
					continue
				}
			}
		}
	}

	tx.Commit()
}

func AllowDexunScdnSell(ctx *gin.Context) {
	var scdnInfo form.ScdnMove
	if err := ctx.BindJSON(scdnInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := Dexun.CopyDexunCombos(config.GetDB(), scdnInfo.Id); err != nil {
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
}

func UpdateScdnCombosStatus(ctx *gin.Context) {
	var scdn Dexun.ScdnCombos
	var reqBody struct {
		Status int `json:"status"`
	}

	id, _ := strconv.ParseInt(ctx.Param("scdn"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := scdn.ChangeScdnCombosStatus(config.GetDB(), id, reqBody.Status); err != nil {
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

func DexunScdnLists(ctx *gin.Context) {
	var info Dexun.DexunCombos
	var reqPages form.ScdnFilterForm
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if lists, total, err := info.GetDCByParams(config.GetDB(), &reqPages); err != nil {
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

func ScdnLists(ctx *gin.Context) {
	var info Dexun.ScdnCombos
	var reqPages form.ScdnFilterForm
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
		if lists, total, err := info.UserGetSCByParams(config.GetDB(), &reqPages); err != nil {
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
		if lists, total, err := info.AdminGetSCByParams(config.GetDB(), &reqPages); err != nil {
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

func AllowDexunSPSell(ctx *gin.Context) {
	var scdnInfo form.ScdnMove
	if err := ctx.BindJSON(scdnInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := Dexun.CopyPackageService(config.GetDB(), scdnInfo.Id); err != nil {
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
}

func UpdateDexunSPStatus(ctx *gin.Context) {
	var scdn Dexun.PackageService
	var reqBody struct {
		Status int `json:"status"`
	}
	id, _ := strconv.ParseInt(ctx.Param("packages"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := scdn.ChangeScdnPackageStatus(config.GetDB(), id, reqBody.Status); err != nil {
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

func DexunPackageLists(ctx *gin.Context) {
	var info Dexun.ScdnPackage
	var reqPages form.ScdnFilterForm
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if lists, total, err := info.GetDxPackagesByParams(config.GetDB(), &reqPages); err != nil {
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

func PackageLists(ctx *gin.Context) {
	var info Dexun.PackageService
	var reqPages form.ScdnFilterForm
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if lists, total, err := info.GetPackagesByParams(config.GetDB(), &reqPages); err != nil {
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

func DomainsLists(ctx *gin.Context) {
	var info model.Domain
	var reqPages form.DomainFilterForm
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

func UpdateDomainsStatus(ctx *gin.Context) {
	var domain model.Domain
	var order model.ScdnService
	var reqBody form.DomainInfo

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

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := domain.UpdateDomainStatus(tx, domain_id, reqBody.Status); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": nil,
		})
	}

	if err := d.DeleteDomains(order_uuid, pro_type, domain_uuid); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func AddDomain(ctx *gin.Context) {
	var domains form.DomainInfo
	var reqInfo form.DomainsInfo
	var domain model.Domain
	var user model.Account
	var order model.ScdnService
	var info model.ConfigList
	var clInfo form.ConfigListInfo
	var saInfo form.SourceAddressInfo

	username := ctx.GetHeader("username")
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
	domains.OrderId = &order_id
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
	domain_id := domain.GetIdByOIAndDomain(tx, order_id, domains.Domain, "scdn")

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
				if err := domain.UpdateDomain(tx, &d, domains.Domain, "scdn", order_id); err != nil {
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

func AddDomainConfig(ctx *gin.Context) {
	var info form.Config
	var domain model.Domain
	var order model.ScdnService
	var clInfo form.ConfigListInfo
	var saInfo form.SourceAddressInfo
	var dxInfo form.DomainDevices

	if err := ctx.BindJSON(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	domain_uuid := domain.GetUuidById(config.GetDB(), domain_id)
	order_id := domain.GetOrderIdById(config.GetDB(), domain_id)
	order_uuid := order.GetUuidById(config.GetDB(), order_id)
	pro_type := order.GetProTypeById(config.GetDB(), order_id)

	clInfo.DomainId = domain_id
	clInfo.LoadBalancing = info.LoadBalancing
	clInfo.OverloadRedirectUrl = info.OverloadRedirectURL
	clInfo.OverloadStatusCode = *info.OverloadStatusCode
	clInfo.OverloadType = info.OverloadType
	clInfo.Port = info.Port
	clInfo.Protocol = info.Protocol
	clInfo.Redirect = info.Redirect
	clInfo.Server = info.Server
	clInfo.UriForward = *info.URIForward

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err, config_list_id := model.CreateCL(tx, &clInfo, domain_id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	for i, _ := range info.SourceAddresses {
		saInfo.ConfigListId = config_list_id
		saInfo.Address = info.SourceAddresses[i].Address
		saInfo.Concurrent = info.SourceAddresses[i].Concurrent
		saInfo.Port = info.SourceAddresses[i].Port
		saInfo.Protocol = info.SourceAddresses[i].Protocol
		saInfo.Sni = info.SourceAddresses[i].Sni
		saInfo.Weight = info.SourceAddresses[i].Weight
		if err := model.CreateSA(tx, &saInfo, config_list_id); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
	}

	dxInfo.DomainUUID = domain_uuid
	dxInfo.DDUUID = order_uuid
	dxInfo.ProType = pro_type
	dxInfo.Config = append(dxInfo.Config, info)

	if err := d.UpdateHttpDevices(dxInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func AddConfigLists(ctx *gin.Context) {

}

func UpdateConfigLists(ctx *gin.Context) {

}

func ConfigLists(ctx *gin.Context) {

}

func AddSourceAddress(ctx *gin.Context) {

}

func UpdateSourceAddress(ctx *gin.Context) {

}

func SourceAddresses(ctx *gin.Context) {

}

func GetDomainSni(ctx *gin.Context) {
	var originPage form.PageOrigin
	var configlist model.ConfigList
	var reqPages form.ConfigListInfo
	var user model.Account

	domain_id, _ := strconv.ParseInt(ctx.Param("domain"), 10, 64)
	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	username := ctx.GetHeader("Username")
	if err := user.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if domains, total, err := configlist.GetCLParams(config.GetDB(), &reqPages, domain_id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"lists": domains,
				"total": total,
			},
		})
	}
}

func GetSCDNUserDomains(ctx *gin.Context) {
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

	if lists, total, err := info.GetDomainsByUser(config.GetDB(), &reqPages, id, order_id, "scdn"); err != nil {
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

func CleanDomainCache(ctx *gin.Context) {

}
