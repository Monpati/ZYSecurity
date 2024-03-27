package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type CorpCert struct {
	Id            int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	UserId        int64  `gorm:"column:user_id" json:"user_id"`
	CorpName      string `gorm:"column:corp_name" json:"corp_name"`
	RegNum        string `gorm:"column:reg_num" json:"reg_num"`
	LgMan         string `gorm:"column:lgman" json:"lgman"`
	LgPersonNum   string `gorm:"column:lgperson_num" json:"lgperson_num"`
	CorpAddress   string `gorm:"column:corp_address" json:"corp_address"`
	CorpDoc       string `gorm:"column:corp_doc" json:"corp_doc"`
	Status        int    `gorm:"column:status" json:"status"`
	LgPersonFront string `gorm:"column:lgperson_front" json:"lgperson_front"`
	LgPersonBack  string `gorm:"column:lgperson_back" json:"lgperson_back"`
	DelTime       int    `gorm:"column:del_time" json:"del_time"`
	Create        int    `gorm:"column:create" json:"create"`
}

func (CorpCert) TableName() string {
	return "Enterprise"
}

func CorpCertCreate(db *gorm.DB, info *CorpCert) error {
	items := db.Table("Enterprise")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&CorpCert{
		Id:            sf.Generate(),
		UserId:        info.UserId,
		CorpName:      info.CorpName,
		RegNum:        info.RegNum,
		LgMan:         info.LgMan,
		LgPersonNum:   info.LgPersonNum,
		CorpAddress:   info.CorpAddress,
		CorpDoc:       info.CorpDoc,
		LgPersonFront: info.LgPersonFront,
		LgPersonBack:  info.LgPersonBack,
		Status:        2,
		DelTime:       0,
		Create:        int(time.Now().Unix()),
	}).Error
}

func (p *CorpCert) UpdateStatus(db *gorm.DB, userid int64, status int) error {
	return db.Table(p.TableName()).
		Where("user_id = ?", userid).
		UpdateColumn("status", status).Error
}

func (p *CorpCert) GetStatusByName(db *gorm.DB, username string) error {
	return db.Table(p.TableName()).Where("username = ?", username).Find(p.Status).Error
}

func (p *CorpCert) GetStatusById(db *gorm.DB, id int64) error {
	return db.Table(p.TableName()).Where("user_id = ?", id).Find(p).Error
}

func (p *CorpCert) GetByParams(db *gorm.DB, info *form.CorpFilter) (*[]CorpCert, int, error) {
	var corp []CorpCert
	var total int

	if info.Field != "" {
		if info.Field == "username" {
			if err := db.Model(&CorpCert{}).Where("`real_name` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`real_name` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&corp).Error; err != nil {
				return nil, 0, err
			}
			return &corp, total, nil
		}

		if info.Field == "tel_num" {
			if err := db.Model(&CorpCert{}).Where("`card_id` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`card_id` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&corp).Error; err != nil {
				return nil, 0, err
			}
			return &corp, total, nil
		}

		if info.Field == "email" {
			if err := db.Model(&CorpCert{}).Where("`sex` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`sex` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&corp).Error; err != nil {
				return nil, 0, err
			}
			return &corp, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&CorpCert{}).Where("del_time = ?", 0).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Where("del_time = ?", 0).Offset(info.Offset).Find(&corp).Error; err != nil {
			return nil, 0, err
		}
		return &corp, total, nil
	}
	return &corp, total, nil
}
