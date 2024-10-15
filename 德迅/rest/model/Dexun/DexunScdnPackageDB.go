package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type ScdnPackage struct {
	ID           int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	UUID         string `gorm:"column:uuid" json:"uuid"`
	ProtectNote  string `gorm:"column:protect_note" json:"protect_note"`
	ProtectLv    int64  `gorm:"column:protect_lv" json:"protect_lv"`
	ProtectPrice int64  `gorm:"column:protect_price" json:"protect_price"`
	Source       string `gorm:"column:source" json:"source"`
	Status       int    `gorm:"column:status" json:"status"`
	Type         string `gorm:"column:type" json:"type"`
}

type PackageService struct {
	ID           int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	PackageId    int64  `gorm:"column:package_id" json:"package_id"`
	UUID         string `gorm:"column:uuid" json:"uuid"`
	ProtectNote  string `gorm:"column:protect_note" json:"protect_note"`
	ProtectLv    int64  `gorm:"column:protect_lv" json:"protect_lv"`
	ProtectPrice int64  `gorm:"column:protect_price" json:"protect_price"`
	Source       string `gorm:"column:source" json:"source"`
	Status       int    `gorm:"column:status" json:"status"`
	SellStatus   int    `gorm:"column:sell_status" json:"sell_status"`
	Type         string `gorm:"column:type" json:"type"`
	SellPrice    int64  `gorm:"column:sell_price" json:"sell_price"`
}

func (ScdnPackage) TableName() string {
	return "ScdnPackage"
}

func CreateScdnPackage(db *gorm.DB, info Data) error {
	items := db.Table("ScdnPackage")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&ScdnPackage{
		ID:           sf.Generate(),
		UUID:         info.UUID,
		ProtectNote:  info.ProtectNote,
		ProtectLv:    info.ProtectLV,
		ProtectPrice: info.ProtectPrice,
		Source:       "DeXun",
		Status:       1,
		Type:         info.Type,
	}).Error
}

func CopyPackageService(db *gorm.DB, id int64) error {
	items := db.Table("PackageService")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	var packages ScdnPackage
	var data PackageService
	if err := db.Table("ScdnPackage").Where("id = ?", id).Find(&packages).Error; err != nil {
		return err
	}

	db.Table("PackageService").Where("package_id = ?", packages.ID).Find(&data)
	if data.ID == 0 {
		items.Create(&PackageService{
			ID:           sf.Generate(),
			PackageId:    packages.ID,
			UUID:         packages.UUID,
			ProtectNote:  packages.ProtectNote,
			ProtectLv:    packages.ProtectLv,
			ProtectPrice: packages.ProtectPrice,
			Status:       packages.Status,
			SellStatus:   1,
			Source:       packages.Source,
			Type:         packages.Type,
		})
	} else {
		items.Where("package_id = ?", packages.ID).Updates(&PackageService{
			ProtectNote:  packages.ProtectNote,
			ProtectLv:    packages.ProtectLv,
			ProtectPrice: packages.ProtectPrice,
			Source:       packages.Source,
			Type:         packages.Type,
		})
	}
	return nil
}

func (p *PackageService) ChangeScdnPackageStatus(db *gorm.DB, id int64, sell_status int) error {
	return db.Table("PackageService").
		Where("id= ?", id).
		UpdateColumn("sell_status = ?", sell_status).Error
}

func (p *ScdnPackage) FindPackageByUuid(db *gorm.DB, uuid string) *ScdnPackage {
	info := &ScdnPackage{}
	db.Model(&ScdnPackage{}).Where("uuid = ?", uuid).Find(&info)
	return info
}

func (p *ScdnPackage) UpdatePackageByUuid(db *gorm.DB, uuid string, info Data) error {
	return db.Table(p.TableName()).
		Where("uuid = ?", uuid).
		Where("status = ?", 1).
		Updates(ScdnPackage{
			ProtectNote:  info.ProtectNote,
			ProtectLv:    info.ProtectLV,
			ProtectPrice: info.ProtectPrice,
			Type:         info.Type,
			Source:       "DeXun",
		}).
		Error
}

func (p *ScdnPackage) GetDxPackagesByParams(db *gorm.DB, info *form.ScdnFilterForm) (*[]ScdnPackage, int, error) {
	var packages []ScdnPackage
	var total int

	if info.Field != "protect_note" {
		if info.Field == "pro" {
			if err := db.Model(&ScdnPackage{}).Where("`protect_note` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`protect_note` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&packages).Error; err != nil {
				return nil, 0, err
			}
			return &packages, total, nil
		}

		if info.Field == "protect_lv" {
			if err := db.Model(&ScdnPackage{}).Where("`protect_lv` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`protect_lv` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&packages).Error; err != nil {
				return nil, 0, err
			}
			return &packages, total, nil
		}

		if info.Field == "email" {
			if err := db.Model(&ScdnPackage{}).Where("`protect_price` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`protect_price` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&packages).Error; err != nil {
				return nil, 0, err
			}
			return &packages, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&ScdnPackage{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Offset(info.Offset).Find(&packages).Error; err != nil {
			return nil, 0, err
		}
		return &packages, total, nil
	}
	return &packages, total, nil
}

func (p *PackageService) GetPackagesByParams(db *gorm.DB, info *form.ScdnFilterForm) (*[]PackageService, int, error) {
	var packages []PackageService
	var total int

	if info.Field != "protect_note" {
		if info.Field == "pro" {
			if err := db.Model(&PackageService{}).Where("`protect_note` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`protect_note` LIKE ?", "%"+info.Value+"%").Where("source = ?", "DeXun").Offset(info.Offset).Find(&packages).Error; err != nil {
				return nil, 0, err
			}
			return &packages, total, nil
		}

		if info.Field == "protect_lv" {
			if err := db.Model(&PackageService{}).Where("`protect_lv` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`protect_lv` LIKE ?", "%"+info.Value+"%").Where("source = ?", "DeXun").Offset(info.Offset).Find(&packages).Error; err != nil {
				return nil, 0, err
			}
			return &packages, total, nil
		}

		if info.Field == "email" {
			if err := db.Model(&PackageService{}).Where("`protect_price` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`protect_price` LIKE ?", "%"+info.Value+"%").Where("source = ?", "DeXun").Offset(info.Offset).Find(&packages).Error; err != nil {
				return nil, 0, err
			}
			return &packages, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Table("PackageService").Where("source = ?", "DeXun").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("PackageService").Limit(info.Limit).Where("source = ?", "DeXun").Offset(info.Offset).Find(&packages).Error; err != nil {
			return nil, 0, err
		}
		return &packages, total, nil
	}
	return &packages, total, nil
}
