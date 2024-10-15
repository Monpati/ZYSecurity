package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type PreAcc struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	OrderUuid  string `gorm:"column:order_uuid" json:"order_uuid"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	DaId       int64  `gorm:"column:da_id" json:"da_id"`
	Action     string `gorm:"column:action" json:"action"`
	Active     int64  `gorm:"column:active" json:"active"`
	CheckList  string `gorm:"column:check_list" json:"check_list"`
	MItem      string `gorm:"column:m_item" json:"m_item"`
	MValue     string `gorm:"column:m_value" json:"m_value"`
	MOperate   string `gorm:"column:m_operate" json:"m_operate"`
	MValueXs   string `gorm:"column:m_value_xs" json:"m_value_xs"`
}

func (PreAcc) TableName() string {
	return "ScdnPreAccessCon"
}

func CreatePreAcc(db *gorm.DB, info *form.PreAccConInfo) error {
	items := db.Table("ScdnPreAccessCon")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&PreAcc{
		Id:         sf.Generate(),
		DomainUuid: info.DomainUuid,
		DomainId:   info.DomainId,
		OrderUuid:  info.OrderUuid,
		OrderId:    info.OrderId,
		DaId:       info.DaId,
		Action:     info.Action,
		Active:     info.Active,
		CheckList:  info.CheckList,
		MItem:      info.MItem,
		MValue:     info.MValue,
		MOperate:   info.MOperate,
		MValueXs:   info.MValueXs,
	}).Error
}

func (p *PreAcc) GetPreAccLists(db *gorm.DB, info *form.Filter) (*[]PreAcc, int, error) {
	var lists []PreAcc
	var total int

	if info.Field != "" {
		if info.Field == "policy" {
			if err := db.Model(&PreAcc{}).Where("`policy` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`policy` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}

		if info.Field == "concurrency" {
			if err := db.Model(&PreAcc{}).Where("`concurrency` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`concurrency` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&lists).Error; err != nil {
				return nil, 0, err
			}
			return &lists, total, nil
		}
	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&PreAcc{}).Count(&total).Error; err != nil {
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

func (p *PreAcc) UpdatePreAcc(db *gorm.DB, id int64, info *form.PreAccInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&PreAcc{
			Active: info.Active,
		}).Error
}
