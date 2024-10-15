package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxAreaStats struct {
	Id      int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId int64  `gorm:"column:order_id" json:"order_id"`
	Source  string `gorm:"column:source" json:"source"`
	Count   int64  `gorm:"column:count" json:"count"`
}

func (DxAreaStats) TableName() string {
	return "DexunAreaStats"
}

func AddAreaStats(db *gorm.DB, info *form.AreaStatsInfo) error {
	items := db.Table("DexunAreaStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxAreaStats{
		Id:      sf.Generate(),
		OrderId: info.OrderId,
		Source:  info.Source,
		Count:   info.Count,
	}).Error
}
