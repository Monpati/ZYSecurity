package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

var salt = utils.GetRandomString(32)

type Account struct {
	Id       int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
	TelNum   string `gorm:"tel_num" json:"tel_num"`
	Email    string `gorm:"email" json:"email"`
	Balance  int    `gorm:"balance" json:"balance"`
	Create   int    `gorm:"create" json:"create"`
	DelTime  int    `gorm:"del_time" json:"del_time"`
	Enabled  int    `gorm:"enabled" json:"enabled"`
	CertType string `gorm:"cert_type" json:"cert_type"`
	Role     string `gorm:"role" json:"role"`
}

func (Account) TableName() string {
	return "User"
}

func AccountCreate(db *gorm.DB, info *form.AccountInfo) error {
	items := db.Table("User")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	hashedPassword, _ := utils.HashSalt(info.Password, salt)
	return items.Create(&Account{
		Id:       sf.Generate(),
		Username: info.Username,
		Password: hashedPassword,
		Salt:     salt,
		TelNum:   info.TelNum,
		Email:    info.Email,
		Balance:  0,
		Create:   int(time.Now().Unix()),
		Role:     "user",
	}).Error
}

func (p *Account) GetIdByAccountName(db *gorm.DB, username string) int64 {
	if err := db.Table(p.TableName()).Where("username = ?", username).Find(&p).Error; err != nil {
		return 0
	} else {
		return p.Id
	}
}

func (p *Account) GetCertType(db *gorm.DB, username string) string {
	if err := db.Table("Person").Where("username = ?", username).Find(&p.Id).Error; err == nil {
		p.CertType = "Person"
	}
	if err := db.Table("Enterprise").Where("username = ?", username).Find(&p.Id).Error; err == nil {
		p.CertType = "corp"
	}
	return p.CertType
}

func (p *Account) UpdateCertType(db *gorm.DB, username, cert string) error {
	return db.Table(p.TableName()).
		Where("username = ?", username).
		UpdateColumn("cert_type", cert).Error
}

func (p *Account) FindSaltByAccountName(db *gorm.DB, username string) string {
	if err := db.Table(p.TableName()).Where("username = ?", username).Find(&p).Error; err != nil {
		return ""
	} else {
		return p.Salt
	}
}

func (p *Account) AdminFindByAccountName(db *gorm.DB, username, password string) error {
	return db.Table(p.TableName()).Where("username = ? and password = ? and role = ?", username, password, "admin").Find(p).Error
}

func (p *Account) GetRole(db *gorm.DB, username string) error {
	return db.Table(p.TableName()).Where("username = ?", username).Find(p).Error
}

func (p *Account) FindByAccountName(db *gorm.DB, username, password string) error {
	//return db.Model(&Account{}).Where("username = ? and password = ?", username, password).Find(p).Error
	return db.Table(p.TableName()).Where("username = ? and password = ?", username, password).Find(p).Error
}

func (p *Account) GetInfo(db *gorm.DB, id int64) (*Account, error) {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&Account{}).Error; err != nil {
		return nil, err
	} else {
		return &Account{}, nil
	}
}

func (p *Account) GetByParams(db *gorm.DB, info *form.AccountFilterForm) (*[]Account, int, error) {
	var account []Account
	var total int

	if info.Field != "" {
		if info.Field == "username" {
			if err := db.Model(&Account{}).Where("`username` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`username` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&account).Error; err != nil {
				return nil, 0, err
			}
			return &account, total, nil
		}

		if info.Field == "tel_num" {
			if err := db.Model(&Account{}).Where("`tel_num` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tel_num` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&account).Error; err != nil {
				return nil, 0, err
			}
			return &account, total, nil
		}

		if info.Field == "email" {
			if err := db.Model(&Account{}).Where("`email` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`email` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&account).Error; err != nil {
				return nil, 0, err
			}
			return &account, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&Account{}).Where("del_time = ?", 0).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Where("del_time = ?", 0).Offset(info.Offset).Find(&account).Error; err != nil {
			return nil, 0, err
		}
		return &account, total, nil
	}
	return &account, total, nil
}
