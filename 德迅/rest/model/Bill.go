package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	Id        int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	UserId    int64  `gorm:"column:user_id" json:"user_id"`
	Amount    int64  `gorm:"column:amount" json:"amount"`
	ServiceId *int64 `gorm:"column:service_id" json:"service_id"`
	Create    int    `gorm:"column:create" json:"create"`
	Payee     string `gorm:"column:payee" json:"payee"`
	Payer     string `gorm:"column:payer" json:"payer"`
	Explain   string `gorm:"column:explain" json:"explain"`
	Status    int    `gorm:"column:status" json:"status"`
	Agent     string `gorm:"column:agent" json:"agent"`
	DDoSId    *int64 `gorm:"column:ddos_id" json:"ddos_id"`
	Recharge  int64  `gorm:"column:recharge" json:"recharge"`
	Method    string `gorm:"column:method" json:"method"`
}

func (Bill) TableName() string {
	return "Bill_Details"
}

func CreateBill(db *gorm.DB, info form.Purchase) error {
	items := db.Table("Bill_Details")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&Bill{
		Id:        sf.Generate(),
		ServiceId: info.ServiceId,
		DDoSId:    info.DDoSId,
		UserId:    info.UserId,
		Payer:     info.Username,
		Agent:     info.Agent,
		Create:    int(time.Now().Unix()),
		Recharge:  info.Recharge,
		Method:    info.Method,
	}).Error
}

func (p *Bill) GetBills(db *gorm.DB, info *form.BillsFilterForm) (*[]Bill, int, error) {
	var bills []Bill
	var total int

	if err := db.Table(p.TableName()).
		Where("user_id", p.UserId).
		Find(&bills).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).
		Limit(info.Limit).
		Where("user_id", p.UserId).
		Offset(info.Offset).
		Find(&bills).Error; err != nil {
		return nil, 0, err
	}
	return &bills, total, nil
}

func (p *Bill) GetByParams(db *gorm.DB, info *form.BillsFilterForm) (*[]Bill, int, error) {
	var bills []Bill
	var total int

	query := db.Model(&Account{})

	if info.Username != "" {
		query = query.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	if info.Months != "" {
		query = query.Where("`months` LIKE ?", "%"+info.Months+"%")
	}
	if info.Agent != "" {
		query = query.Where("`agent` LIKE ?", "%"+info.Agent+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&bills).Error; err != nil {
		return nil, 0, err
	}

	return &bills, total, nil
}
