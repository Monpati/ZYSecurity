package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateWaf(ctx *gin.Context) {
	var reqBody *form.WafInfo
	var domain model.Domain
	var order model.ScdnService
	var info form.WafStatus
	var proxy_info form.WafProxyStatus
	var xss_info form.WafXssStatus

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
	reqBody.DDUUID = order_uuid
	reqBody.DomainUUID = domain_uuid
	reqBody.ProType = strconv.FormatInt(pro_type, 10)
	reqBody.DomainId = domain_id

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	info.DomainUUID = reqBody.DomainUUID
	info.Active = reqBody.Active
	info.DDUUID = reqBody.DDUUID
	info.ProType = reqBody.ProType

	switch reqBody.Type {
	case "WAF":
		if err := domain.UpdateTotalWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}

		if err := d.UpdateTotalWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "File":
		if err := domain.UpdateFileWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateFileWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Code":
		if err := domain.UpdateCodeWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateCodeWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Session":
		if err := domain.UpdateSessionWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateSessionWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Shellshock":
		if err := domain.UpdateShellshockWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateShellShockWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Script":
		if err := domain.UpdateScriptWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateScriptWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Meta":
		if err := domain.UpdateMetaWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateMetaWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Sql":
		if err := domain.UpdateSqlWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateSqlWaf(&info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Proxy":
		proxy_info.ProType = reqBody.ProType
		proxy_info.DomainUUID = reqBody.DomainUUID
		proxy_info.DDUUID = reqBody.DDUUID
		proxy_info.Proxy = reqBody.Active
		if err := domain.UpdateProxyWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateProxyWaf(&proxy_info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
	case "Xss":
		xss_info.ProType = reqBody.ProType
		xss_info.DomainUUID = reqBody.DomainUUID
		xss_info.DDUUID = reqBody.DDUUID
		xss_info.XSS = reqBody.Active
		if err := domain.UpdateXssWaf(tx, reqBody); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}
		if err := d.UpdateXssWaf(&xss_info); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
			tx.Rollback()
		}

	}

	tx.Commit()
}
