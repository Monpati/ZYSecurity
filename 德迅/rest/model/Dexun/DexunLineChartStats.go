package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DxLineChartStats struct {
	Id           int64 `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId      int64 `gorm:"column:order_id" json:"order_id"`
	Time         int64 `gorm:"column:time" json:"time"`
	ResponseSize int64 `gorm:"column:response_size" json:"response_size"`
	RequestSize  int64 `gorm:"column:request_size" json:"request_size"`
}

func (DxLineChartStats) TableName() string {
	return "DexunLineChartStats"
}

func AddLineChartStats(db *gorm.DB, info *form.LineChartStatsInfo) error {
	items := db.Table("DexunLineChartStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxLineChartStats{
		Id:           sf.Generate(),
		OrderId:      info.OrderId,
		Time:         info.Time,
		ResponseSize: info.ResponseSize,
		RequestSize:  info.RequestSize,
	}).Error
}

func (p *DxLineChartStats) GetDataByOrderId(db *gorm.DB, order_id int64) *[]DxLineChartStats {
	var datas []DxLineChartStats
	if err := db.Table(p.TableName()).Where("order_id = ?", order_id).Find(&datas).Error; err != nil {
		return nil
	}
	return &datas
}
