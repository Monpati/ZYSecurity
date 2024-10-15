package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxQueryLog struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	CacheCalls int64  `gorm:"column:cache_calls" json:"cache_calls"`
	CacheRate  int64  `gorm:"column:cache_rate" json:"cache_rate"`
	Domain     string `gorm:"column:domain" json:"domain"`
	TotalCalls int64  `gorm:"column:total_calls" json:"total_calls"`
}

func (DxQueryLog) TableName() string {
	return "DexunDomainQuery"
}

func AddQueryLog(db *gorm.DB, info *form.QueryLogsInfo) error {
	items := db.Table("DexunDomainQuery")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxQueryLog{
		Id:         sf.Generate(),
		CacheCalls: info.CacheCalls,
		CacheRate:  info.CacheRate,
		Domain:     info.Domain,
		TotalCalls: info.TotalCalls,
	}).Error
}
