package Dexun

import (
	"Dexun/form"
	"Dexun/utils"
	"github.com/jinzhu/gorm"
)

type DexunCloudEyeList struct {
	ID          int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	UUID        string     `gorm:"column:uuid" json:"uuid"`
	TcName      string     `gorm:"column:tc_name" json:"tc_name"`
	KsMoney     int64      `gorm:"column:ks_money" json:"ks_money"`
	Content     string     `gorm:"column:content" json:"content"`
	TaskType    utils.JSON `gorm:"column:task_type" json:"task_type"`
	MonitorType utils.JSON `gorm:"column:monitor_type" json:"monitor_type"`
	StartCount  int64      `gorm:"column:start_count" json:"start_count"`
	OrderMoney  int64      `gorm:"column:order_money" json:"order_money"`
	ZkMoney     int64      `gorm:"column:zk_money" json:"zk_money"`
}

type CloudEyeList struct {
	ID          int64      `gorm:"primary_key" gorm:"column:id" json:"id"`
	CEId        int64      `gorm:"column:ce_id" json:"ce_id"`
	UUID        string     `gorm:"column:uuid" json:"uuid"`
	TcName      string     `gorm:"column:tc_name" json:"tc_name"`
	KsMoney     int64      `gorm:"column:ks_money" json:"ks_money"`
	Content     string     `gorm:"column:content" json:"content"`
	TaskType    utils.JSON `gorm:"column:task_type" json:"task_type"`
	MonitorType utils.JSON `gorm:"column:monitor_type" json:"monitor_type"`
	StartCount  int64      `gorm:"column:start_count" json:"start_count"`
	OrderMoney  int64      `gorm:"column:order_money" json:"order_money"`
	ZkMoney     int64      `gorm:"column:zk_money" json:"zk_money"`
	SellStatus  int        `gorm:"column:sell_status" json:"sell_status"`
	ZyksMoney   int64      `gorm:"column:zyks_money" json:"zyks_money"`
	ZyzkMoney   int64      `gorm:"column:zyzk_money" json:"zyzk_money"`
}

func (DexunCloudEyeList) TableName() string {
	return "DexunCloudEyeList"
}

func CreateDxCloudEye(db *gorm.DB, info *Datum) error {
	items := db.Table("DexunCloudEyeList")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	return items.Create(&DexunCloudEyeList{
		ID:          sf.Generate(),
		UUID:        info.UUID,
		TcName:      info.TcName,
		KsMoney:     info.KsMoney,
		Content:     info.Content,
		TaskType:    info.TaskType,
		MonitorType: info.MonitorType,
		StartCount:  info.StartCount,
		OrderMoney:  info.OrderMoney,
		ZkMoney:     info.ZkMoney,
	}).Error
}

func (p *DexunCloudEyeList) UpdateCloudEyeByUUID(db *gorm.DB, uuid string, info *Datum) error {
	return db.Table(p.TableName()).
		Where("uuid = ?", uuid).
		Updates(&DexunCloudEyeList{
			KsMoney:     info.KsMoney,
			Content:     info.Content,
			TaskType:    info.TaskType,
			MonitorType: info.MonitorType,
			StartCount:  info.StartCount,
			OrderMoney:  info.OrderMoney,
			ZkMoney:     info.ZkMoney,
		}).Error
}

func CopyDxCloudEyeList(db *gorm.DB) error {
	items := db.Table("CloudEyeList")
	sf, _ := utils.NewSnowflake(utils.GenerateRand())
	var combos []DexunCloudEyeList
	var service CloudEyeList
	if err := db.Table("DexunCloudEyeList").Find(&combos).Error; err != nil {
		return err
	}

	for _, combo := range combos {
		db.Table("CloudEyeList").Where("ce_id = ?", combo.ID).Find(&service)
		if service.ID == 0 {
			items.Create(&CloudEyeList{
				ID:          sf.Generate(),
				CEId:        combo.ID,
				UUID:        combo.UUID,
				TcName:      combo.TcName,
				KsMoney:     combo.KsMoney,
				Content:     combo.Content,
				TaskType:    combo.TaskType,
				MonitorType: combo.MonitorType,
				StartCount:  combo.StartCount,
				OrderMoney:  combo.OrderMoney,
				ZkMoney:     combo.ZkMoney,
				SellStatus:  1,
				ZyksMoney:   3 * combo.KsMoney,
				ZyzkMoney:   3 * combo.ZkMoney,
			})
		} else {
			items.Where("ssl_id = ?", combo.ID).Updates(&CloudEyeList{
				KsMoney:     combo.KsMoney,
				Content:     combo.Content,
				TaskType:    combo.TaskType,
				MonitorType: combo.MonitorType,
				StartCount:  combo.StartCount,
				OrderMoney:  combo.OrderMoney,
				ZkMoney:     combo.ZkMoney,
				SellStatus:  1,
				ZyksMoney:   3 * combo.KsMoney,
				ZyzkMoney:   3 * combo.ZkMoney,
			})
		}
	}
	return nil
}

func (p *CloudEyeList) AdminGetCloudEyeByParams(db *gorm.DB, info *form.Filter) (*[]CloudEyeList, int, error) {
	var combos []CloudEyeList
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&CloudEyeList{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&CloudEyeList{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&CloudEyeList{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		//if err := db.Model(&ScdnCombos{}).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
		if err := db.Table("CloudEyeList").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("CloudEyeList").Limit(info.Limit).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *CloudEyeList) UserGetCloudEyeByParams(db *gorm.DB, info *form.Filter) (*[]CloudEyeList, int, error) {
	var combos []CloudEyeList
	var total int

	if info.Field != "" {
		if info.Field == "tc_name" {
			if err := db.Model(&CloudEyeList{}).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`tc_name` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "ks_money" {
			if err := db.Model(&CloudEyeList{}).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`ks_money` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

		if info.Field == "domain_num" {
			if err := db.Model(&CloudEyeList{}).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("`sell_status` = ?", 1).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err := db.Limit(info.Limit).Where("`domain_num` LIKE ?", "%"+info.Value+"%").Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
				return nil, 0, err
			}
			return &combos, total, nil
		}

	}

	if info.Field == "" {
		//获取总记录数
		//if err := db.Model(&ScdnCombos{}).Where("sell_status = ?", 1).Count(&total).Error; err != nil {
		if err := db.Table("CloudEyeList").Where("sell_status = ?", 1).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		//分页查询
		if err := db.Table("CloudEyeList").Limit(info.Limit).Where("sell_status = ?", 1).Offset(info.Offset).Find(&combos).Error; err != nil {
			return nil, 0, err
		}
		return &combos, total, nil
	}
	return &combos, total, nil
}

func (p *DexunCloudEyeList) FindComboByUuid(db *gorm.DB, uuid string) *DexunCloudEyeList {
	info := &DexunCloudEyeList{}
	db.Model(&DexunCloudEyeList{}).Where("uuid = ?", uuid).Find(&info)
	return info
}
