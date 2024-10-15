package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type CC struct {
	Id                  int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	CcId                int64      `gorm:"column:cc_id" json:"cc_id"`
	DomainId            int64      `gorm:"column:domain_id" json:"domain_id"`
	OrderId             int64      `gorm:"column:order_id" json:"order_id"`
	DomainUuid          string     `gorm:"column:domain_uuid" json:"domain_uuid"`
	CsActive            string     `gorm:"column:cs_active" json:"cs_active"`
	Policy              string     `gorm:"column:policy" json:"policy"`
	Url                 string     `gorm:"column:url" json:"url"`
	Rate                string     `gorm:"column:rate" json:"rate"`
	WaitSeconds         string     `gorm:"column:waitseconds" json:"waitseconds"`
	BlockMinutes        string     `gorm:"column:blockminutes" json:"blockminutes"`
	RedirectLocation    string     `gorm:"column:redirectlocation" json:"redirectlocation"`
	GlobalConcurrent    string     `gorm:"column:global_concurrent" json:"global_concurrent"`
	WaitPolicyMinutes   string     `gorm:"column:waitpolicyminutes" json:"waitpolicyminutes"`
	RedirectWaitSeconds string     `gorm:"column:redirectwaitseconds" json:"redirectwaitseconds"`
	Count               string     `gorm:"column:count" json:"count"`
	BlockTime           string     `gorm:"column:block_time" json:"block_time"`
	BlockActive         string     `gorm:"column:block_active" json:"block_active"`
	RrActive            string     `gorm:"column:rr_active" json:"rr_active"`
	RrRate              string     `gorm:"column:rr_rate" json:"rr_rate"`
	UrUrl               string     `gorm:"column:ur_url" json:"ur_url"`
	UrRate              string     `gorm:"column:ur_rate" json:"ur_rate"`
	CookieName          string     `gorm:"column:cookieName" json:"cookieName"`
	ExcludeExt          string     `gorm:"column:excludeExt" json:"excludeExt"`
	Concurrency         string     `gorm:"column:concurrency" json:"concurrency"`
	RBlockMinutes       string     `gorm:"column:r_blockMinutes" json:"r-blockMinutes"`
	WhiteMinutes        string     `gorm:"column:whiteMinutes" json:"whiteMinutes"`
	ChallengeLimit      string     `gorm:"column:challengeLimit" json:"challengeLimit"`
	ProtectMinutes      string     `gorm:"column:protectMinutes" json:"protectMinutes"`
	ChallengeMethods    utils.JSON `gorm:"column:challengeMethods" json:"challengeMethods"`
	ChallengePolicy     string     `gorm:"column:challengePolicy" json:"challengePolicy"`
	Active              string     `gorm:"column:active" json:"active"`
	UseDefault          string     `gorm:"column:use_default" json:"use_default"`
}

func (CC) TableName() string {
	return "ScdnCC"
}

func CreateCC(db *gorm.DB, info *form.CC) error {
	items := db.Table("ScdnCC")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&CC{
		Id:                  sf.Generate(),
		DomainId:            info.DomainId,
		OrderId:             info.OrderId,
		DomainUuid:          info.DomainUUID,
		CsActive:            info.Config.Site.Active,
		Policy:              info.Config.Site.Policy,
		WaitSeconds:         info.Config.Site.Waitseconds,
		BlockMinutes:        info.Config.Site.Blockminutes,
		RedirectLocation:    info.Config.Site.Redirectlocation,
		GlobalConcurrent:    info.Config.Site.GlobalConcurrent,
		WaitPolicyMinutes:   info.Config.Site.Waitpolicyminutes,
		RedirectWaitSeconds: info.Config.Site.Redirectwaitseconds,
		Count:               info.Config.BlockConfig.Count,
		BlockTime:           info.Config.BlockConfig.BlockTime,
		BlockActive:         info.Config.BlockConfig.BlockActive,
		RrActive:            info.Config.ResuestRate.Active,
		RrRate:              info.Config.ResuestRate.Rate,
		CookieName:          info.Config.ResuestRate.CookieName,
		ExcludeExt:          info.Config.ResuestRate.ExcludeEXT,
		Concurrency:         info.Config.ResuestRate.Concurrency,
		RBlockMinutes:       info.Config.ResuestRate.BlockMinutes,
		WhiteMinutes:        info.Config.ResuestRate.WhiteMinutes,
		ChallengeLimit:      info.Config.ResuestRate.ChallengeLimit,
		ProtectMinutes:      info.Config.ResuestRate.ProtectMinutes,
		ChallengeMethods:    info.Config.ResuestRate.ChallengeMethod,
		ChallengePolicy:     info.Config.ResuestRate.ChallengePolicy,
		Active:              info.Active,
		UseDefault:          info.UseDefault,
	}).Error
}

func (p *CC) GetCCLists(db *gorm.DB, info *form.CcFilterForm) (*[]CC, int, error) {
	var lists []CC
	var total int
	query := db.Model(&Account{})

	if info.Policy != "" {
		query = query.Where("`policy` LIKE ?", "%"+info.Policy+"%")
	}
	if info.Url != "" {
		query = query.Where("`url` LIKE ?", "%"+info.Url+"%")
	}
	if info.CsActive != "" {
		query = query.Where("`cs_active` LIKE ?", "%"+info.CsActive+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
		return nil, 0, err
	}

	return &lists, total, nil
}

func (p *CC) UpdateCC(db *gorm.DB, id int64, info *form.CCInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&CC{
			CsActive:            info.CsActive,
			Policy:              info.Policy,
			Url:                 info.Url,
			Rate:                info.Rate,
			WaitSeconds:         info.WaitSeconds,
			BlockMinutes:        info.BlockMinutes,
			RedirectLocation:    info.RedirectLocation,
			GlobalConcurrent:    info.GlobalConcurrent,
			WaitPolicyMinutes:   info.WaitPolicyMinutes,
			RedirectWaitSeconds: info.RedirectWaitSeconds,
			Count:               info.Count,
			BlockTime:           info.BlockTime,
			BlockActive:         info.BlockActive,
			RrActive:            info.RrActive,
			RrRate:              info.RrRate,
			UrUrl:               info.UrUrl,
			UrRate:              info.UrRate,
			CookieName:          info.CookieName,
			ExcludeExt:          info.ExcludeExt,
			Concurrency:         info.Concurrency,
			RBlockMinutes:       info.RBlockMinutes,
			WhiteMinutes:        info.WhiteMinutes,
			ChallengeLimit:      info.ChallengeLimit,
			ProtectMinutes:      info.ProtectMinutes,
			ChallengeMethods:    info.ChallengeMethods,
			ChallengePolicy:     info.ChallengePolicy,
			Active:              info.Active,
			UseDefault:          info.UseDefault,
		}).Error
}
