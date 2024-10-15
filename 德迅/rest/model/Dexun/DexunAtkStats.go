package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxAtkStats struct {
	Id         int64 `gorm:"primary_key" gorm:"column:id" json:"id"`
	Time       int64 `gorm:"column:time" json:"time"`
	OrderId    int64 `gorm:"column:order_id" json:"order_id"`
	TotalCount int64 `gorm:"column:total_count" json:"total_count"`
}

func (DxAtkStats) TableName() string {
	return "DexunAtkStats"
}

func AddAtkStats(db *gorm.DB, info *form.AtkStatsInfo) error {
	items := db.Table("DexunAtkStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxAtkStats{
		Id:         sf.Generate(),
		Time:       info.Time,
		OrderId:    info.OrderId,
		TotalCount: info.TotalCount,
	}).Error
}
