package controller

import (
	"Dexun/config"
	"Dexun/model"
	"Dexun/model/Dexun"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"time"
)

func GetTotalAccounts(ctx *gin.Context) {
	var agent model.Agent
	var account model.Account

	username := ctx.GetHeader("Username")
	if err := agent.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err,
		})
	}

	_, total := account.GetAccountsByAgentId(config.GetDB(), agent.Id)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": total,
	})
}

func GetSellStats(ctx *gin.Context) {

}

func GetExistOrders(ctx *gin.Context) {

}

func GetTotalFlowStats(ctx *gin.Context) {
	var user model.Account
	var order model.ScdnService
	var data Dexun.DxTotalFlow
	var entries []struct {
		Name  string `json:"name"`
		Value int64  `json:"value"`
	}

	username := ctx.GetHeader("Username")
	if err := user.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	orders := order.GetIdByUserId(config.GetDB(), user.Id)

	for _, tmp := range *orders {
		if err := data.GetDataByOrderId(config.GetDB(), tmp.Id); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
		val := reflect.ValueOf(data)
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldValue := val.Field(i).Interface()

			if fieldName == "Id" || fieldName == "OrderId" || fieldName == "CreateTime" || fieldName == "TotalRequestFlows" || fieldName == "TotalResponseFlows" || fieldName == "Requests" || fieldName == "UnidentifiedAttack" {
				continue
			}

			if fieldName == "RequestBandWidthPeak" {
				fieldName = "上行带宽峰值"
			}
			if fieldName == "ResponseBandWidthPeak" {
				fieldName = "下行带宽峰值"
			}

			entry := struct {
				Name  string `json:"name"`
				Value int64  `json:"value"`
			}{
				Name:  fieldName,
				Value: fieldValue.(int64),
			}
			entries = append(entries, entry)
		}
	}

	if len(*orders) == 0 {
		tmp := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "上行带宽峰值",
			Value: 0,
		}
		entries = append(entries, tmp)
		tmp2 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "下行带宽峰值",
			Value: 0,
		}
		entries = append(entries, tmp2)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": entries,
	})
}

func GetAttackStats(ctx *gin.Context) {

}

func GetInterceptStats(ctx *gin.Context) {
	var user model.Account
	var order model.ScdnService
	var data Dexun.DxInterceptStats
	var entries []struct {
		Name  string `json:"name"`
		Value int64  `json:"value"`
	}

	username := ctx.GetHeader("Username")
	if err := user.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	orders := order.GetIdByUserId(config.GetDB(), user.Id)

	for _, tmp := range *orders {
		if err := data.GetDataByOrderId(config.GetDB(), tmp.Id); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err,
			})
		}
		val := reflect.ValueOf(data)
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldValue := val.Field(i).Interface()

			if fieldName == "Id" || fieldName == "OrderId" || fieldName == "CreateTime" {
				continue
			}

			entry := struct {
				Name  string `json:"name"`
				Value int64  `json:"value"`
			}{
				Name:  fieldName,
				Value: fieldValue.(int64),
			}
			entries = append(entries, entry)
		}

	}

	if len(*orders) == 0 {
		tmp := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "APP专用防CC策略",
			Value: 0,
		}
		entries = append(entries, tmp)
		tmp2 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "CC防护",
			Value: 0,
		}
		entries = append(entries, tmp2)
		tmp3 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "IP黑名单",
			Value: 0,
		}
		entries = append(entries, tmp3)
		tmp4 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "Referer防盗链",
			Value: 0,
		}
		entries = append(entries, tmp4)
		tmp5 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "URL黑名单",
			Value: 0,
		}
		entries = append(entries, tmp5)
		tmp6 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "Web攻击防护",
			Value: 0,
		}
		entries = append(entries, tmp6)
		tmp7 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "其他",
			Value: 0,
		}
		entries = append(entries, tmp7)
		tmp8 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "区域访问限制",
			Value: 0,
		}
		entries = append(entries, tmp8)
		tmp9 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "安全访问控制",
			Value: 0,
		}
		entries = append(entries, tmp9)
		tmp10 := struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		}{
			Name:  "精准访问控制",
			Value: 0,
		}
		entries = append(entries, tmp10)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": entries,
	})
}

func GetLineChart(ctx *gin.Context) {
	var user model.Account
	var order model.ScdnService
	var data Dexun.DxLineChartStats
	var datas *[]Dexun.DxLineChartStats
	var result interface{}

	username := ctx.GetHeader("Username")
	if err := user.GetIdByName(config.GetDB(), username); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
	}
	orders := order.GetIdByUserId(config.GetDB(), user.Id)

	for _, tmp := range *orders {
		datas = data.GetDataByOrderId(config.GetDB(), tmp.Id)
		result = append(*datas)
	}

	if len(*orders) == 0 {
		data := Dexun.DxLineChartStats{
			Id:           0,
			OrderId:      0,
			Time:         int64(int(time.Now().Unix())),
			ResponseSize: 0,
			RequestSize:  0,
		}
		datas = &[]Dexun.DxLineChartStats{data}
		result = append(*datas)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": result,
	})
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": nil,
	})
}
