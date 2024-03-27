package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func PersonalCertList(ctx *gin.Context) {
	var cert model.PersonalCert
	var reqPages form.PersonFilter
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if certs, total, err := cert.GetByParams(config.GetDB(), &reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "",
			"data": gin.H{
				"certs": certs,
				"total": total,
			},
		})
	}
}

func CorpCertList(ctx *gin.Context) {
	var corp model.CorpCert
	var reqPages form.CorpFilter
	var originPage form.PageOrigin

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}

	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize
	if corps, total, err := corp.GetByParams(config.GetDB(), &reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "",
			"data": gin.H{
				"corps": corps,
				"total": total,
			},
		})
	}
}

func CorpCert(ctx *gin.Context) {
	var corpInfo model.CorpCert
	var userInfo model.Account
	username := ctx.GetHeader("username")
	if err := ctx.BindJSON(&corpInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if userInfo.GetCertType(config.GetDB(), username) != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "您已完成实名认证",
		})
	}
	corpInfo.UserId = userInfo.GetIdByAccountName(config.GetDB(), username)
	//De Xun Api

	if err := model.CorpCertCreate(config.GetDB(), &corpInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		if err := userInfo.UpdateCertType(config.GetDB(), username, "corp"); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "提交成功，等待审核",
			})
		}
	}
}

func PersonalCert(ctx *gin.Context) {
	var personInfo model.PersonalCert
	var userInfo model.Account
	username := ctx.GetHeader("username")
	if err := ctx.BindJSON(&personInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if userInfo.GetCertType(config.GetDB(), username) != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "您已完成实名认证",
		})
	}
	personInfo.UserId = userInfo.GetIdByAccountName(config.GetDB(), username)
	//if err := personInfo.GetCertDetails(personInfo.RealName, personInfo.CardId); err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code":    http.StatusBadRequest,
	//		"message": err,
	//	})
	//}
	if err := model.PersonalCertCreate(config.GetDB(), &personInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		if err := userInfo.UpdateCertType(config.GetDB(), username, "person"); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "提交成功，等待审核",
			})
		}
	}
}

func CardCert(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	path := model.UploadCard(file)
	imageurl := strings.Replace(path, "/opt/data/cert", "http://localhost:8088/cert", -1)
	if imageurl != "" {
		if err := ctx.SaveUploadedFile(file, path); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		} else {
			os.Chmod(path, 0777)
			ctx.JSON(http.StatusOK, gin.H{
				"code":      0,
				"card_path": imageurl,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "未生成路径",
		})
	}
}

func UpdatePersonCertStatus(ctx *gin.Context) {
	var user model.PersonalCert
	var reqBody struct {
		Status int `json:"status"`
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := user.UpdateStatus(config.GetDB(), id, reqBody.Status); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Updated",
		})
	}
}

func UpdateCorpCertStatus(ctx *gin.Context) {
	var corp model.CorpCert
	var reqBody struct {
		Status int `json:"status"`
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := corp.UpdateStatus(config.GetDB(), id, reqBody.Status); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Updated",
		})
	}
}

func GetPersonStatus(ctx *gin.Context) {
	var personInfo model.PersonalCert
	var user model.Account
	username := ctx.GetHeader("username")

	if err := user.GetIdByName(config.GetDB(), username); err == nil {
		if err := personInfo.GetStatusById(config.GetDB(), user.Id); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   0,
				"status": personInfo.Status,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
}

func GetCorpStatus(ctx *gin.Context) {
	var corpInfo model.CorpCert
	var user model.Account
	username := ctx.GetHeader("username")
	if err := user.GetIdByName(config.GetDB(), username); err == nil {
		if err := corpInfo.GetStatusById(config.GetDB(), user.Id); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   0,
				"status": corpInfo.Status,
			})
		}
	}
}
