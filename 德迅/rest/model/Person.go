package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
	"mime/multipart"
	"path"
	"strconv"
	"time"
)

type PersonalCert struct {
	Id        int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	UserId    int64  `gorm:"column:user_id" json:"user_id"`
	RealName  string `gorm:"column:real_name" json:"real_name"`
	CardId    string `gorm:"column:card_id" json:"card_id"`
	Sex       string `gorm:"column:sex" json:"sex"`
	Birthday  string `gorm:"column:birthday" json:"birthday"`
	City      string `gorm:"column:city" json:"city"`
	Status    int    `gorm:"column:status" json:"status"`
	CardFront string `gorm:"column:card_front" json:"card_front"`
	CardBack  string `gorm:"column:card_back" json:"card_back"`
	DelTime   int    `gorm:"column:del_time" json:"del_time"`
	Create    int    `gorm:"column:create" json:"create"`
}

type CardStore struct {
	Username string `json:"username"`
	Kind     string `json:"kind"`
}

func (PersonalCert) TableName() string {
	return "Person"
}

// card_front、card_back
func UploadCard(file *multipart.FileHeader) string {
	var saveDir string
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg": true,
		".png": true,
		"jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "格式不符"
	}
	fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
	currentTime := time.Now().Format("20240322")
	saveDir = path.Join("/opt/data/cert/"+currentTime+"/"+salt, fileUnixName+extName)
	return saveDir
}

func (p *PersonalCert) DeletedById(db *gorm.DB, id int64) error {
	return db.Table(p.TableName()).
		Where("user_id = ?", id).
		UpdateColumn("del_time", int(time.Now().Unix())).
		UpdateColumn("status", 0).Error
}

func (p *PersonalCert) GetStatusById(db *gorm.DB, id int64) error {
	return db.Table(p.TableName()).Where("user_id = ?", id).Where("del_time = ?", 0).Find(p).Error
}

func (p *Account) GetIdByName(db *gorm.DB, username string) error {
	return db.Table("User").Where("username = ?", username).Find(p).Error
}

func (p *PersonalCert) UpdateStatus(db *gorm.DB, userid int64, status int) error {
	return db.Table(p.TableName()).
		Joins("User.cert_type = ?", "person").
		Where("user_id = ?", userid).
		UpdateColumn("status", status).Error
}

func (p *PersonalCert) GetByParams(db *gorm.DB, info *form.PersonFilter) (*[]PersonalCert, int, error) {
	var cert []PersonalCert
	var total int
	query := db.Model(&PersonalCert{})

	if info.RealName != "" {
		query = query.Where("`real_name` LIKE ?", "%"+info.RealName+"%")
	}
	if info.CardId != "" {
		query = query.Where("`card_id` LIKE ?", "%"+info.CardId+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&cert).Error; err != nil {
		return nil, 0, err
	}
	return &cert, total, nil
}

func PersonalCertCreate(db *gorm.DB, info *PersonalCert) error {
	items := db.Table("Person")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&PersonalCert{
		Id:        sf.Generate(),
		UserId:    info.UserId,
		RealName:  info.RealName,
		CardId:    info.CardId,
		Sex:       info.Sex,
		Birthday:  info.Birthday,
		City:      info.City,
		Status:    2,
		CardFront: info.CardFront,
		CardBack:  info.CardBack,
		DelTime:   0,
		Create:    int(time.Now().Unix()),
	}).Error
}
