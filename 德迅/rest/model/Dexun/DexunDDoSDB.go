package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DexunDDoSCombos struct {
	ID            int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	Uuid          string `gorm:"column:uuid" json:"uuid"`
	TcName        string `gorm:"column:tc_name" json:"tc_name"`
	KsMoney       int64  `gorm:"column:ks_money" json:"ks_money"`
	ZkMoney       int64  `gorm:"column:zk_money" json:"zk_money"`
	Yms           string `gorm:"column:yms" json:"yms"`
	Ccfy          string `gorm:"column:ccfy" json:"ccfy"`
	Gjwaf         string `gorm:"column:gjwaf" json:"gjwaf"`
	Ywll          string `gorm:"column:ywll" json:"ywll"`
	Gjcs          string `gorm:"column:gjcs" json:"gjcs"`
	XL            string `gorm:"column:xl" json:"xl"`
	Zfdks         string `gorm:"column:zfdks" json:"zfdks"`
	Fhyms         string `gorm:"column:fhyms" json:"fhyms"`
	Ywdk          string `gorm:"column:ywdk" json:"ywdk"`
	ProFlow       int64  `gorm:"column:pro_flow" json:"pro_flow"`
	DdosHh        string `gorm:"column:ddos_hh" json:"ddos_hh"`
	DomainNum     int64  `gorm:"column:domain_num" json:"domain_num"`
	CompleteState int64  `gorm:"column:complete_state" json:"complete_state"`
	FirewallState int64  `gorm:"firewall_state" json:"firewall_state"`
	WafState      int64  `gorm:"column:waf_state" json:"waf_state"`
	GroupType     int64  `gorm:"column:group_type" json:"group_type"`
	PortNum       int64  `gorm:"column:port_num" json:"port_num"`
	YwdkNum       int64  `gorm:"column:ywdk_num" json:"ywdk_num"`
}

type DDoSCombos struct {
	ID            int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	ComboId       int64  `gorm:"column:combo_id" json:"combo_id"`
	Uuid          string `gorm:"column:uuid" json:"uuid"`
	TcName        string `gorm:"column:tc_name" json:"tc_name"`
	KsMoney       int64  `gorm:"column:ks_money" json:"ks_money"`
	ZkMoney       int64  `gorm:"column:zk_money" json:"zk_money"`
	Yms           string `gorm:"column:yms" json:"yms"`
	Ccfy          string `gorm:"column:ccfy" json:"ccfy"`
	Gjwaf         string `gorm:"column:gjwaf" json:"gjwaf"`
	Ywll          string `gorm:"column:ywll" json:"ywll"`
	Gjcs          string `gorm:"column:gjcs" json:"gjcs"`
	XL            string `gorm:"column:xl" json:"xl"`
	Zfdks         string `gorm:"column:zfdks" json:"zfdks"`
	Fhyms         string `gorm:"column:fhyms" json:"fhyms"`
	Ywdk          string `gorm:"column:ywdk" json:"ywdk"`
	ProFlow       int64  `gorm:"column:pro_flow" json:"pro_flow"`
	DdosHh        string `gorm:"column:ddos_hh" json:"ddos_hh"`
	DomainNum     int64  `gorm:"column:domain_num" json:"domain_num"`
	CompleteState int64  `gorm:"column:complete_state" json:"complete_state"`
	FirewallState int64  `gorm:"firewall_state" json:"firewall_state"`
	WafState      int64  `gorm:"column:waf_state" json:"waf_state"`
	GroupType     int64  `gorm:"column:group_type" json:"group_type"`
	ZyksMoney     int64  `gorm:"column:zyks_money" json:"zyks_money"`
	ZyzkMoney     int64  `gorm:"column:zyzk_money" json:"zyzk_money"`
	SellStatus    int    `gorm:"column:sell_status" json:"sell_status"`
	Source        string `gorm:"column:source" json:"source"`
	PortNum       int64  `gorm:"column:port_num" json:"port_num"`
	YwdkNum       int64  `gorm:"column:ywdk_num" json:"ywdk_num"`
}

func (DexunDDoSCombos) TableName() string {
	return "DexunDDoSCombos"
}

func CreateDxDDoSCombos(db *gorm.DB, info Data) error {
	items := db.Table("DexunDDoSCombos")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DexunDDoSCombos{
		ID:            sf.Generate(),
		Uuid:          info.UUID,
		TcName:        info.TcName,
		KsMoney:       info.KsMoney,
		ZkMoney:       info.ZkMoney,
		Yms:           info.ProNote.Yms,
		Ccfy:          info.ProNote.Ccfy,
		Gjwaf:         info.ProNote.Gjwaf,
		Ywll:          info.ProNote.Ywll,
		Gjcs:          info.ProNote.Gjcs,
		XL:            info.ProNote.XL,
		Zfdks:         info.ProNote.Zfdks,
		Fhyms:         info.ProNote.Fhyms,
		Ywdk:          info.ProNote.Ywdk,
		ProFlow:       info.ProFlow,
		DdosHh:        info.DdosHh,
		DomainNum:     info.DomainNum,
		CompleteState: info.CompleteState,
		FirewallState: info.FirewallState,
		WafState:      info.WafState,
		GroupType:     info.GroupType,
		PortNum:       info.PortNum,
		YwdkNum:       info.YwdkNum,
	}).Error
}

func CopyDxDDoSCombos(db *gorm.DB) error {
	items := db.Table("DDoSCombos")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	var combos []DexunDDoSCombos
	var service DDoSCombos
	if err := db.Table("DexunDDoSCombos").Find(&combos).Error; err != nil {
		return err
	}

	for _, combo := range combos {
		db.Table("DDoSCombos").Where("combo_id = ?", combo.ID).Find(&service)
		if service.ID == 0 {
			items.Create(&DDoSCombos{
				ID:            sf.Generate(),
				ComboId:       combo.ID,
				Uuid:          combo.Uuid,
				TcName:        combo.TcName,
				KsMoney:       combo.KsMoney,
				ZkMoney:       combo.ZkMoney,
				Yms:           combo.Yms,
				Ccfy:          combo.Ccfy,
				Gjwaf:         combo.Gjwaf,
				Ywll:          combo.Ywll,
				Gjcs:          combo.Gjcs,
				XL:            combo.XL,
				Zfdks:         combo.Zfdks,
				Fhyms:         combo.Fhyms,
				Ywdk:          combo.Ywdk,
				ProFlow:       combo.ProFlow,
				DdosHh:        combo.DdosHh,
				DomainNum:     combo.DomainNum,
				CompleteState: combo.CompleteState,
				FirewallState: combo.FirewallState,
				WafState:      combo.WafState,
				GroupType:     combo.GroupType,
				ZyksMoney:     3 * combo.KsMoney,
				ZyzkMoney:     3 * combo.ZkMoney,
				SellStatus:    1,
				Source:        "dexun",
				PortNum:       combo.PortNum,
				YwdkNum:       combo.YwdkNum,
			})
		} else {
			items.Where("combo_id = ?", combo.ID).Updates(&DDoSCombos{
				TcName:        combo.TcName,
				KsMoney:       combo.KsMoney,
				ZkMoney:       combo.ZkMoney,
				Yms:           combo.Yms,
				Ccfy:          combo.Ccfy,
				Gjwaf:         combo.Gjwaf,
				Ywll:          combo.Ywll,
				Gjcs:          combo.Gjcs,
				XL:            combo.XL,
				Zfdks:         combo.Zfdks,
				Fhyms:         combo.Fhyms,
				Ywdk:          combo.Ywdk,
				ProFlow:       combo.ProFlow,
				DdosHh:        combo.DdosHh,
				DomainNum:     combo.DomainNum,
				CompleteState: combo.CompleteState,
				FirewallState: combo.FirewallState,
				WafState:      combo.WafState,
				GroupType:     combo.GroupType,
				ZyksMoney:     3 * combo.KsMoney,
				ZyzkMoney:     3 * combo.ZkMoney,
				Source:        "dexun",
				PortNum:       combo.PortNum,
				YwdkNum:       combo.YwdkNum,
			})
		}
	}
	return nil
}

func (p *DexunDDoSCombos) UpdateComboByUuid(db *gorm.DB, uuid string, info Data) error {
	return db.Table(p.TableName()).
		Where("uuid = ?", uuid).
		Updates(&DexunDDoSCombos{
			TcName:        info.TcName,
			KsMoney:       info.KsMoney,
			ZkMoney:       info.ZkMoney,
			Yms:           info.ProNote.Yms,
			Ccfy:          info.ProNote.Ccfy,
			Gjwaf:         info.ProNote.Gjwaf,
			Ywll:          info.ProNote.Ywll,
			XL:            info.ProNote.XL,
			Zfdks:         info.ProNote.Zfdks,
			Fhyms:         info.ProNote.Fhyms,
			Ywdk:          info.ProNote.Ywdk,
			ProFlow:       info.ProFlow,
			DdosHh:        info.DdosHh,
			DomainNum:     info.DomainNum,
			CompleteState: info.CompleteState,
			FirewallState: info.FirewallState,
			WafState:      info.WafState,
			PortNum:       info.PortNum,
			YwdkNum:       info.YwdkNum,
		}).
		Error
}

func (p *DexunDDoSCombos) FindComboByUuid(db *gorm.DB, uuid string) *DexunDDoSCombos {
	info := &DexunDDoSCombos{}
	db.Model(&DexunDDoSCombos{}).Where("uuid = ?", uuid).Find(&info)
	return info
}

func (p *DDoSCombos) GetIdByUuid(db *gorm.DB, combo_id int64) int64 {
	db.Table("DDoSCombos").Where("combo_id = ?", combo_id).Find(&p)
	return p.ID
}

func (p *DDoSCombos) GetSourceById(db *gorm.DB, id int64) string {
	db.Table("DDoSCombos").Where("id = ?", id).Find(&p)
	return p.Source
}

func (p *DDoSCombos) ChangeDDoSCombosStatus(db *gorm.DB, id int64, sell_status int) error {
	return db.Table("DDoSCombos").
		Where("id = ?", id).
		UpdateColumn("sell_status", sell_status).Error
}

func (p *DDoSCombos) AdminGetDDoSByParams(db *gorm.DB, info *form.Filter) (*[]DDoSCombos, int, error) {
	var combos []DDoSCombos
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&DDoSCombos{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&DDoSCombos{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&DDoSCombos{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		//if err := db.Model(&ScdnCombos{}).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
		if err := db.Table("DDoSCombos").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("DDoSCombos").Limit(info.Limit).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *DDoSCombos) UserGetDDoSByParams(db *gorm.DB, info *form.Filter) (*[]DDoSCombos, int, error) {
	var combos []DDoSCombos
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&DDoSCombos{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&DDoSCombos{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&DDoSCombos{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		//if err := db.Model(&ScdnCombos{}).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
		if err := db.Table("DDoSCombos").Where("sell_status = ?", 1).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("DDoSCombos").Limit(info.Limit).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}
