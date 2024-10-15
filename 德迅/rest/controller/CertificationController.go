package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCertification(ctx *gin.Context) {
	var reqBody *form.CertificationInfo
	var info form.Certification
	var domain model.Domain
	var order model.ScdnService
	var cert model.Certification

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
	reqBody.OrderId = order_id
	reqBody.DomainId = domain_id
	reqBody.DomainUuid = domain_uuid

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err, cert_id := model.CreateCertification(tx, reqBody)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	info.CERT = reqBody.Cert
	info.CERTName = reqBody.CertName
	info.DDUUID = order_uuid
	info.DomainUUID = reqBody.DomainUuid
	info.Status = reqBody.Status
	info.SSLAlways = reqBody.SslAlways
	info.Hsts = reqBody.Hsts
	info.Key = reqBody.Key
	info.ProType = strconv.FormatInt(pro_type, 10)

	if err := d.CertUpdate(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	if err := d.CertGetDomains(&info); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	reqBody.Desc = d.Desc
	reqBody.Createtime = d.Createtime
	reqBody.UpdateTime = d.Updatetime

	if err := model.CreateDexunCert(tx, d); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	if err := cert.UpdateCertification(tx, cert_id, reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	}

	tx.Commit()
}

func UpdateCertification(ctx *gin.Context) {
	var reqBody *form.CertificationInfo
	var cm model.Certification

	id, _ := strconv.ParseInt(ctx.Param("cert"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": nil,
		})
	}

	if err := cm.UpdateCertification(config.GetDB(), id, reqBody); err != nil {
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

func CertificationLists(ctx *gin.Context) {
	var reqPages form.CertificationFilterForm
	var list model.Certification
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if lists, total, err := list.GetCertificationLists(config.GetDB(), &reqPages); err != nil {
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
