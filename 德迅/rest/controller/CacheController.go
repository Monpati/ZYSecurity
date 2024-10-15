package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"Dexun/model/Dexun"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CacheAdd(ctx *gin.Context) {
	var reqBody *form.CacheInfo
	var order model.ScdnService
	var domain model.Domain
	var cache model.Cache
	var info form.Cache

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
	reqBody.DomainId = domain_id
	reqBody.OrderId = order_id
	reqBody.DomainUuid = domain_uuid
	reqBody.DdUuid = order_uuid
	reqBody.ProType = strconv.FormatInt(pro_type, 10)
	active, _ := strconv.ParseInt(reqBody.Active, 10, 64)

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err, cache_id := model.CreateCache(tx, reqBody, active)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.Active = active
	info.Cacheextensions = reqBody.CacheExtensions
	info.Cachemode = reqBody.CacheMode
	info.Cachepath = reqBody.CachePath
	info.Cachereg = reqBody.CacheReg
	info.DDUUID = reqBody.DdUuid
	info.DomainUUID = reqBody.DomainUuid
	info.ProType = reqBody.ProType
	info.Timeout = reqBody.TimeOut
	info.Urlmode = reqBody.UrlMode
	info.Weight = reqBody.Weight

	if err := d.CacheAdd(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := d.CacheLists(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	for i, _ := range d.List {
		if err := Dexun.CreateDexunCache(tx, d.List[i]); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if i == 0 {
			reqBody.CacheUuid = d.List[i].CacheUUID
			reqBody.UrlMode = d.List[i].Urlmode
			reqBody.CacheMode = d.List[i].Cachemode
			reqBody.CachePath = d.List[i].Cachepath
			reqBody.CacheExtensions = d.List[i].Cacheextensions
			reqBody.CacheReg = d.List[i].Cachereg
			reqBody.TimeOut = strconv.FormatInt(d.List[i].Timeout, 10)
			reqBody.Weight = strconv.FormatInt(d.List[i].Weight, 10)
			reqBody.CreateTime = d.List[i].Createtime
			reqBody.UpdateTime = d.List[i].Updatetime
			if err := cache.UpdateCache(tx, cache_id, reqBody, active); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    http.StatusBadRequest,
					"message": err,
				})
				tx.Rollback()
			}
		}
	}

	tx.Commit()
}

func CacheLists(ctx *gin.Context) {
	var reqPages form.CacheFilterForm
	var list model.Cache
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetCacheLists(config.GetDB(), &reqPages); err != nil {
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

func UpdateCache(ctx *gin.Context) {
	var reqBody *form.CacheInfo
	var cm model.Cache

	id, _ := strconv.ParseInt(ctx.Param("model"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}
	active, _ := strconv.ParseInt(reqBody.Active, 10, 64)

	if err := cm.UpdateCache(config.GetDB(), id, reqBody, active); err != nil {
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
