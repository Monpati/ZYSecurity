package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	Old = flag.String("index", "/opt/data/cert", "index")
	New = flag.String("ip", "http://localhost:8088/cert", "ip")
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
	if err := userInfo.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"err":  err,
		})
	}
	if userInfo.GetCertType(config.GetDB(), userInfo.Id) != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "您已完成实名认证",
		})
	}
	corpInfo.UserId = userInfo.GetIdByAccountName(config.GetDB(), username)

	d.GetCorpCertDetails(corpInfo.CorpName, corpInfo.RegNum, corpInfo.LgMan, corpInfo.LgPersonNum, corpInfo.CorpAddress)
	if corpInfo.LgPersonNum != d.SmrzGnum || corpInfo.LgMan != d.FullName || corpInfo.CorpName != d.SmrzCname || corpInfo.RegNum != d.SmrzCnum {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "实名信息错误！",
		})
	}

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := model.CorpCertCreate(tx, &corpInfo); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := userInfo.UpdateCertType(tx, "corp", corpInfo.UserId); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "提交成功，等待审核",
	})
}

func PersonalCert(ctx *gin.Context) {
	var personInfo model.PersonalCert
	var userInfo model.Account
	username := ctx.GetHeader("username")
	id := userInfo.GetIdByAccountName(config.GetDB(), username)

	if err := ctx.BindJSON(&personInfo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	if err := userInfo.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"err":  err,
		})
	}

	if userInfo.GetCertType(config.GetDB(), userInfo.Id) != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "您已完成实名认证",
		})
	}

	d.GetPersonCertDetails(personInfo.RealName, personInfo.CardId)
	if personInfo.CardId != d.SmrzGnum || personInfo.RealName != d.FullName {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "真实姓名或身份证号码错误！",
		})
	}
	personInfo.Birthday = d.SmrzBirthday
	personInfo.City = d.SmrzCity
	personInfo.Sex = d.SmrzSex

	tx := config.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Internal Server Error",
			})
		}
	}()

	if err := model.PersonalCertCreate(tx, &personInfo); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := userInfo.UpdateCertType(tx, "person", id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "提交成功，等待审核",
	})
}

func CardCert(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	path := model.UploadCard(file)
	imageurl := strings.Replace(path, *Old, *New, -1)
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
	var userInfo model.Account

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
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

	if reqBody.Status == 0 {
		if err := user.UpdateStatus(tx, id, reqBody.Status); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}

		if err := userInfo.UpdateCertType(tx, "", id); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}

	} else if reqBody.Status == 1 {
		if err := user.UpdateStatus(config.GetDB(), id, reqBody.Status); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Updated",
	})
}

func UpdateCorpCertStatus(ctx *gin.Context) {
	var corp model.CorpCert
	var reqBody struct {
		Status int `json:"status"`
	}
	var userInfo model.Account

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
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

	if reqBody.Status == 0 {
		if err := corp.UpdateStatus(tx, id, reqBody.Status); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}

		if err := userInfo.UpdateCertType(tx, "", id); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
	} else if reqBody.Status == 1 {
		if err := corp.UpdateStatus(tx, id, reqBody.Status); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Updated",
	})
}

func GetPersonStatus(ctx *gin.Context) {
	var personInfo model.PersonalCert
	var user model.Account
	username := ctx.GetHeader("username")

	if err := user.GetIdByName(config.GetDB(), username); err == nil {
		if err := personInfo.GetStatusById(config.GetDB(), user.Id); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   0,
				"status": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   0,
				"status": personInfo.Status,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"status": err,
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
				"code":   0,
				"status": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   0,
				"status": corpInfo.Status,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"status": err,
		})
	}
}

func PersonalCertDelete(ctx *gin.Context) {
	var reqPersonCert model.PersonalCert
	var reqAccount model.Account

	username := ctx.GetHeader("username")
	id := reqAccount.GetIdByAccountName(config.GetDB(), username)
	if err := reqPersonCert.DeletedById(config.GetDB(), id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"err":  err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
		})
	}
}

func CorpCertDelete(ctx *gin.Context) {
	var reqCorpCert model.CorpCert
	var reqAccount model.Account

	username := ctx.GetHeader("username")
	id := reqAccount.GetIdByAccountName(config.GetDB(), username)

	if err := reqCorpCert.DeletedById(config.GetDB(), id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"err":  err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
		})
	}
}
