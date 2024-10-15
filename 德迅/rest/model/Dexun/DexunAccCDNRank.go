package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type AccCDNRank struct {
	Id       int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId  int64  `gorm:"column:order_id" json:"order_id"`
	ClientIp string `gorm:"column:client_ip" json:"client_ip"`
	CountSum int64  `gorm:"column:count_sum" json:"count_sum"`
}

func (AccCDNRank) TableName() string {
	return "DexunAccCDNRank"
}

func AddAccCDNRank(db *gorm.DB, info *form.AccCDNRankInfo) error {
	items := db.Table("DexunAccCDNRank")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&AccCDNRank{
		Id:       sf.Generate(),
		OrderId:  info.OrderId,
		ClientIp: info.ClientIp,
		CountSum: info.CountSum,
	}).Error
}
