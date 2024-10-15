package model

import (
	"Dexun/form"
	"Dexun/model/Dexun"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type ScdnService struct {
	Id              int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	ComboId         int64  `gorm:"column:combo_id" json:"combo_id"`
	Uuid            string `gorm:"column:uuid" json:"uuid"`
	TcName          string `gorm:"column:tc_name" json:"tc_name"`
	KsMoney         int64  `gorm:"column:ks_money" json:"ks_money"`
	ProFlow         int64  `gorm:"column:pro_flow" json:"pro_flow"`
	DdosHh          string `gorm:"column:ddos_hh" json:"ddos_hh"`
	DomainNum       int64  `gorm:"column:domain_num" json:"domain_num"`
	ZkMoney         int64  `gorm:"column:zk_money" json:"zk_money"`
	Source          string `gorm:"column:source" json:"source"`
	Status          int    `gorm:"column:status" json:"status"`
	ZyzkMoney       int64  `gorm:"column:zyzk_money" json:"zyzk_money"`
	ZyksMoney       int64  `gorm:"column:zyks_money" json:"zyks_money"`
	Months          int64  `gorm:"column:months" json:"months"`
	ActuaFlow       int64  `gorm:"column:actua_flow" json:"actua_flow"`
	EndTime         string `gorm:"column:end_time" json:"end_time"`
	KsStart         int64  `gorm:"column:ks_start" json:"ks_start"`
	ProductSitename string `gorm:"column:product_sitename" json:"product_sitename"`
	RechargeFlow    int64  `gorm:"column:recharge_flow" json:"recharge_flow"`
	RechargeDomain  int64  `gorm:"column:recharge_domain" json:"recharge_domain"`
	ServerIp        string `gorm:"column:server_ip" json:"server_ip"`
	StartTime       string `gorm:"column:stat_time" json:"stat_time"`
	TotalFlow       int64  `gorm:"column:total_flow" json:"total_flow"`
	UUserId         int64  `gorm:"column:u_user_id" json:"u_user_id"`
	SiteStart       int64  `gorm:"column:site_stat" json:"site_stat"`
	UserId          int64  `gorm:"column:user_id" json:"user_id"`
	UserName        string `gorm:"column:username" json:"username"`
	Agent           string `gorm:"column:agent" json:"agent"`
	PackageId       int64  `gorm:"column:package_id" json:"package_id"`
	ProType         int64  `gorm:"column:pro_type" json:"pro_type"`
}

func (ScdnService) TableName() string {
	return "ScdnService"
}

func (p *ScdnService) GetByParams(db *gorm.DB, info *form.Filter) (*[]ScdnService, int, error) {
	var services []ScdnService
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&ScdnService{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

		if info.Field == "source" {
			if err := db.Model(&ScdnService{}).Where("`source` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`source` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

		if info.Field == "u_user_id" {
			if err := db.Model(&ScdnService{}).Where("`u_user_id` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
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
		if err := db.Model(&ScdnService{}).Count(&total).Error; err != nil {
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

func (p *ScdnService) GetUuidById(db *gorm.DB, id int64) string {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error; err != nil {
		return ""
	} else {
		return p.Uuid
	}
}

func (p *ScdnService) GetProTypeById(db *gorm.DB, id int64) int64 {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error; err != nil {
		return 0
	}
	return p.ProType
}

func (p *ScdnService) GetIdByUuid(db *gorm.DB, uuid string) int64 {
	if err := db.Table(p.TableName()).Where("uuid = ?", uuid).Find(&p).Error; err != nil {
		return 0
	} else {
		return p.Id
	}
}

func (p *ScdnService) GetIdByUserId(db *gorm.DB, user_id int64) *[]ScdnService {
	var orders []ScdnService
	if err := db.Table(p.TableName()).Where("user_id = ?", user_id).Find(&orders).Error; err != nil {
		return nil
	}
	return &orders
}

func (p *ScdnService) GetOrdersByUser(db *gorm.DB, info *form.Filter) (*[]ScdnService, int, error) {
	var orders []ScdnService
	var total int

	if err := db.Model(&ScdnService{}).
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

func (p *ScdnService) GetOrdersByAgent(db *gorm.DB, info *form.Filter) (*[]ScdnService, int, error) {
	var orders []ScdnService
	var total int

	if err := db.Table(p.TableName()).Where("agent = ?", p.Agent).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).Limit(info.Limit).Where("agent = ?", p.Agent).Offset(info.Offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return &orders, total, nil
}

func CreateScdnService(db *gorm.DB, datum Dexun.Datum, combo_id, months, user_id, pro_type int64, username, source string) error {
	items := db.Table("ScdnService")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&ScdnService{
		Id:              sf.Generate(),
		ComboId:         combo_id,
		Uuid:            datum.UUID,
		TcName:          datum.TcName,
		KsMoney:         datum.KsMoney,
		ZyksMoney:       3 * datum.KsMoney,
		ProFlow:         datum.ProFlow,
		DdosHh:          datum.DdosHh,
		DomainNum:       datum.DomainNum,
		Months:          months,
		Source:          source,
		ActuaFlow:       datum.ActuaFlow,
		EndTime:         datum.EndTime,
		KsStart:         datum.KsStart,
		ProductSitename: datum.ProductSitename,
		RechargeFlow:    datum.RechargeFlow,
		ServerIp:        datum.ServerIP,
		StartTime:       datum.StatTime,
		TotalFlow:       datum.TotalFlow,
		UUserId:         datum.UUserID,
		SiteStart:       datum.SiteStart,
		UserId:          user_id,
		UserName:        username,
		PackageId:       0,
		ProType:         pro_type,
	}).Error
}

func (p *ScdnService) GetAllRecord(db *gorm.DB) *[]ScdnService {
	var orders []ScdnService

	if err := db.Table("ScdnService").Find(&orders).Error; err != nil {
		return nil
	}
	return &orders
}

func (p *Account) UpdateBalancePurchase(db *gorm.DB, id, balance, money int64) error {
	return db.Table(p.TableName()).Where("id = ?", id).UpdateColumn("balance", balance-money).Error
}

func (p *Account) UpdateBalanceRecharge(db *gorm.DB, id, balance, money int64) error {
	return db.Table(p.TableName()).Where("id = ?", id).UpdateColumn("balance", balance+money).Error
}
