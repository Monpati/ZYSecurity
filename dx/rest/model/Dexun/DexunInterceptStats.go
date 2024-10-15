package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type DxInterceptStats struct {
	Id         int64 `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64 `gorm:"column:order_id" json:"order_id"`
	AppCC      int64 `gorm:"column:app_cc" json:"app_cc"`
	CC         int64 `gorm:"column:cc" json:"cc"`
	IpBlack    int64 `gorm:"column:ip_black" json:"ip_black"`
	Referer    int64 `gorm:"column:referer" json:"referer"`
	UrlBlack   int64 `gorm:"column:url_black" json:"url_black"`
	WebProtect int64 `gorm:"column:web_protect" json:"web_protect"`
	Other      int64 `gorm:"column:other" json:"other"`
	AreaAcc    int64 `gorm:"column:area_acc" json:"area_acc"`
	SafeAcc    int64 `gorm:"column:safe_acc" json:"safe_acc"`
	PreAcc     int64 `gorm:"column:pre_acc" json:"pre_acc"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
}

func (DxInterceptStats) TableName() string {
	return "DexunInterceptStats"
}

func AddInterceptStats(db *gorm.DB, info *form.InterceptStatsInfo) error {
	items := db.Table("DexunInterceptStats")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DxInterceptStats{
		Id:         sf.Generate(),
		OrderId:    info.OrderId,
		AppCC:      info.AppCC,
		CC:         info.CC,
		IpBlack:    info.IpBlack,
		Referer:    info.Referer,
		UrlBlack:   info.UrlBlack,
		WebProtect: info.WebProtect,
		Other:      info.Other,
		AreaAcc:    info.AreaAcc,
		SafeAcc:    info.SafeAcc,
		PreAcc:     info.PreAcc,
		CreateTime: int64(int(time.Now().Unix())),
	}).Error
}

func (p *DxInterceptStats) GetDataByOrderId(db *gorm.DB, order_id int64) error {
	if err := db.Table(p.TableName()).Where("order_id = ?", order_id).Order("create_time DESC").Limit(1).Find(&p).Error; err != nil {
		return err
	}
	return nil
}
