package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type LeechLink struct {
	Id         int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	DlId       int64      `gorm:"column:dl_id" json:"dl_id"`
	DomainId   int64      `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string     `gorm:"column:domain_uuid" json:"domain_uuid"`
	OrderId    int64      `gorm:"column:order_id" json:"order_id"`
	DdUuid     string     `gorm:"column:dd_uuid" json:"dd_uuid"`
	ProType    string     `gorm:"column:pro_type" json:"pro_type"`
	Type       string     `gorm:"column:type" json:"type"`
	Active     string     `gorm:"column:active" json:"active"`
	Domains    utils.JSON `gorm:"column:domains" json:"domains"`
	AllowEmpty string     `gorm:"column:allow_empty" json:"allow_empty"`
}

func (LeechLink) TableName() string {
	return "ScdnLeechlink"
}

func CreateLeechLink(db *gorm.DB, info *form.LeechLinkInfo) error {
	items := db.Table("ScdnLeechlink")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&LeechLink{
		Id:         sf.Generate(),
		DomainId:   info.DomainId,
		DomainUuid: info.DomainUuid,
		OrderId:    info.OrderId,
		DdUuid:     info.DdUuid,
		ProType:    info.ProType,
		Type:       info.Type,
		Active:     info.Active,
		Domains:    info.Domains,
		AllowEmpty: info.AllowEmpty,
	}).Error
}

func (p *LeechLink) GetLeechLinkLists(db *gorm.DB, info *form.LeechLinkFilterForm) (*[]LeechLink, int, error) {
	var lists []LeechLink
	var total int
	query := db.Model(&LeechLink{})

	if info.ProType != "" {
		query = query.Where("`pro_type` LIKE ?", "%"+info.ProType+"%")
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
		return nil, 0, err
	}
	return &lists, total, nil
}

func (p *LeechLink) UpdateLeechLink(db *gorm.DB, id int64, info *form.LeechLinkInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&LeechLink{
			DdUuid:     info.DdUuid,
			ProType:    info.ProType,
			Type:       info.Type,
			Active:     info.Active,
			Domains:    info.Domains,
			AllowEmpty: info.AllowEmpty,
		}).Error
}
