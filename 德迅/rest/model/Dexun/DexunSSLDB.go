package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DexunSSLList struct {
	ID          int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	UUID        string `gorm:"column:uuid" json:"uuid"`
	PTypeId     int64  `gorm:"column:p_type_id" json:"p_type_id"`
	PTypeName   string `gorm:"column:p_type_name" json:"p_type_name"`
	SSLName     string `gorm:"column:ssl_name" json:"ssl_name"`
	SSLType     string `gorm:"column:ssl_type" json:"ssl_type"`
	DomainType  string `gorm:"column:domain_type" json:"domain_type"`
	DomainNum   string `gorm:"column:domain_num" json:"domain_num"`
	SSLCode     string `gorm:"column:ssl_code" json:"ssl_code"`
	KsMoney     int64  `gorm:"column:ks_money" json:"ks_money"`
	EyMoney     int64  `gorm:"column:ey_money" json:"ey_money"`
	PNote       string `gorm:"column:p_note" json:"p_note"`
	Term        string `gorm:"column:term" json:"term"`
	SType       int64  `gorm:"column:s_type" json:"s_type"`
	MarketMoney int64  `gorm:"column:market_money" json:"market_money"`
}

type SSLList struct {
	ID          int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	SSLId       int64  `gorm:"column:ssl_id" json:"ssl_id"`
	UUID        string `gorm:"column:uuid" json:"uuid"`
	PTypeId     int64  `gorm:"column:p_type_id" json:"p_type_id"`
	PTypeName   string `gorm:"column:p_type_name" json:"p_type_name"`
	SSLName     string `gorm:"column:ssl_name" json:"ssl_name"`
	SSLType     string `gorm:"column:ssl_type" json:"ssl_type"`
	DomainType  string `gorm:"column:domain_type" json:"domain_type"`
	DomainNum   string `gorm:"column:domain_num" json:"domain_num"`
	SSLCode     string `gorm:"column:ssl_code" json:"ssl_code"`
	KsMoney     int64  `gorm:"column:ks_money" json:"ks_money"`
	EyMoney     int64  `gorm:"column:ey_money" json:"ey_money"`
	PNote       string `gorm:"column:p_note" json:"p_note"`
	Term        string `gorm:"column:term" json:"term"`
	SType       int64  `gorm:"column:s_type" json:"s_type"`
	MarketMoney int64  `gorm:"column:market_money" json:"market_money"`
	SellStatus  int    `gorm:"column:sell_status" json:"sell_status"`
	ZyksMoney   int64  `gorm:"column:zyks_money" json:"zyks_money"`
	Zyeymoney   int64  `gorm:"column:zyey_money" json:"zyey_money"`
}

func (DexunSSLList) TableName() string {
	return "DexunSSLList"
}

func CreateDexunSSLList(db *gorm.DB, info List) error {
	items := db.Table("DexunSSLList")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DexunSSLList{
		ID:          sf.Generate(),
		UUID:        info.UUID,
		PTypeId:     info.PTypeID,
		PTypeName:   info.PTypeName,
		SSLName:     info.SSLName,
		SSLType:     info.SSLType,
		DomainType:  info.DomainType,
		DomainNum:   info.DomainNum,
		SSLCode:     info.SSLCode,
		KsMoney:     info.KsMoney,
		EyMoney:     info.EyMoney,
		PNote:       info.PNote,
		Term:        info.Term,
		SType:       info.SType,
		MarketMoney: info.MarketMoney,
	}).Error
}

func (p *DexunSSLList) UpdateSSLByUUID(db *gorm.DB, uuid string, info List) error {
	return db.Table(p.TableName()).
		Where("uuid = ?", uuid).
		Updates(&DexunSSLList{
			PTypeId:     info.PTypeID,
			PTypeName:   info.PTypeName,
			SSLName:     info.SSLName,
			SSLType:     info.SSLType,
			DomainType:  info.DomainType,
			DomainNum:   info.DomainNum,
			SSLCode:     info.SSLCode,
			KsMoney:     info.KsMoney,
			EyMoney:     info.EyMoney,
			PNote:       info.PNote,
			Term:        info.Term,
			SType:       info.SType,
			MarketMoney: info.MarketMoney,
		}).Error
}

func (p *SSLList) GetMoneyById(db *gorm.DB, id int64) int64 {
	if err := db.Table("SSLList").Where("id = ?", id).Find(&p); err != nil {
		return 0
	} else {
		return p.ZyksMoney
	}
}

func (p *DexunSSLList) FindComboByUuid(db *gorm.DB, uuid string) *DexunSSLList {
	info := &DexunSSLList{}
	db.Model(&DexunSSLList{}).Where("uuid = ?", uuid).Find(&info)
	return info
}

func CopyDxSSLList(db *gorm.DB) error {
	items := db.Table("SSLList")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	var combos []DexunSSLList
	var service SSLList
	if err := db.Table("DexunSSLList").Find(&combos).Error; err != nil {
		return err
	}

	for _, combo := range combos {
		db.Table("SSLList").Where("ssl_id = ?", combo.ID).Find(&service)
		if service.ID == 0 {
			items.Create(&SSLList{
				ID:          sf.Generate(),
				SSLId:       combo.ID,
				UUID:        combo.UUID,
				PTypeId:     combo.PTypeId,
				PTypeName:   combo.PTypeName,
				SSLName:     combo.SSLName,
				SSLType:     combo.SSLType,
				DomainType:  combo.DomainType,
				DomainNum:   combo.DomainNum,
				SSLCode:     combo.SSLCode,
				KsMoney:     combo.KsMoney,
				EyMoney:     combo.EyMoney,
				PNote:       combo.PNote,
				Term:        combo.Term,
				SType:       combo.SType,
				MarketMoney: combo.MarketMoney,
				SellStatus:  1,
				ZyksMoney:   3 * combo.KsMoney,
				Zyeymoney:   3 * combo.EyMoney,
			})
		} else {
			items.Where("ssl_id = ?", combo.ID).Updates(&SSLList{
				PTypeId:     combo.PTypeId,
				PTypeName:   combo.PTypeName,
				SSLName:     combo.SSLName,
				SSLType:     combo.SSLType,
				DomainType:  combo.DomainType,
				DomainNum:   combo.DomainNum,
				SSLCode:     combo.SSLCode,
				KsMoney:     combo.KsMoney,
				EyMoney:     combo.EyMoney,
				PNote:       combo.PNote,
				Term:        combo.Term,
				SType:       combo.SType,
				MarketMoney: combo.MarketMoney,
				SellStatus:  1,
				ZyksMoney:   3 * combo.KsMoney,
				Zyeymoney:   3 * combo.EyMoney,
			})
		}
	}
	return nil
}

func (p *SSLList) UserGetSSLByParams(db *gorm.DB, ssl_type string, info *form.Filter) (*[]SSLList, int, error) {
	var combos []SSLList
	var total int

	if info.Field != "" {
		if info.Field == "ssl_name" {
			if err := db.Model(&SSLList{}).Where("`ssl_name` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ssl_name` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "zyks_money" {
			if err := db.Model(&SSLList{}).Where("`zyks_money` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`zyks_money` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&SSLList{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		//if err := db.Model(&ScdnCombos{}).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
		if err := db.Table("SSLList").Where("ssl_type = ?", ssl_type).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("SSLList").Limit(info.Limit).Where("ssl_type = ?", ssl_type).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *SSLList) AdminGetSSLByParams(db *gorm.DB, ssl_type string, info *form.Filter) (*[]SSLList, int, error) {
	var combos []SSLList
	var total int

	if info.Field != "" {
		if info.Field == "ssl_name" {
			if err := db.Model(&SSLList{}).Where("`ssl_name` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ssl_name` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "zyks_money" {
			if err := db.Model(&SSLList{}).Where("`zyks_money` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`zyks_money` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&SSLList{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("ssl_type = ?", ssl_type).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		//if err := db.Model(&ScdnCombos{}).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
		if err := db.Table("SSLList").Where("ssl_type = ?", ssl_type).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("SSLList").Where("ssl_type = ?", ssl_type).Limit(info.Limit).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *SSLList) GetUUIDById(db *gorm.DB, id int64) string {
	if err := db.Table("SSLList").Where("id = ?", id).Find(&p).Error; err != nil {
		return ""
	}
	return p.UUID
}
