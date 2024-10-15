package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"strconv"
)

type DexunCombos struct {
	ID            int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	UUID          string `gorm:"column:uuid" json:"uuid"`
	TcName        string `gorm:"column:tc_name" json:"tc_name"`
	KsMoney       int64  `gorm:"column:ks_money" json:"ks_money"`
	ZkMoney       int64  `gorm:"column:zk_money" json:"zk_money"`
	Yms           string `gorm:"column:yms" json:"yms"`
	Ccfy          string `gorm:"column:ccfy" json:"ccfy"`
	Gjwaf         string `gorm:"column:gjwaf" json:"gjwaf"`
	Ywll          string `gorm:"column:ywll" json:"ywll"`
	ProFlow       int64  `gorm:"column:pro_flow" json:"pro_flow"`
	DdosHh        string `gorm:"column:ddos_hh" json:"ddos_hh"`
	DomainNum     int64  `gorm:"column:domain_num" json:"domain_num"`
	CompleteState int64  `gorm:"column:complete_state" json:"complete_state"`
	FirewallState int64  `gorm:"column:firewall_state" json:"firewall_state"`
	WafState      int64  `gorm:"column:waf_state" json:"waf_state"`
	Status        int    `gorm:"column:status" json:"status"`
}

type ScdnCombos struct {
	ID            int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	ComboId       int64  `gorm:"column:combo_id" json:"combo_id"`
	UUID          string `gorm:"column:uuid" json:"uuid"`
	TcName        string `gorm:"column:tc_name" json:"tc_name"`
	KsMoney       int64  `gorm:"column:ks_money" json:"ks_money"`
	ZkMoney       int64  `gorm:"column:zk_money" json:"zk_money"`
	Yms           string `gorm:"column:yms" json:"yms"`
	Ccfy          string `gorm:"column:ccfy" json:"ccfy"`
	Gjwaf         string `gorm:"column:gjwaf" json:"gjwaf"`
	Ywll          string `gorm:"column:ywll" json:"ywll"`
	ProFlow       int64  `gorm:"column:pro_flow" json:"pro_flow"`
	DdosHh        string `gorm:"column:ddos_hh" json:"ddos_hh"`
	DomainNum     int64  `gorm:"column:domain_num" json:"domain_num"`
	CompleteState int64  `gorm:"column:complete_state" json:"complete_state"`
	FirewallState int64  `gorm:"column:firewall_state" json:"firewall_state"`
	WafState      int64  `gorm:"column:waf_state" json:"waf_state"`
	Source        string `gorm:"column:source" json:"source"`
	Status        int    `gorm:"column:status" json:"status"`
	ZyzkMoney     int64  `gorm:"column:zyzk_money" json:"zyzk_money"`
	ZyksMoney     int64  `gorm:"column:zyks_money" json:"zyks_money"`
	SellStatus    int    `gorm:"column:sell_status" json:"sell_status"`
}

type DexunCache struct {
	ID              int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	CacheId         int64  `gorm:"column:cache_id" json:"cache_id"`
	DDUuid          string `gorm:"column:dd_uuid" json:"dd_uuid"`
	CacheUuid       string `gorm:"column:cache_uuid" json:"cache_uuid"`
	DomainUuid      string `gorm:"column:domain_uuid" json:"domain_uuid"`
	Active          int64  `gorm:"column:active" json:"active"`
	UrlMode         string `gorm:"column:urlmode" json:"urlmode"`
	CacheMode       string `gorm:"column:cachemode" json:"cachemode"`
	CachePath       string `gorm:"column:cachepath" json:"cachepath"`
	CacheExtensions string `gorm:"column:cacheextensions" json:"cacheextensions"`
	CacheReg        string `gorm:"column:cachereg" json:"cachereg"`
	TimeOut         string `gorm:"column:timeout" json:"timeout"`
	Weight          string `gorm:"column:weight" json:"weight"`
	CreateTime      string `gorm:"column:createtime" json:"createtime"`
	UpdateTime      string `gorm:"column:updatetime" json:"updatetime"`
}

func (DexunCombos) TableName() string {
	return "DexunCombos"
}

func CreateDexunCache(db *gorm.DB, info List) error {
	items := db.Table("DexunCache")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DexunCache{
		ID:              sf.Generate(),
		CacheUuid:       info.CacheUUID,
		DomainUuid:      info.DomainUUID,
		Active:          info.Active,
		UrlMode:         info.Urlmode,
		CacheMode:       info.Cachemode,
		CachePath:       info.Cachepath,
		CacheExtensions: info.Cacheextensions,
		CacheReg:        info.Cachereg,
		TimeOut:         strconv.FormatInt(info.Timeout, 10),
		Weight:          strconv.FormatInt(info.Weight, 10),
		CreateTime:      info.Createtime,
		UpdateTime:      info.Updatetime,
	}).Error
}

func CreateDexunCombos(db *gorm.DB, info Data) error {
	items := db.Table("DexunCombos")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DexunCombos{
		ID:            sf.Generate(),
		UUID:          info.UUID,
		TcName:        info.TcName,
		KsMoney:       info.KsMoney,
		ZkMoney:       info.ZkMoney,
		Yms:           info.ProNote.Yms,
		Ccfy:          info.ProNote.Ccfy,
		Gjwaf:         info.ProNote.Gjwaf,
		Ywll:          info.ProNote.Ywll,
		ProFlow:       info.ProFlow,
		DdosHh:        info.DdosHh,
		DomainNum:     info.DomainNum,
		CompleteState: info.CompleteState,
		FirewallState: info.FirewallState,
		WafState:      info.WafState,
		Status:        1,
	}).Error
}

func CopyDexunCombos(db *gorm.DB, id int64) error {
	items := db.Table("ScdnCombos")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	var combo DexunCombos
	var service ScdnCombos
	if err := db.Table("DexunCombos").Where("id = ?", id).Find(&combo).Error; err != nil {
		return err
	}

	db.Table("ScdnCombos").Where("combo_id = ?", combo.ID).Find(&service)
	if service.ID == 0 {
		items.Create(&ScdnCombos{
			ID:            sf.Generate(),
			ComboId:       combo.ID,
			UUID:          combo.UUID,
			TcName:        combo.TcName,
			KsMoney:       combo.KsMoney,
			ZkMoney:       combo.ZkMoney,
			Yms:           combo.Yms,
			Ccfy:          combo.Ccfy,
			Gjwaf:         combo.Gjwaf,
			Ywll:          combo.Ywll,
			ProFlow:       combo.ProFlow,
			DdosHh:        combo.DdosHh,
			DomainNum:     combo.DomainNum,
			CompleteState: combo.CompleteState,
			FirewallState: combo.FirewallState,
			WafState:      combo.WafState,
			Source:        "Dexun",
			Status:        combo.Status,
			SellStatus:    1,
			ZyksMoney:     3 * combo.KsMoney,
			ZyzkMoney:     3 * combo.ZkMoney,
		})
	} else {
		items.Where("combo_id = ?", combo.ID).Updates(&ScdnCombos{
			TcName:        combo.TcName,
			KsMoney:       combo.KsMoney,
			ZkMoney:       combo.ZkMoney,
			Yms:           combo.Yms,
			Ccfy:          combo.Ccfy,
			Gjwaf:         combo.Gjwaf,
			Ywll:          combo.Ywll,
			ProFlow:       combo.ProFlow,
			DdosHh:        combo.DdosHh,
			DomainNum:     combo.DomainNum,
			CompleteState: combo.CompleteState,
			FirewallState: combo.FirewallState,
			WafState:      combo.WafState,
			Source:        "Dexun",
			ZyksMoney:     3 * combo.KsMoney,
			ZyzkMoney:     3 * combo.ZkMoney,
		})
	}
	return nil
}

func (p *ScdnCombos) ChangeScdnCombosStatus(db *gorm.DB, id int64, sell_status int) error {
	return db.Table("ScdnCombos").
		Where("id = ?", id).
		UpdateColumn("sell_status", sell_status).Error
}

func (p *ScdnCombos) GetIdByUuid(db *gorm.DB, uuid int64) int64 {
	if err := db.Table("ScdnCombos").Where("combo_id = ?", uuid).Find(&p).Error; err != nil {
		return 0
	}
	return p.ID
}

func (p *ScdnCombos) GetSourceById(db *gorm.DB, id int64) string {
	if err := db.Table("ScdnCombos").Where("id = ?", id).Find(&p).Error; err != nil {
		return ""
	}
	return p.Source
}

func (p *DexunCombos) FindComboByUuid(db *gorm.DB, uuid string) *DexunCombos {
	info := &DexunCombos{}
	db.Model(&DexunCombos{}).Where("uuid = ?", uuid).Find(&info)
	return info
}

func (p *DexunCombos) UpdateComboByUuid(db *gorm.DB, uuid string, info Data) error {
	return db.Table(p.TableName()).
		Where("uuid = ?", uuid).
		Where("status = ?", 1).
		Updates(DexunCombos{
			TcName:        info.TcName,
			KsMoney:       info.KsMoney,
			ZkMoney:       info.ZkMoney,
			Yms:           info.ProNote.Yms,
			Ccfy:          info.ProNote.Ccfy,
			Gjwaf:         info.ProNote.Gjwaf,
			Ywll:          info.ProNote.Ywll,
			ProFlow:       info.ProFlow,
			DdosHh:        info.DdosHh,
			DomainNum:     info.DomainNum,
			CompleteState: info.CompleteState,
			FirewallState: info.FirewallState,
			WafState:      info.WafState,
		}).
		Error
}

func (p *DexunCombos) GetDCByParams(db *gorm.DB, info *form.ScdnFilterForm) (*[]DexunCombos, int, error) {
	var combos []DexunCombos
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&DexunCombos{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("`status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&DexunCombos{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("`status` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("status = ?", 0).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&DexunCombos{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("`status` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("status = ?", 0).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&DexunCombos{}).Where("status = ?", 0).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Where("status = ?", 0).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *ScdnCombos) AdminGetSCByParams(db *gorm.DB, info *form.ScdnFilterForm) (*[]ScdnCombos, int, error) {
	var combos []ScdnCombos
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&ScdnCombos{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&ScdnCombos{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&ScdnCombos{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
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
		if err := db.Table("ScdnCombos").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("ScdnCombos").Limit(info.Limit).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *ScdnCombos) UserGetSCByParams(db *gorm.DB, info *form.ScdnFilterForm) (*[]ScdnCombos, int, error) {
	var combos []ScdnCombos
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&ScdnCombos{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&ScdnCombos{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&ScdnCombos{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
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
		if err := db.Table("ScdnCombos").Where("sell_status = ?", 1).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("ScdnCombos").Limit(info.Limit).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}
