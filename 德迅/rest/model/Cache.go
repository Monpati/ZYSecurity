package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type Cache struct {
	Id              int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	DxCacheId       int64  `gorm:"column:dxcache_id" json:"dxcache_id"`
	CacheId         int64  `gorm:"column:cache_id" json:"cache_id"`
	DdUuid          string `gorm:"column:dd_uuid" json:"dd_uuid"`
	OrderId         int64  `gorm:"column:order_id" json:"order_id"`
	CacheUuid       string `gorm:"column:cache_uuid" json:"cache_uuid"`
	DomainUuid      string `gorm:"column:domain_uuid" json:"domain_uuid"`
	DomainId        int64  `gorm:"column:domain_id" json:"domain_id"`
	Active          int64  `gorm:"column:active" json:"active"`
	UrlMode         string `gorm:"column:urlmode" json:"urlmode"`
	CacheMode       string `gorm:"column:cachemode" json:"cachemode"`
	CachePath       string `gorm:"column:cachepath" json:"cachepath"`
	CacheExtensions string `gorm:"column:cacheextensions" json:"cacheextensions"`
	CacheReg        string `gorm:"column:cachereg" json:"cachereg"`
	TimeOut         string `gorm:"column:timeout" json:"timeout"`
	Weight          string `gorm:"column:weight" json:"weight"`
	CreateTime      string `gorm:"column:createtime" json:"createtime"`
	UpdateTime      string `gorm:"column:updatetime" json:"updatetime"`
	Status          int    `gorm:"column:status" json:"status"`
}

func (Cache) TableName() string {
	return "ScdnCache"
}

func CreateCache(db *gorm.DB, info *form.CacheInfo, active int64) (error, int64) {
	items := db.Table("ScdnCache")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	tmp := sf.Generate()
	return items.Create(&Cache{
		Id:              tmp,
		DxCacheId:       info.DxCacheId,
		CacheId:         info.CacheId,
		DdUuid:          info.DdUuid,
		OrderId:         info.OrderId,
		CacheUuid:       info.CacheUuid,
		DomainUuid:      info.DomainUuid,
		DomainId:        info.DomainId,
		Active:          active,
		UrlMode:         info.UrlMode,
		CacheMode:       info.CacheMode,
		CachePath:       info.CachePath,
		CacheExtensions: info.CacheExtensions,
		CacheReg:        info.CacheReg,
		TimeOut:         info.TimeOut,
		Weight:          info.Weight,
		CreateTime:      info.CreateTime,
		UpdateTime:      info.UpdateTime,
		Status:          1,
	}).Error, tmp
}

func (p *Cache) GetCacheLists(db *gorm.DB, info *form.CacheFilterForm) (*[]Cache, int, error) {
	var lists []Cache
	var total int
	query := db.Model(&Account{})

	if info.UrlMode != "" {
		query = query.Where("`urlmode` LIKE ?", "%"+info.UrlMode+"%")
	}
	if info.CacheMode != "" {
		query = query.Where("`cachemode` LIKE ?", "%"+info.CacheMode+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Limit(info.Limit).Offset(info.Offset).Find(&lists).Error; err != nil {
		return nil, 0, err
	}

	return &lists, total, nil
}

func (p *Cache) UpdateCache(db *gorm.DB, id int64, info *form.CacheInfo, active int64) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&Cache{
			CacheUuid:       info.CacheUuid,
			Active:          active,
			UrlMode:         info.UrlMode,
			CacheMode:       info.CacheMode,
			CachePath:       info.CachePath,
			CacheExtensions: info.CacheExtensions,
			CacheReg:        info.CacheReg,
			TimeOut:         info.TimeOut,
			Weight:          info.Weight,
			CreateTime:      info.CreateTime,
			UpdateTime:      info.UpdateTime,
			Status:          1,
		}).Error
}
