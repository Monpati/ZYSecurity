package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type WordRC struct {
	Id         int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64      `gorm:"column:order_id" json:"order_id"`
	OrderUuid  string     `gorm:"column:order_uuid" json:"order_uuid"`
	DomainId   int64      `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string     `gorm:"column:domain_uuid" json:"domain_uuid"`
	DwId       int64      `gorm:"column:dw_id" json:"dw_id"`
	Gzip       string     `gorm:"column:gzip" json:"gzip"`
	Active     string     `gorm:"column:active" json:"active"`
	KeyWords   utils.JSON `gorm:"column:keywords" json:"keywords"`
}

func (WordRC) TableName() string {
	return "ScdnWordRC"
}

func CreateWordRC(db *gorm.DB, info *form.WordRCInfo) error {
	items := db.Table("ScdnWordRC")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&WordRC{
		Id:         sf.Generate(),
		DomainUuid: info.DomainUuid,
		DomainId:   info.DomainId,
		OrderUuid:  info.OrderUuid,
		OrderId:    info.OrderId,
		DwId:       info.DwId,
		Gzip:       info.Gzip,
		Active:     info.Active,
		KeyWords:   info.KeyWords,
	}).Error
}

func (p *WordRC) GetWordRCLists(db *gorm.DB, info *form.Filter) (*[]WordRC, int, error) {
	var lists []WordRC
	var total int

	if info.Field != "" {
		if info.Field == "domain_uuid" {
			if err := db.Model(&WordRC{}).Where("`domain_uuid` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_uuid` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}

		if info.Field == "order_uuid" {
			if err := db.Model(&WordRC{}).Where("`order_uuid` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`order_uuid` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}
	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&WordRC{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
			return nil, 0, err
		}
		return &lists, total, nil
	}
	return &lists, total, nil
}

func (p *WordRC) UpdateWordRC(db *gorm.DB, id int64, info *form.WordRCInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&WordRC{
			DwId:     info.DwId,
			Gzip:     info.Gzip,
			Active:   info.Active,
			KeyWords: info.KeyWords,
		}).Error
}
