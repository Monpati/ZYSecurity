package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DDoSService struct {
	Id              int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	ComboId         int64  `gorm:"column:combo_id" json:"combo_id"`
	Uuid            string `gorm:"column:uuid" json:"uuid"`
	UUserId         int64  `gorm:"column:u_user_id" json:"u_user_id"`
	ServerIp        string `gorm:"column:server_ip" json:"server_ip"`
	TcName          string `gorm:"column:tc_name" json:"tc_name"`
	KsMoney         int64  `gorm:"column:ks_money" json:"ks_money"`
	ProductSitename string `gorm:"column:product_sitename" json:"product_sitename"`
	StatTime        string `gorm:"column:stat_time" json:"stat_time"`
	EndTime         string `gorm:"column:end_time" json:"end_time"`
	SiteStart       int64  `gorm:"column:site_start" json:"site_start"`
	DdosHh          string `gorm:"column:ddos_hh" json:"ddos_hh"`
	KsStart         int64  `gorm:"column:ks_start" json:"ks_start"`
	DomainNum       int64  `gorm:"column:domain_num" json:"domain_num"`
	RechargeDomain  int64  `gorm:"column:recharge_domain" json:"recharge_domain"`
	PortNum         int64  `gorm:"column:port_num" json:"port_num"`
	RechargePort    int64  `gorm:"column:recharge_port" json:"recharge_port"`
	UserId          int64  `gorm:"column:user_id" json:"user_id"`
	Agent           string `gorm:"column:agent" json:"agent"`
	ProType         int64  `gorm:"column:pro_type" json:"pro_type"`
}

func (DDoSService) TableName() string {
	return "DDoSService"
}

func CreateDDoSService(db *gorm.DB, info *form.DDoSServiceInfo) error {
	items := db.Table("DDoSService")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DDoSService{
		Id:              sf.Generate(),
		ComboId:         info.ComboId,
		Uuid:            info.Uuid,
		UUserId:         info.UUserId,
		ServerIp:        info.ServerIp,
		TcName:          info.TcName,
		KsMoney:         info.KsMoney,
		ProductSitename: info.ProductSitename,
		StatTime:        info.StatTime,
		EndTime:         info.EndTime,
		SiteStart:       info.SiteStart,
		DdosHh:          info.DdosHh,
		KsStart:         info.KsStart,
		DomainNum:       info.DomainNum,
		RechargeDomain:  info.RechargeDomain,
		PortNum:         info.PortNum,
		RechargePort:    info.RechargePort,
		UserId:          info.UserId,
		Agent:           info.Agent,
		ProType:         info.ProType,
	}).Error
}

func (p *DDoSService) GetByParams(db *gorm.DB, info *form.Filter) (*[]DDoSService, int, error) {
	var services []DDoSService
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&DDoSService{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

		if info.Field == "source" {
			if err := db.Model(&DDoSService{}).Where("`source` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`source` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

		if info.Field == "u_user_id" {
			if err := db.Model(&DDoSService{}).Where("`u_user_id` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`u_user_id` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&DDoSService{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Offset(info.Offset).Find(&services).Error; err != nil {
			return nil, 0, err
		}
		return &services, total, nil
	}
	return &services, total, nil
}

func (p *DDoSService) GetOrdersByUser(db *gorm.DB, info *form.Filter) (*[]DDoSService, int, error) {
	var orders []DDoSService
	var total int

	if err := db.Model(&DDoSService{}).
		Where("user_id = ?", p.UserId).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).
		Limit(info.Limit).
		Where("user_id = ?", p.UserId).
		Offset(info.Offset).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return &orders, total, nil

}

func (p *DDoSService) GetOrdersByAgent(db *gorm.DB, info *form.Filter) (*[]DDoSService, int, error) {
	var orders []DDoSService
	var total int

	if err := db.Table(p.TableName()).Where("agent = ?", p.Agent).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).Limit(info.Limit).Where("agent = ?", p.Agent).Offset(info.Offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return &orders, total, nil
}

func (p *DDoSService) GetIdByUuid(db *gorm.DB, uuid string) int64 {
	db.Table(p.TableName()).Where("uuid = ?", uuid).Find(&p)
	return p.Id
}

func (p *DDoSService) GetUuidById(db *gorm.DB, id int64) string {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error; err != nil {
		return ""
	} else {
		return p.Uuid
	}
}
