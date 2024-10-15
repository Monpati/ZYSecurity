package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type DxTotalFlow struct {
	Id                    int64 `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId               int64 `gorm:"column:order_id" json:"order_id"`
	RequestBandWidthPeak  int64 `gorm:"column:request_bandwidth_peak" json:"request_bandwidth_peak"`
	Requests              int64 `gorm:"column:requests" json:"requests"`
	ResponseBandWidthPeak int64 `gorm:"column:response_bandwidth_peak" json:"response_bandwidth_peak"`
	TotalRequestFlows     int64 `gorm:"column:total_request_flows" json:"total_request_flows"`
	TotalResponseFlows    int64 `gorm:"column:total_response_flows" json:"total_response_flows"`
	UnidentifiedAttack    int64 `gorm:"column:unidentified_attack" json:"unidentified_attack"`
	CreateTime            int64 `gorm:"column:create_time" json:"create_time"`
}

func (DxTotalFlow) TableName() string {
	return "DexunTotalFlow"
}

func AddTotalFlow(db *gorm.DB, info *form.TotalFlowInfo) error {
	items := db.Table("DexunTotalFlow")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxTotalFlow{
		Id:                    sf.Generate(),
		OrderId:               info.OrderId,
		RequestBandWidthPeak:  info.RequestBandWidthPeak,
		Requests:              info.Requests,
		ResponseBandWidthPeak: info.ResponseBandWidthPeak,
		TotalRequestFlows:     info.TotalRequestFlows,
		TotalResponseFlows:    info.TotalResponseFlows,
		UnidentifiedAttack:    info.UnidentifiedAttack,
		CreateTime:            int64(int(time.Now().Unix())),
	}).Error
}

func (p *DxTotalFlow) GetDataByOrderId(db *gorm.DB, order_id int64) error {
	if err := db.Table(p.TableName()).Where("order_id = ?", order_id).Order("create_time DESC").Limit(1).Error; err != nil {
		return err
	}
	return nil
}
