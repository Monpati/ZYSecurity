package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type OperationLog struct {
	Id        int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	UserId    int64      `gorm:"column:user_id" json:"user_id"`
	ReqUrl    string     `gorm:"column:req_url" json:"req_url"`
	OriginUrl string     `gorm:"column:origin_url" json:"origin_url"`
	UserAgent string     `gorm:"column:user_agent" json:"user_agent"`
	Request   utils.JSON `gorm:"column:request" json:"request"`
	ReqTime   int        `gorm:"column:req_time" json:"req_time"`
	Create    int        `gorm:"column:create" json:"create"`
}

func (OperationLog) TableName() string {
	return "Logs"
}

func OperationLogCreate(db *gorm.DB, info *form.OperationLog) error {
	items := db.Table("Logs")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&OperationLog{
		Id:        sf.Generate(),
		UserId:    info.UserId,
		ReqUrl:    info.ReqUrl,
		OriginUrl: info.OriginUrl,
		UserAgent: info.UserAgent,
		Request:   info.Request,
		ReqTime:   info.ReqTime,
		Create:    int(time.Now().Unix()),
	}).Error
}
