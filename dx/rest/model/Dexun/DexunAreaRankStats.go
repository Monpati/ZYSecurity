package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxAreaRankStats struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	Domain     string `gorm:"column:domain" json:"domain"`
	TotalCount int64  `gorm:"column:total_count" json:"total_count"`
}

func (DxAreaRankStats) TableName() string {
	return "DexunAreaRankStats"
}

func AddAreaRankStats(db *gorm.DB, info *form.AreaRankStatsInfo) error {
	items := db.Table("DexunAreaRankStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxAreaRankStats{
		Id:         sf.Generate(),
		OrderId:    info.OrderId,
		Domain:     info.Domain,
		TotalCount: info.TotalCount,
	}).Error
}
