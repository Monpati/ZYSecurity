package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type SafeAcc struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	OrderUuid  string `gorm:"column:order_uuid" json:"order_uuid"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	DsId       int64  `gorm:"column:ds_id" json:"ds_id"`
	Password   string `gorm:"column:password" json:"password"`
	URL        string `gorm:"column:url" json:"url"`
}

func (SafeAcc) TableName() string {
	return "ScdnSafeAccessCon"
}

func CreateSafeAcc(db *gorm.DB, info *form.SafeAccCon) error {
	items := db.Table("ScdnSafeAccessCon")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&SafeAcc{
		Id:         sf.Generate(),
		DomainUuid: info.DomainUuid,
		DomainId:   info.DomainId,
		OrderUuid:  info.OrderUuid,
		OrderId:    info.OrderId,
		DsId:       info.DsId,
		Password:   info.Password,
		URL:        info.URL,
	}).Error
}

func (p *SafeAcc) GetSafeAccLists(db *gorm.DB, info *form.Filter) (*[]SafeAcc, int, error) {
	var lists []SafeAcc
	var total int

	if info.Field != "" {
		if info.Field == "domain_uuid" {
			if err := db.Model(&SafeAcc{}).Where("`domain_uuid` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_uuid` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}

		if info.Field == "order_uuid" {
			if err := db.Model(&SafeAcc{}).Where("`order_uuid` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
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
		if err := db.Model(&SafeAcc{}).Count(&total).Error; err != nil {
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

func (p *SafeAcc) UpdateSafeAcc(db *gorm.DB, id int64, info *form.SafeAccInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&SafeAcc{
			DsId: info.DsId,
		}).Error
}
