package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Agent struct {
	Id       int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	TelNum   string `gorm:"column:tel_num" json:"tel_num"`
	Create   int    `gorm:"column:create" json:"create"`
	Salt     string `gorm:"column:salt" json:"salt"`
	Email    string `gorm:"column:email" json:"email"`
	Amount   int64  `gorm:"column:amount" json:"amount"`
	Status   int    `gorm:"column:status" json:"status"`
	Role     string `gorm:"column:role" json:"role"`
}

type Invite struct {
	Id     int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	Code   string `gorm:"column:code" json:"code"`
	Status int    `gorm:"column:status" json:"status"`
}

func (Agent) TableName() string {
	return "Agent"
}

func InviteCodeAdd(db *gorm.DB, info *form.InviteCode) error {
	items := db.Table("Invite")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&Invite{
		Id:     sf.Generate(),
		Code:   info.Code,
		Status: 0,
	}).Error
}

func AgentCreate(db *gorm.DB, info *form.AgentInfo) error {
	items := db.Table("Agent")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	hashedPassword, _ := utils.HashSalt(info.Password, salt)
	return items.Create(&Agent{
		Id:       sf.Generate(),
		Username: info.Username,
		Password: hashedPassword,
		Salt:     salt,
		TelNum:   info.TelNum,
		Email:    info.Email,
		Amount:   0,
		Create:   int(time.Now().Unix()),
		Status:   1,
		Role:     "agent",
	}).Error
}

func (p *Agent) FindAgentByUP(db *gorm.DB, username, password string) error {
	return db.Model(&Agent{}).Where("username = ? and password = ?", username, password).Find(p).Error
}

func (p *Agent) FindAgentByUsername(db *gorm.DB, username string) error {
	return db.Model(&Agent{}).Where("username = ?", username).Find(p).Error
}

func (p *Agent) FindSaltByAgentName(db *gorm.DB, username string) string {
	if err := db.Table(p.TableName()).Where("username = ?", username).Find(&p).Error; err != nil {
		return ""
	} else {
		return p.Salt
	}
}

func (p *Agent) UpdateStatus(db *gorm.DB, agentid int64, status int) error {
	return db.Table(p.TableName()).
		Where("id = ?", agentid).
		UpdateColumn("status", status).Error
}

func (p *Invite) UpdateInviteStatus(db *gorm.DB, id int64, status int) error {
	return db.Table("Invite").
		Where("id = ?", id).
		UpdateColumn("status", status).
		Error
}

func (p *Agent) GetByParams(db *gorm.DB, info *form.AgentFilterForm) (*[]Agent, int, error) {
	var agent []Agent
	var total int

	query := db.Model(&Agent{})

	if info.Username != "" {
		query = query.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	if info.TelNum != "" {
		query = query.Where("`tel_num` LIKE ?", "%"+info.TelNum+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&agent).Error; err != nil {
		return nil, 0, err
	}
	return &agent, total, nil
}

func (p *Agent) GetUsersByAgent(db *gorm.DB, id int64, info *form.Filter) (*[]Account, int, error) {
	var account []Account
	var total int

	if info.Field != "" {
		if info.Field == "username" {
			if err := db.Model(&Account{}).Where("`username` LIKE ?", "%"+info.Value+"%").Where("agent_id = ?", id).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`username` LIKE ?", "%"+info.Value+"%").Where("agent_id = ?", id).Offset(info.Offset).Find(&account).Error; err != nil {
				return nil, 0, err
			}
			return &account, total, nil
		}

		if info.Field == "tel_num" {
			if err := db.Model(&Account{}).Where("`tel_num` LIKE ?", "%"+info.Value+"%").Where("agent_id = ?", id).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tel_num` LIKE ?", "%"+info.Value+"%").Where("agent_id = ?", id).Offset(info.Offset).Find(&account).Error; err != nil {
				return nil, 0, err
			}
			return &account, total, nil
		}

		if info.Field == "email" {
			if err := db.Model(&Account{}).Where("`email` LIKE ?", "%"+info.Value+"%").Where("agent_id = ?", id).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`email` LIKE ?", "%"+info.Value+"%").Where("agent_id = ?", id).Offset(info.Offset).Find(&account).Error; err != nil {
				return nil, 0, err
			}
			return &account, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&Account{}).Where("agent_id = ?", id).Where("role = ?", "user").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Where("agent_id = ?", id).Where("role = ?", "user").Offset(info.Offset).Find(&account).Error; err != nil {
			return nil, 0, err
		}
		return &account, total, nil
	}
	return &account, total, nil
}

func (p *Agent) GetIdByName(db *gorm.DB, username string) error {
	return db.Table(p.TableName()).Where("username = ?", username).Find(&p).Error
}
