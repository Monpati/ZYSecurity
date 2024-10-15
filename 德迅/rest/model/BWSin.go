package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type BWSingle struct {
	Id         int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	BwId       int64      `gorm:"column:bw_id" json:"bw_id"`
	OrderId    int64      `gorm:"column:order_id" json:"order_id"`
	DomainId   int64      `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string     `gorm:"column:domain_uuid" json:"domain_uuid"`
	Type       int64      `gorm:"column:type" json:"type"`
	IpList     utils.JSON `gorm:"column:ip_list" json:"ip_list"`
}

func (BWSingle) TableName() string {
	return "ScdnBWInstance"
}

func CreateBWSingle(db *gorm.DB, info *form.BWInfo) error {
	items := db.Table("ScdnBWInstance")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&BWSingle{
		Id:         sf.Generate(),
		BwId:       info.BwId,
		OrderId:    info.OrderId,
		DomainId:   info.DomainId,
		DomainUuid: info.DomainUuid,
		Type:       info.Type,
		IpList:     info.IpList,
	}).Error
}

func (p *BWSingle) GetBWSingleLists(db *gorm.DB, info *form.BWInfoFilterForm) (*[]BWInstance, int, error) {
	var lists []BWInstance
	var total int
	query := db.Model(&BWInstance{})

	if info.DomainUuid != "" {
		query = query.Where("`domain_uuid` LIKE ?", "%"+info.DomainUuid+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
		return nil, 0, err
	}

	return &lists, total, nil
}

func (p *BWSingle) UpdateBWSingle(db *gorm.DB, id int64, info *form.BWInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&BWSingle{
			DomainId:   info.DomainId,
			DomainUuid: info.DomainUuid,
			Type:       info.Type,
		}).Error
}
