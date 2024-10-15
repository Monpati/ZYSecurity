package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxBWStats struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	IP         string `gorm:"column:ip" json:"ip"`
	BwListType string `gorm:"column:bw_list_type" json:"bw_list_type"`
	TotalCount int64  `gorm:"column:total_count" json:"total_count"`
}

func (DxBWStats) TableName() string {
	return "DexunBWStats"
}

func AddBWStats(db *gorm.DB, info *form.BWStatsInfo) error {
	items := db.Table("DexunBWStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxBWStats{
		Id:         sf.Generate(),
		OrderId:    info.OrderId,
		IP:         info.IP,
		BwListType: info.BwListType,
		TotalCount: info.TotalCount,
	}).Error
}
