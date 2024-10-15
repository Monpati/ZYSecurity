package model

import (
	"Dexun/form"
	"Dexun/model/Dexun"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type AreaAcc struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	OrderUuid  string `gorm:"column:order_uuid" json:"order_uuid"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	DaId       int64  `gorm:"column:da_id" json:"da_id"`
	Regions    JSON   `gorm:"column:regions" json:"regions"`
	Active     string `gorm:"column:active" json:"active"`
}

type DexunAreaAcc struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	Regions    JSON   `gorm:"column:regions" json:"regions"`
	Active     string `gorm:"column:active" json:"active"`
}

func (AreaAcc) TableName() string {
	return "ScdnAreaAccessCon"
}

func CreateDexunAreaAcc(db *gorm.DB, info *Dexun.DeXunBody) error {
	items := db.Table("DexunAreaAccessCon")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&AreaAcc{
		Id:      sf.Generate(),
		Active:  info.Active,
		Regions: JSON(info.Config.Regions),
	}).Error
}

func CreateAreaAcc(db *gorm.DB, info *form.AreaAccInfo) error {
	items := db.Table("ScdnAreaAccessCon")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&AreaAcc{
		Id:         sf.Generate(),
		DomainUuid: info.DomainUuid,
		DomainId:   info.DomainId,
		OrderUuid:  info.OrderUuid,
		OrderId:    info.OrderId,
		DaId:       info.DaId,
		Active:     info.Active,
		Regions:    JSON(info.Regions),
	}).Error
}

func (p *AreaAcc) GetAreaAccLists(db *gorm.DB, info *form.AreaAccFilterForm) (*[]AreaAcc, int, error) {
	var lists []AreaAcc
	var total int

	query := db.Model(&Account{})

	if info.Active != "" {
		query = query.Where("`active` LIKE ?", "%"+info.Active+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
		return nil, 0, err
	}

	return &lists, total, nil
}

func (p *AreaAcc) UpdateAreaAcc(db *gorm.DB, id int64, info *form.AreaAccInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&AreaAcc{
			DaId:    info.DaId,
			Active:  info.Active,
			Regions: JSON(info.Regions),
		}).Error
}
