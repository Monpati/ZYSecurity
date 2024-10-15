package model

import (
	"Dexun/form"
	"Dexun/model/Dexun"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type Certification struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	Ddc_id     int64  `gorm:"column:ddc_id" json:"ddc_Id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	Status     int64  `gorm:"column:status" json:"status"`
	SslAlways  int64  `gorm:"column:ssl_always" json:"ssl_always"`
	Hsts       int64  `gorm:"column:hsts" json:"hsts"`
	CertName   string `gorm:"column:cert_name" json:"cert_name"`
	Cert       string `gorm:"column:cert" json:"cert"`
	Key        string `gorm:"column:key" json:"key"`
	Desc       string `gorm:"column:desc" json:"desc"`
	Createtime string `gorm:"column:createtime" json:"createtime"`
	UpdateTime string `gorm:"column:updatetime" json:"updatetime"`
}

type DexunCert struct {
	Id         int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	Dc_id      int64  `gorm:"column:dc_id" json:"dc_Id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	Status     int64  `gorm:"column:status" json:"status"`
	SslAlways  int64  `gorm:"column:ssl_always" json:"ssl_always"`
	Hsts       int64  `gorm:"column:hsts" json:"hsts"`
	CertName   string `gorm:"column:cert_name" json:"cert_name"`
	Cert       string `gorm:"column:cert" json:"cert"`
	Key        string `gorm:"column:key" json:"key"`
	Desc       string `gorm:"column:desc" json:"desc"`
	Createtime string `gorm:"column:createtime" json:"createtime"`
	UpdateTime string `gorm:"column:updatetime" json:"updatetime"`
}

func (Certification) TableName() string {
	return "ScdnDomainCert"
}

func CreateCertification(db *gorm.DB, info *form.CertificationInfo) (error, int64) {
	items := db.Table("ScdnDomainCert")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	tmp := sf.Generate()
	return items.Create(&Certification{
		Id:         tmp,
		Ddc_id:     info.Ddc_id,
		OrderId:    info.OrderId,
		DomainId:   info.DomainId,
		DomainUuid: info.DomainUuid,
		Status:     info.Status,
		SslAlways:  info.SslAlways,
		Hsts:       info.Hsts,
		CertName:   info.CertName,
		Cert:       info.Cert,
		Key:        info.Key,
		Desc:       info.Desc,
		Createtime: info.Createtime,
		UpdateTime: info.UpdateTime,
	}).Error, tmp
}

func CreateDexunCert(db *gorm.DB, info Dexun.DeXunBody) error {
	items := db.Table("DexunDomainCert")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DexunCert{
		Id:         sf.Generate(),
		DomainUuid: info.DomainUUID,
		Status:     info.Status,
		SslAlways:  info.SSLAlways,
		Hsts:       info.Hsts,
		CertName:   info.CERTName,
		Cert:       info.CERT,
		Key:        info.Key,
		Desc:       info.Desc,
		Createtime: info.Createtime,
		UpdateTime: info.Updatetime,
	}).Error
}

func (p *Certification) GetCertificationLists(db *gorm.DB, info *form.CertificationFilterForm) (*[]Certification, int, error) {
	var lists []Certification
	var total int
	query := db.Model(&Account{})

	if info.CertName != "" {
		query = query.Where("`cert_name` LIKE ?", "%"+info.CertName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
		return nil, 0, err
	}
	return &lists, total, nil
}

func (p *Certification) UpdateCertification(db *gorm.DB, id int64, info *form.CertificationInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&Certification{
			Desc:       info.Desc,
			Createtime: info.Createtime,
			UpdateTime: info.UpdateTime,
		}).Error
}
