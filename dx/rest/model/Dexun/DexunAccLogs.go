package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxAccLogs struct {
	Id                 int64       `gorm:"primary_key" gorm:"column:id" json:"id"`
	CacheHit           string      `gorm:"column:cachehit" json:"cachehit"`
	ClientIp           string      `gorm:"column:clientip" json:"clientip"`
	ClientPort         int64       `gorm:"column:clientport" json:"clientport"`
	ClientRegion       string      `gorm:"column:clientregion" json:"clientregion"`
	Count              int64       `gorm:"column:count" json:"count"`
	CreateDat          interface{} `gorm:"column:createdat" json:"createdat"`
	Domain             string      `gorm:"column:domain" json:"domain"`
	DomainId           string      `gorm:"column:domainid" json:"domainid"`
	Form               string      `gorm:"column:form" json:"form"`
	DxId               int64       `gorm:"column:dx_id" json:"dx_id"`
	InstanceId         int64       `gorm:"column:instanceid" json:"instanceid"`
	LocalAddr          string      `gorm:"column:localaddr" json:"localaddr"`
	LocalIp            string      `gorm:"column:localip" json:"localip"`
	LocalPort          int64       `gorm:"column:localport" json:"localport"`
	Method             string      `gorm:"column:method" json:"method"`
	Nodeid             string      `gorm:"column:nodeid" json:"nodeid"`
	Packagesize        int64       `gorm:"column:packagesize" json:"packagesize"`
	Remoteaddr         string      `gorm:"column:remoteaddr" json:"remoteaddr"`
	Responsesize       int64       `gorm:"column:responsesize" json:"responsesize"`
	Responsestatuscode int64       `gorm:"column:responsestatuscode" json:"responsestatuscode"`
	Timerangeend       string      `gorm:"column:timerangeend" json:"timerangeend"`
	Timerangestart     string      `gorm:"column:timerangestart" json:"timerangestart"`
	URL                string      `gorm:"column:url" json:"url"`
	Wblist             string      `gorm:"column:wblist" json:"wblist"`

	Accept                  string  `gorm:"column:Accept" json:"Accept"`
	AcceptEncoding          string  `gorm:"column:Accept-Encoding" json:"accept-encoding"`
	AcceptLanguage          string  `gorm:"column:Accept-Language" json:"accept-language"`
	Authorization           string  `gorm:"column:Authorization" json:"authorization"`
	CacheControl            string  `gorm:"column:Cache-Control" json:"cache-control"`
	Connection              string  `gorm:"column:Connection" json:"connection"`
	Pragma                  string  `gorm:"column:Pragma" json:"pragma"`
	Purpose                 *string `gorm:"column:Purpose" json:"purpose,omitempty"`
	Referer                 string  `gorm:"column:Referer" json:"referer"`
	UpgradeInsecureRequests *string `gorm:"column:Upgrade-Insecure-Requests" json:"upgrade-insecure-requests,omitempty"`
	UserAgent               string  `gorm:"column:User-Agent" json:"user-agent"`
	XForwardedHost          string  `gorm:"column:X-Forwarded-Host" json:"x-forwarded-host"`
	XForwardedPort          string  `gorm:"column:X-Forwarded-Port" json:"x-forwarded-port"`
	XForwardedProto         string  `gorm:"column:X-Forwarded-Porto" json:"x-forwarded-proto"`
	XForwardedServer        string  `gorm:"column:X-Forwarded-Server" json:"x-forwarded-server"`
	XRealIP                 string  `gorm:"column:X-Real-Ip" json:"x-real-ip"`

	AcceptRanges    *string `gorm:"column:Accept-Ranges" json:"accept-ranges,omitempty"`
	ContentEncoding string  `gorm:"column:Content-Encoding" json:"content-encoding"`
	ContentLength   string  `gorm:"column:Content-Length" json:"content-length"`
	ContentType     string  `gorm:"column:Content-Type" json:"content-type"`
	Date            string  `gorm:"Date" json:"date"`
	Etag            string  `gorm:"column:Etag" json:"etag"`
	LastModified    string  `gorm:"column:Last-Modified" json:"Last-Modified"`
	Server          string  `gorm:"column:Server" json:"server"`
	Vary            string  `gorm:"column:Vary" json:"vary"`
	WWWAuthenticate string  `gorm:"column:Www-Authenticate" json:"www-authenticate"`
}

func (DxAccLogs) YableName() string {
	return "DexunAccLogs"
}

func AddAccLogs(db *gorm.DB, info *form.AccLogsInfo) error {
	items := db.Table("DexunAtkLogs")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxAccLogs{
		Id:                      sf.Generate(),
		CacheHit:                info.CacheHit,
		ClientIp:                info.ClientIp,
		ClientPort:              info.ClientPort,
		ClientRegion:            info.ClientRegion,
		Count:                   info.Count,
		CreateDat:               info.CreateDat,
		Domain:                  info.Domain,
		DomainId:                info.DomainId,
		Form:                    info.Form,
		DxId:                    info.DxId,
		InstanceId:              info.InstanceId,
		LocalAddr:               info.LocalAddr,
		LocalIp:                 info.LocalIp,
		LocalPort:               info.LocalPort,
		Method:                  info.Method,
		Nodeid:                  info.Nodeid,
		Packagesize:             info.Packagesize,
		Remoteaddr:              info.Remoteaddr,
		Responsesize:            info.Responsesize,
		Responsestatuscode:      info.Responsestatuscode,
		Timerangeend:            info.Timerangeend,
		Timerangestart:          info.Timerangestart,
		URL:                     info.URL,
		Wblist:                  info.Wblist,
		Accept:                  info.Accept,
		AcceptEncoding:          info.AcceptEncoding,
		AcceptLanguage:          info.AcceptLanguage,
		Authorization:           info.Authorization,
		CacheControl:            info.CacheControl,
		Connection:              info.Connection,
		Pragma:                  info.Pragma,
		Purpose:                 info.Purpose,
		Referer:                 info.Referer,
		UpgradeInsecureRequests: info.UpgradeInsecureRequests,
		UserAgent:               info.UserAgent,
		XForwardedHost:          info.XForwardedHost,
		XForwardedPort:          info.XForwardedPort,
		XForwardedProto:         info.XForwardedProto,
		XForwardedServer:        info.XForwardedServer,
		XRealIP:                 info.XRealIP,
		AcceptRanges:            info.AcceptRanges,
		ContentEncoding:         info.ClientRegion,
		ContentLength:           info.ContentLength,
		ContentType:             info.ContentType,
		Date:                    info.Date,
		Etag:                    info.Etag,
		LastModified:            info.LastModified,
		Server:                  info.Server,
		Vary:                    info.Vary,
		WWWAuthenticate:         info.WWWAuthenticate,
	}).Error
}
