package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type ScdnHeat struct {
	Id       int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	DomainId int64  `gorm:"column:domain_id" json:"domain_id"`
	Url      string `gorm:"column:url" json:"url"`
	Status   int    `gorm:"column:status" json:"status"`
}

func (ScdnHeat) TableName() string {
	return "ScdnHeat"
}

func CreateHeat(db *gorm.DB, info *form.ScdnHeatInfo) error {
	items := db.Table("ScdnHeat")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&ScdnHeat{
		Id:       sf.Generate(),
		DomainId: info.DomainId,
		Url:      info.Url,
	}).Error
}

func (p *ScdnHeat) UpdateHeatStatus(db *gorm.DB, info *form.ScdnHeatInfo) error {
	return db.Table("ScdnHeat").Where("domain_id = ?", info.DomainId).UpdateColumn("status", info.Active).Error
}

func (p *ScdnHeat) GetHeatList(db *gorm.DB, info *form.HeatFilterForm) (*[]ScdnHeat, error) {
	var domains []ScdnHeat
	var total int
	query := db.Model(&ScdnHeat{})

	query = query.Where("domain_id = ?", info.DomainId)

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&domains).Error; err != nil {
		return nil, err
	}
	return &domains, nil
}
