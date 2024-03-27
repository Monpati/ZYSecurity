package model

import (
	"Dexun/form"
	"Dexun/utils"
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"mime/multipart"
	"net/http"
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

func (p *PersonalCert) GetStatusById(db *gorm.DB, id int64) error {
	return db.Table(p.TableName()).Where("user_id = ?", id).Find(p).Error
}

func (p *Account) GetIdByName(db *gorm.DB, username string) error {
	return db.Table("User").Where("username = ?", username).Find(p).Error
}

func (p *PersonalCert) UpdateStatus(db *gorm.DB, userid int64, status int) error {
	return db.Table(p.TableName()).
		Where("user_id = ?", userid).
		UpdateColumn("status", status).Error
}

func (p *PersonalCert) GetByParams(db *gorm.DB, info *form.PersonFilter) (*[]PersonalCert, int, error) {
	var cert []PersonalCert
	var total int

	if info.Field != "" {
		if info.Field == "username" {
			if err := db.Model(&PersonalCert{}).Where("`real_name` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`real_name` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&cert).Error; err != nil {
				return nil, 0, err
			}
			return &cert, total, nil
		}

		if info.Field == "tel_num" {
			if err := db.Model(&PersonalCert{}).Where("`card_id` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`card_id` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&cert).Error; err != nil {
				return nil, 0, err
			}
			return &cert, total, nil
		}

		if info.Field == "email" {
			if err := db.Model(&PersonalCert{}).Where("`sex` LIKE ?", "%"+info.Value+"%").Where("`del_time` = ?", 0).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`sex` LIKE ?", "%"+info.Value+"%").Where("del_time = ?", 0).Offset(info.Offset).Find(&cert).Error; err != nil {
				return nil, 0, err
			}
			return &cert, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&PersonalCert{}).Where("del_time = ?", 0).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Where("del_time = ?", 0).Offset(info.Offset).Find(&cert).Error; err != nil {
			return nil, 0, err
		}
		return &cert, total, nil
	}
	return &cert, total, nil
}

// sex、birthday、city
func (p *PersonalCert) GetCertDetails(name, cardid string) error {
	request, err := json.Marshal(map[string]string{
		"gname": name,
		"gnum":  cardid,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://apiagent.dexunyun.com/agent/authentication.real/nameauthgr", bytes.NewBuffer(request))
	if err != nil {
		return err
	}
	//需要改成token
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var responseData form.CertResponseData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return err
	}
	p.Sex = responseData.Data.SmrzSex
	p.Birthday = responseData.Data.SmrzBirthday
	p.City = responseData.Data.SmrzCity

	return nil
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
