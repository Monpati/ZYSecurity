package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxHttpPackStats struct {
	Id         int64 `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64 `gorm:"column:order_id" json:"order_id"`
	Time       int64 `gorm:"column:time" json:"time"`
	TotalCount int64 `gorm:"column:total_count" json:"total_count"`
}

func (DxHttpPackStats) TableName() string {
	return "DexunHttpPackStats"
}

func AddHttpPackStats(db *gorm.DB, info *form.HttpPackStatsInfo) error {
	items := db.Table("DexunHttpPackStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxHttpPackStats{
		Id:         sf.Generate(),
		OrderId:    info.OrderId,
		Time:       info.Time,
		TotalCount: info.TotalCount,
	}).Error
}
