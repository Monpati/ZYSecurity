package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type SSLService struct {
	Id                 int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	SSLId              int64  `gorm:"column:ssl_id" json:"ssl_id"`
	AdminInfo          string `gorm:"column:admin_info" json:"admin_info,omitempty"`
	COMStatus          int64  `gorm:"column:com_status" json:"com_status"`
	DNSHost            string `gorm:"column:dns_host" json:"dns_host,omitempty"`
	DNSType            string `gorm:"column:dns_type" json:"dns_type,omitempty"`
	DNSValue           string `gorm:"column:dns_value" json:"dns_value,omitempty"`
	DomainList         string `gorm:"column:domain_list" json:"domain_list,omitempty"`
	DomainNum          string `gorm:"column:domain_num" json:"domain_num,omitempty"`
	DomainType         string `gorm:"column:domain_type" json:"domain_type,omitempty"`
	FileName           string `gorm:"column:file_name" json:"file_name,omitempty"`
	FileValue          string `gorm:"column:file_value" json:"file_value,omitempty"`
	OrderId            string `gorm:"column:order_id" json:"order_id,omitempty"`
	OrderName          string `gorm:"column:order_name" json:"order_name,omitempty"`
	OrderStart         int64  `gorm:"column:order_start" json:"order_start,omitempty"`
	OrgInfo            string `gorm:"column:org_info" json:"org_info,omitempty"`
	PMethod            string `gorm:"column:p_method" json:"p_method,omitempty"`
	PTypeName          string `gorm:"column:p_type_name" json:"p_type_name,omitempty"`
	SetupServer        int64  `gorm:"column:setup_server" json:"setup_server,omitempty"`
	SSLName            string `gorm:"column:ssl_name" json:"ssl_name,omitempty"`
	SSLCode            string `gorm:"column:ssl_code" json:"ssl_code,omitempty"`
	SSLCsr             string `gorm:"column:ssl_csr" json:"ssl_csr,omitempty"`
	SSLKey             string `gorm:"column:ssl_key" json:"ssl_key,omitempty"`
	SSLPem             string `gorm:"column:ssl_pem" json:"ssl_pem"`
	SSLType            string `gorm:"column:ssl_type" json:"ssl_type,omitempty"`
	TechInfo           string `gorm:"column:tech_info" json:"tech_info,omitempty"`
	UUID               string `gorm:"column:uuid" json:"uuid,omitempty"`
	XufeiOrderid       string `gorm:"column:xufei_orderid" json:"xufei_orderid,omitempty"`
	YnProve            int64  `gorm:"column:yn_prove" json:"yn_prove,omitempty"`
	YnReplace          int64  `gorm:"column:yn_replace" json:"yn_replace,omitempty"`
	YnXufei            string `gorm:"column:yn_xufei" json:"yn_xufei,omitempty"`
	ZDomain            string `gorm:"column:z_domain" json:"z_domain,omitempty"`
	ZDomainList        string `gorm:"column:z_domain_list" json:"z_domain_list,omitempty"`
	KsMoney            int64  `gorm:"column:ks_money" json:"ks_money,omitempty"`
	EdTime             string `gorm:"column:ed_time" json:"ed_time,omitempty"`
	DateDiff           int64  `gorm:"column:date_diff" json:"date_doff,omitempty"`
	Img                string `gorm:"column:img" json:"img,omitempty"`
	ZyksMoney          int64  `gorm:"column:zyks_money" json:"zyks_money"`
	SSLTypeDetail      string `gorm:"column:ssl_type_detail" json:"ssl_type_detail"`
	SetUpServiceDetail string `gorm:"column:setup_service_detail" json:"setup_service_detail"`
	AdminName          string `gorm:"column:admin_name" json:"admin_name"`
	AdminTel           string `gorm:"column:admin_tel" json:"admin_tel"`
	AdminEmail         string `gorm:"column:admin_email" json:"admin_email"`
	AdminJob           string `gorm:"column:admin_job" json:"admin_job"`

	Create string `gorm:"column:create" json:"create"`
	CaNum  string `gorm:"ca_num" json:"ca_num"`
	UserId int64  `gorm:"column:user_id" json:"user_id"`
	Agent  string `gorm:"column:agent" json:"agent"`
}

func (SSLService) TableName() string {
	return "SSLService"
}

func CreateSSLService(db *gorm.DB, info *form.SSLServiceInfo) error {
	items := db.Table("SSLService")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&SSLService{
		Id:                 sf.Generate(),
		SSLId:              info.SSLId,
		AdminInfo:          info.AdminInfo,
		COMStatus:          info.COMStatus,
		DNSHost:            info.DNSHost,
		DNSType:            info.DNSType,
		DNSValue:           info.DNSValue,
		DomainList:         info.DomainList,
		DomainNum:          info.DomainNum,
		DomainType:         info.DomainType,
		FileName:           info.FileName,
		FileValue:          info.FileValue,
		OrderStart:         info.OrderStart,
		OrgInfo:            info.OrgInfo,
		PMethod:            info.PMethod,
		PTypeName:          info.PTypeName,
		SetupServer:        info.SetupServer,
		SSLCode:            info.SSLCode,
		SSLCsr:             info.SSLCsr,
		SSLKey:             info.SSLKey,
		SSLPem:             info.SSLPem,
		SSLType:            info.SSLType,
		TechInfo:           info.TechInfo,
		UUID:               info.UUID,
		XufeiOrderid:       info.XufeiOrderid,
		YnProve:            info.YnProve,
		YnReplace:          info.YnReplace,
		YnXufei:            info.YnXufei,
		ZDomain:            info.ZDomain,
		OrderId:            info.OrderId,
		OrderName:          info.OrderName,
		SSLName:            info.SSLName,
		ZDomainList:        info.ZDomainList,
		KsMoney:            info.KsMoney,
		EdTime:             info.EdTime,
		DateDiff:           info.DateDiff,
		Img:                info.Img,
		ZyksMoney:          3 * info.KsMoney,
		SSLTypeDetail:      info.SSLTypeDetail,
		SetUpServiceDetail: info.SetUpServiceDetail,
		AdminName:          info.AdminName,
		AdminTel:           info.AdminTel,
		AdminEmail:         info.AdminEmail,
		AdminJob:           info.AdminJob,

		Create: info.Create,
		CaNum:  info.CaNum,
		UserId: info.UserId,
		Agent:  info.Agent,
	}).Error
}

func (p *SSLService) GetOrdersByUser(db *gorm.DB, info *form.Filter) (*[]SSLService, int, error) {
	var orders []SSLService
	var total int

	if err := db.Model(&SSLService{}).
		Where("user_id = ?", p.UserId).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).
		Limit(info.Limit).
		Where("user_id = ?", p.UserId).
		Offset(info.Offset).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return &orders, total, nil

}

func (p *SSLService) GetOrdersByAgent(db *gorm.DB, info *form.Filter) (*[]SSLService, int, error) {
	var orders []SSLService
	var total int

	if err := db.Table(p.TableName()).Where("agent = ?", p.Agent).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).Limit(info.Limit).Where("agent = ?", p.Agent).Offset(info.Offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return &orders, total, nil
}

func (p *SSLService) GetByParams(db *gorm.DB, info *form.Filter) (*[]SSLService, int, error) {
	var services []SSLService
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&SSLService{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

		if info.Field == "source" {
			if err := db.Model(&SSLService{}).Where("`source` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`source` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

		if info.Field == "u_user_id" {
			if err := db.Model(&SSLService{}).Where("`u_user_id` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`u_user_id` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&services).Error; err != nil {
				return nil, 0, err
			}
			return &services, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		if err := db.Model(&SSLService{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Limit(info.Limit).Offset(info.Offset).Find(&services).Error; err != nil {
			return nil, 0, err
		}
		return &services, total, nil
	}
	return &services, total, nil
}

func (p *SSLService) GetSSLById(db *gorm.DB, id int64) error {
	return db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error
}
