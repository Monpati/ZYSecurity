package model

import (
	"Dexun/form"
	"Dexun/model/Dexun"
	"Dexun/utils"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Domain struct {
	Id                  int64       `gorm:"primary_key" gorm:"column:id" json:"id"`
	Domain              string      `gorm:"column:domain" json:"domain"`
	DomainUuid          string      `gorm:"column:domain_uuid" json:"domain_uuid"`
	DomainId            int64       `gorm:"column:domain_id" json:"domain_id"`
	OrderId             *int64      `gorm:"column:order_id" json:"order_id"`
	PrimaryDomain       string      `gorm:"column:primary_domain" json:"primary_domain"`
	DomainStatus        int64       `gorm:"column:domain_status" json:"domain_status"`
	DomainRecord        interface{} `gorm:"column:domain_record" json:"domain_record"`
	FourLayersConfig    JSON        `gorm:"column:four_layers_config" json:"four_layers_config"`
	CacheFileSizeLimit  int64       `gorm:"column:cache_file_size_limit" json:"cache_file_size_limit"`
	CacheTotalSizeLimit int64       `gorm:"column:cache_total_size_limit" json:"cache_total_size_limit"`
	CacheConfig         JSON        `gorm:"column:cache_config" json:"cache_config"`
	CacheActive         int64       `gorm:"column:cache_active" json:"cache_active"`
	WhiteNum            int64       `gorm:"column:white_num" json:"white_num"`
	UseFlow             int64       `gorm:"column:use_flow" json:"use_flow"`
	CreateTime          string      `gorm:"column:createtime" json:"createtime"`
	UpdateTime          string      `gorm:"column:updatetime" json:"updatetime"`
	AccessActive        string      `gorm:"column:access_active" json:"access_active"`
	Grouping            string      `gorm:"column:grouping" json:"grouping"`
	IsFiling            string      `gorm:"column:is_filing" json:"is_filing"`
	Status              int         `gorm:"column:status" json:"status"`
	UserId              int64       `gorm:"column:user_id" json:"user_id"`
	Cname               string      `gorm:"column:cname" json:"cname"`
	WafSwitch           bool        `gorm:"column:waf_switch" json:"waf_switch"`
	WafFile             bool        `gorm:"column:waf_file" json:"waf_file"`
	WafCode             bool        `gorm:"column:waf_code" json:"waf_code"`
	WafSession          bool        `gorm:"column:waf_session" json:"waf_session"`
	WafShellShock       bool        `gorm:"column:waf_shellshock" json:"waf_shellshock"`
	WafZombie           bool        `gorm:"column:waf_zombie" json:"waf_zombie"`
	WafMetadata         bool        `gorm:"column:waf_metadata" json:"waf_metadata"`
	WafSql              bool        `gorm:"column:waf_sql" json:"waf_sql"`
	WafProxy            bool        `gorm:"column:waf_proxy" json:"waf_proxy"`
	WafXss              bool        `gorm:"column:waf_xss" json:"waf_xss"`
	DDoSDDId            *int64      `gorm:"column:ddosdd_id" json:"ddosdd_id"`
}

type ConfigList struct {
	Id                  int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	DomainId            int64  `gorm:"column:domain_id" json:"domain_id"`
	LoadBalancing       string `gorm:"column:load_balancing" json:"load_balancing"`
	OverloadRedirectUrl string `gorm:"column:overload_redirect_url" json:"overload_redirect_url"`
	OverloadStatusCode  string `gorm:"column:overload_status_code" json:"overload_status_code"`
	OverloadType        string `gorm:"column:overload_type" json:"overload_type"`
	Port                string `gorm:"column:port" json:"port"`
	Protocol            string `gorm:"column:protocol" json:"protocol"`
	Redirect            string `gorm:"column:redirect" json:"redirect"`
	Server              string `gorm:"column:server" json:"server"`
	UriForward          string `gorm:"column:uri_forward" json:"uri_forward"`
}

type DomainConfig struct {
	Id           int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	ConfigListId int64  `gorm:"column:config_list_id" json:"config_list_id"`
	Address      string `gorm:"column:address" json:"address"`
	Concurrent   string `gorm:"column:concurrent" json:"concurrent"`
	Port         string `gorm:"column:port" json:"port"`
	Protocol     string `gorm:"column:protocol" json:"protocol"`
	Sni          string `gorm:"column:sni" json:"sni"`
	Weight       string `gorm:"column:weight" json:"weight"`
}

func (Domain) TableName() string {
	return "ScdnDomains"
}

func CreateDomains(db *gorm.DB, info *form.DomainInfo) error {
	items := db.Table("ScdnDomains")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&Domain{
		Id:       sf.Generate(),
		Domain:   info.Domain,
		OrderId:  info.OrderId,
		UserId:   info.UserId,
		DDoSDDId: info.DDoSDDId,
	}).Error
}

func (p *Domain) GetIdByOIAndDomain(db *gorm.DB, order_id int64, domain, service string) int64 {
	if service == "scdn" {
		if err := db.Table(p.TableName()).Where("domain = ? and order_id = ?", domain, order_id).Find(&p).Error; err != nil {
			return 0
		}
	}

	if service == "ddos" {
		if err := db.Table(p.TableName()).Where("domain = ? and ddosdd_id = ?", domain, order_id).Find(&p).Error; err != nil {
			return 0
		}
	}

	return p.Id
}

func (p *Domain) GetByParams(db *gorm.DB, info *form.DomainFilterForm) (*[]Domain, int, error) {
	var domains []Domain
	var total int
	query := db.Model(&Domain{})

	if info.Protocol != "" {
		query = query.Where("`protocol` LIKE ?", "%"+info.Protocol+"%")
	}
	if info.Port != "" {
		query = query.Where("`port` LIKE ?", "%"+info.Port+"%")
	}
	if info.Server != "" {
		query = query.Where("`server` LIKE ?", "%"+info.Server+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&domains).Error; err != nil {
		return nil, 0, err
	}
	return &domains, total, nil
}

func (p *Domain) GetDomainsByUser(db *gorm.DB, info *form.Filter, user_id, order_id int64, service string) (*[]Domain, int, error) {
	var domains []Domain
	var total int

	if service == "scdn" {
		if info.Field != "" {
			if info.Field == "protocol" {
				if err := db.Model(&Domain{}).Where("`protocol` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("order_id = ?", order_id).Count(&total).Error; err != nil {
					return nil, 0, err
				}
				if err := db.Limit(info.Limit).Where("`protocol` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("order_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
					return nil, 0, err
				}
				return &domains, total, nil
			}

			if info.Field == "domain" {
				if err := db.Model(&Domain{}).Where("`domain` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("order_id = ?", order_id).Count(&total).Error; err != nil {
					return nil, 0, err
				}
				if err := db.Limit(info.Limit).Where("`domain` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("order_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
					return nil, 0, err
				}
				return &domains, total, nil
			}

			if info.Field == "sni" {
				if err := db.Model(&Domain{}).Where("`sni` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("order_id = ?", order_id).Count(&total).Error; err != nil {
					return nil, 0, err
				}
				if err := db.Limit(info.Limit).Where("`sni` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("order_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
					return nil, 0, err
				}
				return &domains, total, nil
			}

		}

		if info.Field == "" {
			//获取总记录数
			if err := db.Model(&Domain{}).Where("user_id = ?", user_id).Where("order_id = ?", order_id).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			//分页查询
			if err := db.Limit(info.Limit).Where("user_id = ?", user_id).Where("order_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
				return nil, 0, err
			}
			return &domains, total, nil
		}
	}

	if service == "ddos" {
		if info.Field != "" {
			if info.Field == "protocol" {
				if err := db.Model(&Domain{}).Where("`protocol` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Count(&total).Error; err != nil {
					return nil, 0, err
				}
				if err := db.Limit(info.Limit).Where("`protocol` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
					return nil, 0, err
				}
				return &domains, total, nil
			}

			if info.Field == "domain" {
				if err := db.Model(&Domain{}).Where("`domain` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Count(&total).Error; err != nil {
					return nil, 0, err
				}
				if err := db.Limit(info.Limit).Where("`domain` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
					return nil, 0, err
				}
				return &domains, total, nil
			}

			if info.Field == "sni" {
				if err := db.Model(&Domain{}).Where("`sni` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Count(&total).Error; err != nil {
					return nil, 0, err
				}
				if err := db.Limit(info.Limit).Where("`sni` LIKE ?", "%"+info.Value+"%").Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
					return nil, 0, err
				}
				return &domains, total, nil
			}

		}

		if info.Field == "" {
			//获取总记录数
			if err := db.Model(&Domain{}).Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			//分页查询
			if err := db.Limit(info.Limit).Where("user_id = ?", user_id).Where("ddosdd_id = ?", order_id).Offset(info.Offset).Find(&domains).Error; err != nil {
				return nil, 0, err
			}
			return &domains, total, nil
		}
	}

	return &domains, total, nil
}

func (p *Domain) GetDomainsByOrder(db *gorm.DB, info *form.Filter) (*[]Domain, int, error) {
	var domains []Domain
	var total int

	if err := db.Table(p.TableName()).
		Where("user_id", p.OrderId).
		Find(&domains).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).
		Limit(info.Limit).
		Where("user_id", p.OrderId).
		Offset(info.Offset).
		Find(&domains).Error; err != nil {
		return nil, 0, err
	}
	return &domains, total, nil

}

func (p *Domain) GetDomainById(db *gorm.DB, info *form.Filter, user_id int64) (*[]Domain, int, error) {
	var total int
	var domains []Domain

	if err := db.Table(p.TableName()).Where("user_id = ?", user_id).Find(&domains).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Table(p.TableName()).
		Limit(info.Limit).
		Where("user_id = ?", user_id).
		Offset(info.Offset).
		Find(&domains).Error; err != nil {
		return nil, 0, err
	}
	return &domains, total, nil
}

func (p *Domain) GetIdByDomain(db *gorm.DB, id int64) int64 {
	if err := db.Table("ScdnDomains").Where("id = ?", id).Find(&p); err != nil {
		return 0
	}
	return p.Id
}

func (p *Domain) GetOrderIdById(db *gorm.DB, id int64) int64 {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error; err != nil {
		return 0
	}
	return *p.OrderId
}

func (p *Domain) GetUuidById(db *gorm.DB, id int64) string {
	if err := db.Table(p.TableName()).Where("id = ?", id).Find(&p).Error; err != nil {
		return ""
	} else {
		return p.DomainUuid
	}
}

func (p *Domain) UpdateDomainStatus(db *gorm.DB, domain_id int64, status int) error {
	if err := db.Table(p.TableName()).
		Where("id = ?", domain_id).
		UpdateColumn("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (p *Domain) UpdateDomain(db *gorm.DB, info *Dexun.DeXunBody, domain, service string, order_id int64) error {
	cache_config, _ := json.Marshal(info.CacheConfig)
	four_layers_config, _ := json.Marshal(info.FourLayersConfig)
	if service == "scdn" {
		if err := db.Table(p.TableName()).
			Where("domain = ? and order_id = ?", domain, order_id).
			Updates(&Domain{
				DomainUuid:          info.DomainUUID,
				PrimaryDomain:       info.PrimaryDomain,
				DomainStatus:        info.DomainStatus,
				DomainRecord:        info.DomainRecord,
				FourLayersConfig:    four_layers_config,
				CacheFileSizeLimit:  info.DomainConfig.CacheFileSizeLimit,
				CacheTotalSizeLimit: info.DomainConfig.CacheTotalSizeLimit,
				CacheConfig:         cache_config,
				CacheActive:         info.CacheActive,
				WhiteNum:            info.WhiteNum,
				UseFlow:             info.UseFlow,
				CreateTime:          info.Createtime,
				UpdateTime:          info.Updatetime,
				AccessActive:        info.AccessActive,
				Grouping:            info.Grouping,
				IsFiling:            info.IsFilings,
				Status:              1,
				Cname:               info.DomainUUID + "." + info.PrimaryDomain,
			}).Error; err != nil {
			return err
		}
	}
	if service == "ddos" {
		if err := db.Table(p.TableName()).
			Where("domain = ? and ddosdd_id = ?", domain, order_id).
			Updates(&Domain{
				DomainUuid:          info.DomainUUID,
				PrimaryDomain:       info.PrimaryDomain,
				DomainStatus:        info.DomainStatus,
				DomainRecord:        info.DomainRecord,
				FourLayersConfig:    four_layers_config,
				CacheFileSizeLimit:  info.DomainConfig.CacheFileSizeLimit,
				CacheTotalSizeLimit: info.DomainConfig.CacheTotalSizeLimit,
				CacheConfig:         cache_config,
				CacheActive:         info.CacheActive,
				WhiteNum:            info.WhiteNum,
				UseFlow:             info.UseFlow,
				CreateTime:          info.Createtime,
				UpdateTime:          info.Updatetime,
				AccessActive:        info.AccessActive,
				Grouping:            info.Grouping,
				IsFiling:            info.IsFilings,
				Status:              1,
				Cname:               info.DomainUUID + "." + info.PrimaryDomain,
			}).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateCL(db *gorm.DB, info *form.ConfigListInfo, domain_id int64) (error, int64) {
	items := db.Table("ConfigList")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	tmp := sf.Generate()
	return items.Create(&ConfigList{
		Id:                  tmp,
		DomainId:            domain_id,
		LoadBalancing:       info.LoadBalancing,
		OverloadRedirectUrl: info.OverloadRedirectUrl,
		OverloadStatusCode:  info.OverloadStatusCode,
		OverloadType:        info.OverloadType,
		Port:                info.Port,
		Protocol:            info.Protocol,
		Redirect:            info.Redirect,
		Server:              info.Server,
		UriForward:          info.UriForward,
	}).Error, tmp
}

func (p *ConfigList) GetCLParams(db *gorm.DB, info *form.ConfigListInfo, domain_id int64) (*[]ConfigList, int, error) {
	var account []ConfigList
	var total int

	query := db.Table("ConfigList")

	if info.Protocol != "" {
		query = query.Where("domain_id = ?", domain_id).Where("`protocol` LIKE ?", "%"+info.Protocol+"%")
	}
	if info.LoadBalancing != "" {
		query = query.Where("domain_id = ?", domain_id).Where("`load_balancing` LIKE ?", "%"+info.LoadBalancing+"%")
	}
	if info.Server != "" {
		query = query.Where("domain_id = ?", domain_id).Where("`cert_type` LIKE ?", "%"+info.Server+"%")
	}

	if err := query.Where("domain_id = ?", domain_id).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Where("domain_id = ?", domain_id).Limit(info.Limit).Offset(info.Offset).Find(&account).Error; err != nil {
		return nil, 0, err
	}

	return &account, total, nil
}

func (p *ConfigList) GetIdByDomainId(db *gorm.DB, domain_id int64) int64 {
	if err := db.Table("ConfigList").Where("domain_id = ?", domain_id).Find(&p); err != nil {
		return 0
	}
	return p.Id
}

func CreateSA(db *gorm.DB, info *form.SourceAddressInfo, cl_id int64) error {
	items := db.Table("SourceAddress")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DomainConfig{
		Id:           sf.Generate(),
		ConfigListId: cl_id,
		Address:      info.Address,
		Concurrent:   info.Concurrent,
		Port:         info.Port,
		Protocol:     info.Protocol,
		Sni:          info.Sni,
		Weight:       info.Weight,
	}).Error
}

func (p *Domain) UpdateTotalWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_switch", info.Active).Error
}

func (p *Domain) UpdateFileWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_file", info.Active).Error
}

func (p *Domain) UpdateCodeWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_code", info.Active).Error
}

func (p *Domain) UpdateSessionWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_session", info.Active).Error
}

func (p *Domain) UpdateShellshockWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_shellshock", info.Active).Error
}

func (p *Domain) UpdateScriptWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_zombie", info.Active).Error
}

func (p *Domain) UpdateMetaWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_metadata", info.Active).Error
}

func (p *Domain) UpdateSqlWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_sql", info.Active).Error
}

func (p *Domain) UpdateProxyWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_proxy", info.Active).Error
}

func (p *Domain) UpdateXssWaf(db *gorm.DB, info *form.WafInfo) error {
	return db.Table("ScdnDomains").Where("id = ?", info.DomainId).UpdateColumn("waf_xss", info.Active).Error
}
