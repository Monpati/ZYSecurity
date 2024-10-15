package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type UrlBW struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	BwlId      int64  `gorm:"column:bwl_id" json:"bwl_id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	Type       string `gorm:"column:type" json:"type"`
	Path       string `gorm:"column:path" json:"path"`
	Method     string `gorm:"column:method" json:"method"`
	Active     int64  `gorm:"column:active" json:"active"`
	Uuid       string `gorm:"column:uuid" json:"uuid"`
}

func (UrlBW) TableName() string {
	return "ScdnURLBWLists"
}

func CreateUrlBW(db *gorm.DB, info *form.UrlBWCon) error {
	items := db.Table("ScdnURLBWLists")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&UrlBW{
		Id:         sf.Generate(),
		BwlId:      info.BwlId,
		OrderId:    info.OrderId,
		DomainId:   info.DomainId,
		DomainUuid: info.DomainUuid,
		Type:       info.Type,
		Path:       info.Path,
		Method:     info.Method,
		Active:     info.Active,
		Uuid:       info.Uuid,
	}).Error
}

func (p *UrlBW) GetUrlBWLists(db *gorm.DB, info *form.Filter) (*[]UrlBW, int, error) {
	var lists []UrlBW
	var total int

	if info.Field != "" {
		if info.Field == "type" {
			if err := db.Model(&UrlBW{}).Where("`type` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`type` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}

		if info.Field == "method" {
			if err := db.Model(&UrlBW{}).Where("`method` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`method` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}
	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&UrlBW{}).Count(&total).Error; err != nil {
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

func (p *UrlBW) UpdateUrlBW(db *gorm.DB, id int64, info *form.UrlBWCon) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&UrlBW{
			DomainId:   info.DomainId,
			DomainUuid: info.DomainUuid,
			Type:       info.Type,
			Path:       info.Path,
			Method:     info.Method,
			Active:     info.Active,
		}).Error
}
