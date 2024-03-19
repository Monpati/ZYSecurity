package model

import (
	"Dexun/form"
	"Dexun/utils"
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"io"
	"time"
)

var salt = utils.GetRandomString(32)

type Account struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
	TelNum   string `gorm:"tel_num" json:"tel_num"`
	Email    string `gorm:"email" json:"email"`
	Balance  int    `gorm:"balance" json:"balance"`
	Create   int    `gorm:"create" json:"create"`
	Delete   int    `gorm:"delete" json:"delete"`
	CertType string `gorm:"cert_type" json:"cert_type"`
}

func (Account) TableName() string {
	return "User"
}

func HashSalt(str, salt string) (string, error) {
	m := md5.New()
	io.WriteString(m, str)
	m.Sum(nil)
	io.WriteString(m, salt)
	return hex.EncodeToString(m.Sum(nil)), nil
}

func AccountCreate(db *gorm.DB, info *form.AccountInfo) error {
	items := db.Table("User")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	hashedPassword, _ := HashSalt(info.Password, salt)
	return items.Create(&Account{
		Id:       sf.Generate(),
		Username: info.Username,
		Password: hashedPassword,
		Salt:     salt,
		TelNum:   info.TelNum,
		Email:    info.Email,
		Balance:  0,
		Create:   int(time.Now().Unix()),
	}).Error
}

func (p *Account) FindSaltByAccountName(db *gorm.DB, username string) string {
	if err := db.Table(p.TableName()).Where("username = ?", username).Find(&p).Error; err != nil {
		return ""
	} else {
		return p.Salt
	}
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
