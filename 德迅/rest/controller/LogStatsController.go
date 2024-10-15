package controller

import (
	"Dexun/config"
	"Dexun/form"
	"Dexun/model"
	"Dexun/model/Dexun"
	"fmt"
	"time"
)

func GetDexunAtkLogs() {
	var service model.ScdnService
	var reqInfo form.AtkLogsInfo
	var reqBody form.Log

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.End = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.Start = int64(int(result.Unix()))
		reqBody.Limit = 100
		reqBody.Page = 1

		if err := d.GetAtkLog(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.List {
			reqInfo.AttackInfo = list.Attackinfo
			reqInfo.AttackType = list.Attacktype
			reqInfo.ClientIp = list.Clientip
			reqInfo.ClientPort = list.Clientport
			reqInfo.ClientRegion = list.Clientregion
			reqInfo.Count = list.Count
			reqInfo.Domain = list.Domain
			reqInfo.DomainId = list.Domainid
			reqInfo.OrderId = order.Id
			reqInfo.HttpMethod = list.Httpmethod
			reqInfo.AlId = list.ID
			reqInfo.InstanceId = list.Instanceid
			reqInfo.LocalIp = list.Localip
			reqInfo.Method = list.Method
			reqInfo.ProtectType = list.Protecttype
			reqInfo.RequestInfo = list.Requestinfo
			reqInfo.TargetUrl = list.Targeturl
			reqInfo.TimeRangeEnd = list.Timerangeend
			reqInfo.TimeRangeStart = list.Timerangestart

			if err := Dexun.AddAtkLogs(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetDexunAccLogs() {
	var service model.ScdnService
	var reqInfo form.AccLogsInfo
	var reqBody form.Log

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.End = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.Start = int64(int(result.Unix()))
		reqBody.Limit = 100
		reqBody.Page = 1

		if err := d.GetAccLog(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.List {
			reqInfo.CacheHit = list.Cachehit
			reqInfo.ClientIp = list.Clientip
			reqInfo.ClientPort = list.Clientport
			reqInfo.ClientRegion = list.Clientregion
			reqInfo.Count = list.Count
			reqInfo.CreateDat = list.Createdat
			reqInfo.Domain = list.Domain
			reqInfo.DomainId = list.Domainid
			reqInfo.Form = list.Form
			reqInfo.DxId = list.ID
			reqInfo.InstanceId = list.Instanceid
			reqInfo.LocalAddr = list.Localaddr
			reqInfo.LocalIp = list.Localip
			reqInfo.LocalPort = list.Localport
			reqInfo.Method = list.Method
			reqInfo.Nodeid = list.Nodeid
			reqInfo.Packagesize = list.Packagesize
			reqInfo.Remoteaddr = list.Remoteaddr
			reqInfo.Responsesize = list.Responsesize
			reqInfo.Responsestatuscode = list.Responsestatuscode
			reqInfo.Timerangeend = list.Timerangeend
			reqInfo.Timerangestart = list.Timerangestart
			reqInfo.URL = list.URL
			reqInfo.Wblist = list.Wblist
			reqInfo.Accept = list.Requestheaders.Accept
			reqInfo.AcceptEncoding = list.Requestheaders.AcceptEncoding
			reqInfo.AcceptLanguage = list.Requestheaders.AcceptLanguage
			reqInfo.Authorization = list.Requestheaders.Authorization
			reqInfo.CacheControl = list.Requestheaders.CacheControl
			reqInfo.Connection = list.Requestheaders.Connection
			reqInfo.Pragma = list.Requestheaders.Pragma
			reqInfo.Purpose = list.Requestheaders.Purpose
			reqInfo.Referer = list.Requestheaders.Referer
			reqInfo.UpgradeInsecureRequests = list.Requestheaders.UpgradeInsecureRequests
			reqInfo.UserAgent = list.Requestheaders.UserAgent
			reqInfo.XForwardedHost = list.Requestheaders.XForwardedHost
			reqInfo.XForwardedPort = list.Requestheaders.XForwardedPort
			reqInfo.XForwardedProto = list.Requestheaders.XForwardedProto
			reqInfo.XForwardedServer = list.Requestheaders.XForwardedServer
			reqInfo.XRealIP = list.Requestheaders.XRealIP
			reqInfo.AcceptRanges = list.Responseheaders.AcceptRanges
			reqInfo.ContentEncoding = list.Responseheaders.ContentEncoding
			reqInfo.ContentLength = list.Responseheaders.ContentLength
			reqInfo.ContentType = list.Responseheaders.ContentType
			reqInfo.Date = list.Responseheaders.Date
			reqInfo.Etag = list.Responseheaders.Etag
			reqInfo.LastModified = list.Responseheaders.LastModified
			reqInfo.Server = list.Responseheaders.Server
			reqInfo.Vary = list.Responseheaders.Vary
			reqInfo.WWWAuthenticate = list.Responseheaders.WWWAuthenticate

			if err := Dexun.AddAccLogs(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetDexunFlowStats() {
	var service model.ScdnService
	var reqInfo form.FlowLogsInfo
	var reqBody form.OtherLog

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetFlowLog(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.List {
			reqInfo.OrderId = order.Id
			reqInfo.Domain = list.Domain
			reqInfo.RequestSize = list.RequestSize
			reqInfo.ResponseSize = list.ResponseSize

			if err := Dexun.AddFlowLog(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetDexunQueryStats() {
	var service model.ScdnService
	var reqInfo form.QueryLogsInfo
	var reqBody form.QueryLog

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))
		reqBody.Limit = 100
		reqBody.Page = 1

		if err := d.GetQueryLog(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.List {
			reqInfo.TotalCalls = list.TotalCalls
			reqInfo.Domain = list.Domain
			reqInfo.CacheRate = list.CacheRate
			reqInfo.CacheCalls = list.CacheCalls

			if err := Dexun.AddQueryLog(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetDexunInterceptStats() {
	var service model.ScdnService
	var reqInfo form.InterceptStatsInfo
	var reqBody form.AtkKind

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.Unit = "hour"
		reqBody.Interval = 1
		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetAtkKind(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		reqInfo.OrderId = order.Id
		reqInfo.AppCC = d.AppCC
		reqInfo.CC = d.CC
		reqInfo.IpBlack = d.IPBlack
		reqInfo.Referer = d.Referer
		reqInfo.UrlBlack = d.UrLBlack
		reqInfo.WebProtect = d.WebProtect
		reqInfo.Other = d.Other
		reqInfo.AreaAcc = d.AreaAcc
		reqInfo.SafeAcc = d.SafeAcc
		reqInfo.PreAcc = d.PreAcc

		if err := Dexun.AddInterceptStats(config.GetDB(), &reqInfo); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func GetDexunAtkStats() {
	var service model.ScdnService
	var reqInfo form.AtkStatsInfo
	var reqBody form.AtkKind

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.Unit = "hour"
		reqBody.Interval = 1
		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetAtkCount(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.List {
			reqInfo.OrderId = order.Id
			reqInfo.Time = list.Time
			reqInfo.TotalCount = list.TotalCount

			if err := Dexun.AddAtkStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetAreaStats() {
	var service model.ScdnService
	var reqInfo form.AreaStatsInfo
	var reqBody form.AtkKind

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.Unit = "hour"
		reqBody.Interval = 1
		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetAreaCount(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, data := range d.MultiData {
			reqInfo.OrderId = order.Id
			reqInfo.Source = data.Source
			reqInfo.Count = data.Count

			if err := Dexun.AddAreaStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetAreaRankStats() {
	var service model.ScdnService
	var reqInfo form.AreaRankStatsInfo
	var reqBody form.AtkKind

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.Unit = "hour"
		reqBody.Interval = 1
		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetDomainAccRanking(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.MultiData {
			reqInfo.OrderId = order.Id
			reqInfo.TotalCount = list.TotalCount
			reqInfo.Domain = list.Domain

			if err := Dexun.AddAreaRankStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetAtkInterStats() {
	var service model.ScdnService
	var reqInfo form.AtkInterStatsInfo
	var reqBody form.AtkKind

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.Unit = "hour"
		reqBody.Interval = 1
		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetAtkDomain(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.MultiData {
			reqInfo.OrderId = order.Id
			reqInfo.Domain = list.Domain
			reqInfo.TotalCount = list.TotalCount

			if err := Dexun.AddAtkInterStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetHttpPackStat() {
	var service model.ScdnService
	var reqInfo form.HttpPackStatsInfo
	var reqBody form.AtkKind

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.Unit = "hour"
		reqBody.Interval = 1
		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetHttpPack(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.MultiData {
			reqInfo.OrderId = order.Id
			reqInfo.Time = list.Time
			reqInfo.TotalCount = list.TotalCount
			if err := Dexun.AddHttpPackStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetLineChartStats() {
	var service model.ScdnService
	var reqInfo form.LineChartStatsInfo
	var reqBody form.Info

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		if err := d.GetFlowLineChart(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, data := range d.MultiData {
			reqInfo.OrderId = order.Id
			reqInfo.Time = data.Time
			reqInfo.RequestSize = data.RequestSize
			reqInfo.ResponseSize = data.ResponseSize

			if err := Dexun.AddLineChartStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetBWStats() {
	var service model.ScdnService
	var reqInfo form.BWStatsInfo
	var reqBody form.Info

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		if err := d.GetBWList(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, list := range d.List {
			reqInfo.OrderId = order.Id
			reqInfo.IP = list.IP
			reqInfo.BwListType = list.BWListType
			reqInfo.TotalCount = list.TotalCount

			if err := Dexun.AddBWStats(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func GetTotalFlow() {
	var service model.ScdnService
	var reqInfo form.TotalFlowInfo
	var reqBody form.Info

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		if err := d.GetTotalFlow(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		reqInfo.OrderId = order.Id
		reqInfo.RequestBandWidthPeak = d.RequestBandwidthPeak
		reqInfo.Requests = d.Requests
		reqInfo.ResponseBandWidthPeak = d.ResponseBandwidthPeak
		reqInfo.TotalRequestFlows = d.TotalRequestFlows
		reqInfo.TotalResponseFlows = d.TotalResponseFlows
		reqInfo.UnidentifiedAttack = d.UnidentifiedAttack

		if err := Dexun.AddTotalFlow(config.GetDB(), &reqInfo); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func GetAccCDNIPRank() {
	var service model.ScdnService
	var reqInfo form.AccCDNRankInfo
	var reqBody form.IpRanking

	orders := service.GetAllRecord(config.GetDB())

	for _, order := range *orders {
		reqBody.DDUUID = service.GetUuidById(config.GetDB(), order.Id)
		reqBody.ProType = service.GetProTypeById(config.GetDB(), order.Id)

		current := time.Now()

		reqBody.EndTime = int64(int(current.Unix()))
		m, _ := time.ParseDuration("-1h")
		result := current.Add(m)
		reqBody.StartTime = int64(int(result.Unix()))

		if err := d.GetIPRanking(&reqBody); err != nil {
			fmt.Println(err)
			continue
		}

		for _, data := range d.MultiData {
			reqInfo.OrderId = order.Id
			reqInfo.CountSum = data.CountSum
			reqInfo.ClientIp = data.ClientIP

			if err := Dexun.AddAccCDNRank(config.GetDB(), &reqInfo); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
