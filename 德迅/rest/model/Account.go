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
	TelNum   string `gorm:"column:tel_num" json:"tel_num"`
	Email    string `gorm:"column:email" json:"email"`
	Balance  int64  `gorm:"column:balance" json:"balance"`
	Create   int    `gorm:"column:create" json:"create"`
	DelTime  int    `gorm:"column:del_time" json:"del_time"`
	Enabled  int    `gorm:"column:enabled" json:"enabled"`
	CertType string `gorm:"column:cert_type" json:"cert_type"`
	Role     string `gorm:"column:role" json:"role"`
	AgentId  int64  `gorm:"column:agent_id" json:"agent_id"`
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

func (p *Account) GetBalanceById(db *gorm.DB, id int64) int64 {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error; err != nil {
		return 0
	} else {
		return p.Balance
	}
}

func (p *Account) GetCertType(db *gorm.DB, id int64) string {
	if err := db.Table("Person").Where("user_id = ?", id).Find(&p).Error; err == nil {
		p.CertType = "Person"
	}
	if err := db.Table("Enterprise").Where("user_id = ?", id).Find(&p).Error; err == nil {
		p.CertType = "corp"
	}
	return p.CertType
}

func (p *Account) UpdateCertType(db *gorm.DB, cert string, id int64) error {
	db.Transaction(func(tx *gorm.DB) error {
		return tx.Table(p.TableName()).
			Where("id = ?", id).
			UpdateColumn("cert_type", cert).Error
	})
	return nil
}

func (p *Account) UpdateBalance(db *gorm.DB, id, balance, recharge int64) error {
	return db.Table(p.TableName()).Where("id = ?", id).UpdateColumn("balance", balance+recharge).Error
}

func (p *Account) GetAccountsByAgentId(db *gorm.DB, agent_id int64) (*[]Account, int) {
	var accounts []Account
	var total int

	if err := db.Table(p.TableName()).Where("agent_id = ?", agent_id).Count(&total).Error; err != nil {
		return nil, 0
	}
	if err := db.Table(p.TableName()).Where("agent_id = ?", agent_id).Find(&accounts).Error; err != nil {
		return nil, 0
	}
	return &accounts, total
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
	//return db.Table(p.TableName()).Where("username = ? and password = ?", username, password).Find(p).Error
	return db.Model(&Account{}).Where("username = ? and password = ?", username, password).Find(p).Error
}

func (p *Account) GetInfo(db *gorm.DB, username string) (*Account, error) {
	if err := db.Table(p.TableName()).Where("username = ?", username).Find(p).Error; err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (p *Account) GetByParams(db *gorm.DB, info *form.AccountFilterForm) (*[]Account, int, error) {
	var account []Account
	var total int

	query := db.Model(&Account{})

	if info.Username != "" {
		query = query.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	if info.TelNum != "" {
		query = query.Where("`tel_num` LIKE ?", "%"+info.TelNum+"%")
	}
	if info.CertType != "" {
		query = query.Where("`cert_type` LIKE ?", "%"+info.CertType+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&account).Error; err != nil {
		return nil, 0, err
	}

	return &account, total, nil
}

func (p *Account) DeleteById(db *gorm.DB, id int64) error {
	db.Transaction(func(tx *gorm.DB) error {
		return tx.Table(p.TableName()).
			Where("id = ?", id).
			UpdateColumn("del_time", int(time.Now().Unix())).Error
	})
	return nil
}
