package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type PicRC struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	OrderUuid  string `gorm:"column:order_uuid" json:"order_uuid"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	DpId       int64  `gorm:"column:dp_id" json:"dp_id"`
	Active     int64  `gorm:"column:active" json:"active"`
}

func (PicRC) TableName() string {
	return "ScdnPicRC"
}

func CreatePicRC(db *gorm.DB, info *form.PicRCInfo) error {
	items := db.Table("ScdnPicRC")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&PicRC{
		Id:         sf.Generate(),
		DomainUuid: info.DomainUuid,
		DomainId:   info.DomainId,
		OrderUuid:  info.OrderUuid,
		OrderId:    info.OrderId,
		DpId:       info.DpId,
		Active:     info.Active,
	}).Error
}

func (p *PicRC) GetPicRCLists(db *gorm.DB, info *form.Filter) (*[]PicRC, int, error) {
	var lists []PicRC
	var total int

	if info.Field != "" {
		if info.Field == "domain_uuid" {
			if err := db.Model(&PicRC{}).Where("`domain_uuid` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_uuid` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}

		if info.Field == "order_uuid" {
			if err := db.Model(&PicRC{}).Where("`order_uuid` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
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
		if err := db.Model(&PicRC{}).Count(&total).Error; err != nil {
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

func (p *PicRC) UpdatePicRC(db *gorm.DB, id int64, info *form.PicRCInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&PicRC{
			Active: info.Active,
		}).Error
}
