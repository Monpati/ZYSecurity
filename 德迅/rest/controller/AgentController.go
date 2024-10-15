package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InviteCode(ctx *gin.Context) {
	var reqAgent form.AgentInfo
	code, _ := config.ReleaseAgentCode(reqAgent)
	if err := config.StoreCodeInRedis(config.GetRedis(), code, reqAgent.TelNum); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "系统异常",
		})
	}
}

func AgentRegister(ctx *gin.Context) {
	var reqAgent form.AgentInfo

	if err := ctx.BindJSON(&reqAgent); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := model.AgentCreate(config.GetDB(), &reqAgent); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		storedCode, _ := config.GetCodeFormRedis(config.GetRedis(), reqAgent.TelNum)
		if reqAgent.InviteCode == storedCode {
			ctx.JSON(http.StatusOK, gin.H{
				"code":  200,
				"error": "",
				"data":  gin.H{},
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
	}
}

func AgentLogin(ctx *gin.Context) {
	var reqAgent model.Agent
	var reqAccount model.Account

	if err := ctx.BindJSON(&reqAgent); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	fmt.Println(reqAgent.FindSaltByAgentName(config.GetDB(), reqAgent.Username))
	fmt.Println(reqAgent.Password)
	//hashedPassword, _ := model.HashSalt(reqAccount.Password, reqAccount.FindSaltByAccountName(config.GetDB(), reqAccount.Username))
	reqAccount.Username = reqAgent.Username
	if err := reqAgent.FindAgentByUP(config.GetDB(), reqAgent.Username, reqAgent.Password); err == nil {
		token, _ := config.ReleaseToken(reqAccount)
		if err := config.StoreTokenInRedis(config.GetRedis(), token, reqAgent.Username); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "系统异常",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "登录成功",
				"data": gin.H{
					"username": reqAgent.Username,
					"password": reqAgent.Password,
					"token":    token,
				}})
		}
	} else {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": err,
		})
		return
	}
}

func UpdateAgentStatus(ctx *gin.Context) {
	var agent model.Agent
	var invite model.Invite
	var reqBody struct {
		Code   string `json:"code"`
		Status int    `json:"status"`
	}

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

	if err := agent.UpdateStatus(tx, id, reqBody.Status); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		tx.Rollback()
	} else {
		if reqBody.Status == 1 {
			if err := invite.UpdateInviteStatus(tx, id, 1); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    http.StatusBadRequest,
					"message": err,
				})
				tx.Rollback()
			}
		} else if reqBody.Status == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "Failed",
			})
			tx.Rollback()
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
	})
}

func GetUsers(ctx *gin.Context) {
	var agent model.Agent
	var reqPages form.Filter
	var originPage form.PageOrigin

	agentName := ctx.GetHeader("Username")
	if err := agent.GetIdByName(config.GetDB(), agentName); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if users, total, err := agent.GetUsersByAgent(config.GetDB(), agent.Id, &reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "",
			"data": gin.H{
				"users": users,
				"total": total,
			},
		})
	}
}

func AddInviteCode(ctx *gin.Context) {
	var reqAgent form.InviteCode

	if err := ctx.BindJSON(&reqAgent); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := model.InviteCodeAdd(config.GetDB(), &reqAgent); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err,
		})
	}
}

func UpdateInviteCodeStatus(ctx *gin.Context) {
	var invite model.Invite
	var reqBody struct {
		Status int `json:"status"`
	}

	id, _ := strconv.ParseInt(ctx.Param("invite"), 10, 64)
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}

	if err := invite.UpdateInviteStatus(config.GetDB(), id, reqBody.Status); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": err,
		})
	}
}

func AgentList(ctx *gin.Context) {
	var usr model.Agent
	var reqPages form.AgentFilterForm
	var originPage form.PageOrigin
	//if usr.Role == "admin" {

	if err := ctx.BindJSON(&originPage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "解析错误",
		})
	}
	reqPages.Offset = originPage.PageSize * (originPage.PageIndex - 1)
	reqPages.Limit = originPage.PageSize

	if agents, total, err := usr.GetByParams(config.GetDB(), &reqPages); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "",
			"data": gin.H{
				"agents": agents,
				"total":  total,
			},
		})
	}
}
