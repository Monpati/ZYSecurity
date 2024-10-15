package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxAtkLog struct {
	Id             int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	AttackInfo     string `gorm:"column:attackinfo" json:"attackinfo"`
	AttackType     string `gorm:"column:attacktype" json:"attacktype"`
	ClientIp       string `gorm:"column:clientip" json:"clientip"`
	ClientPort     int64  `gorm:"column:clientport" json:"clientport"`
	ClientRegion   string `gorm:"column:clientregion" json:"clientregion"`
	Count          int64  `gorm:"column:count" json:"count"`
	Domain         string `gorm:"column:domain" json:"domain"`
	DomainId       string `gorm:"column:domainid" json:"domainid"`
	OrderId        int64  `gorm:"column:order_id" json:"order_id"`
	HttpMethod     string `gorm:"column:httpmethod" json:"httpmethod"`
	AlId           int64  `gorm:"column:al_id" json:"al_id"`
	InstanceId     int64  `gorm:"column:instanceid" json:"instanceid"`
	LocalIp        string `gorm:"column:localip" json:"localip"`
	Method         string `gorm:"column:method" json:"method"`
	NodeId         string `gorm:"column:nodeid" json:"nodeid"`
	ProtectType    string `gorm:"column:protecttype" json:"protecttype"`
	RequestInfo    string `gorm:"column:requestinfo" json:"requestinfo"`
	TargetUrl      string `gorm:"column:targeturl" json:"targeturl"`
	TimeRangeEnd   string `gorm:"column:timerangeend" json:"timerangeend"`
	TimeRangeStart string `gorm:"column:timerangestart" json:"timerangestart"`
}

func (DxAtkLog) TableName() string {
	return "DexunAtkLogs"
}

func AddAtkLogs(db *gorm.DB, info *form.AtkLogsInfo) error {
	items := db.Table("DexunAtkLogs")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxAtkLog{
		Id:             sf.Generate(),
		AttackInfo:     info.AttackInfo,
		AttackType:     info.AttackType,
		ClientIp:       info.ClientIp,
		ClientPort:     info.ClientPort,
		ClientRegion:   info.ClientRegion,
		Count:          info.Count,
		Domain:         info.Domain,
		DomainId:       info.DomainId,
		OrderId:        info.OrderId,
		HttpMethod:     info.HttpMethod,
		AlId:           info.AlId,
		InstanceId:     info.InstanceId,
		LocalIp:        info.LocalIp,
		Method:         info.Method,
		ProtectType:    info.ProtectType,
		RequestInfo:    info.RequestInfo,
		TargetUrl:      info.TargetUrl,
		TimeRangeEnd:   info.TimeRangeEnd,
		TimeRangeStart: info.TimeRangeStart,
	}).Error
}
