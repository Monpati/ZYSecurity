package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxFlowLog struct {
	Id           int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId      int64  `gorm:"column:domain_id" json:"domain_id"`
	Domain       string `gorm:"column:domain" json:"domain"`
	RequestSize  int64  `gorm:"column:request_size" json:"request_size"`
	ResponseSize int64  `gorm:"column:response_size" json:"response_size"`
}

func (DxFlowLog) TableName() string {
	return "DexunDomainFlow"
}

func AddFlowLog(db *gorm.DB, info *form.FlowLogsInfo) error {
	items := db.Table("DexunDomainFlow")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxFlowLog{
		Id:           sf.Generate(),
		OrderId:      info.OrderId,
		Domain:       info.Domain,
		RequestSize:  info.RequestSize,
		ResponseSize: info.ResponseSize,
	}).Error
}
