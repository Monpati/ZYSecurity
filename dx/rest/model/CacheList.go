package model

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type CacheList struct {
	Id              int64  `gorm:"primary_key" gorm:"column:id" json:"id"`
	CmId            int64  `gorm:"column:cm_id" json:"cm_id"`
	DdId            int64  `gorm:"column:dd_id" json:"dd_id"`
	DdType          int64  `gorm:"column:dd_type" json:"dd_type"`
	CacheName       string `gorm:"column:cache_name" json:"cache_name"`
	Active          string `gorm:"column:active" json:"active"`
	UrlMode         string `gorm:"column:urlmode" json:"urlmode"`
	CacheMode       string `gorm:"column:cachemode" json:"cachemode"`
	CachePath       string `gorm:"column:cachepath" json:"cachepath"`
	Cacheextensions string `gorm:"column:cacheextensions" json:"cacheextensions"`
	CacheReg        string `gorm:"column:cachereg" json:"cachereg"`
	TimeOut         int64  `gorm:"column:timeout" json:"timeout"`
	Weight          int64  `gorm:"column:weight" json:"weight"`
	Status          int    `gorm:"column:status" json:"status"`
}

func (CacheList) TableName() string {
	return "CacheList"
}

func CreateCacheModel(db *gorm.DB, info *form.CacheListInfo) error {
	items := db.Table("ScdnCacheModel")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&CacheList{
		Id:              sf.Generate(),
		CmId:            info.CmId,
		DdId:            info.DdId,
		DdType:          info.DdType,
		CacheName:       info.CacheName,
		Active:          info.Active,
		UrlMode:         info.UrlMode,
		CacheMode:       info.CacheMode,
		CachePath:       info.CachePath,
		Cacheextensions: info.Cacheextensions,
		CacheReg:        info.CacheReg,
		TimeOut:         info.TimeOut,
		Weight:          info.Weight,
		Status:          1,
	}).Error
}

func (p *CacheList) GetCacheModelLists(db *gorm.DB, info *form.CacheListFilterForm) (*[]CacheList, int, error) {
	var lists []CacheList
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

func (p *CacheList) UpdateCacheModel(db *gorm.DB, id int64, info *form.CacheListInfo) error {
	return db.Table(p.TableName()).
		Where("id = ?", id).
		Updates(&CacheList{
			Active:          info.Active,
			UrlMode:         info.UrlMode,
			CacheMode:       info.CacheMode,
			CachePath:       info.CachePath,
			Cacheextensions: info.Cacheextensions,
			Weight:          info.Weight,
			Status:          info.Status,
		}).Error
}
