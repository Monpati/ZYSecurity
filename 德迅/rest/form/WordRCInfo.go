package form

import "Dexun/utils"

type WordRCInfo struct {
	OrderId    int64      `gorm:"column:order_id" json:"order_id"`
	OrderUuid  string     `gorm:"column:order_uuid" json:"order_uuid"`
	DomainId   int64      `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string     `gorm:"column:domain_uuid" json:"domain_uuid"`
	DwId       int64      `gorm:"column:dw_id" json:"dw_id"`
	Gzip       string     `gorm:"column:gzip" json:"gzip"`
	Active     string     `gorm:"column:active" json:"active"`
	KeyWords   utils.JSON `gorm:"column:keywords" json:"keywords"`
}
