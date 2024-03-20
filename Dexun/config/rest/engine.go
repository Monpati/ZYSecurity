package rest

import (
	"Dexun/config/base"
	"Dexun/utils"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	Engine_Module_Snow  = "snow"
	Engine_Module_Http  = "http"
	Engine_Module_Mysql = "mysql"
	Engine_Module_Redis = "redis"
	Engine_Module_ES    = "elasticsearch"
)

var EngineCfg *EngineConf

type EngineConf struct {
	Language *map[int]string

	Snow *utils.Snowflake

	HttpCfg *base.HttpConf

	MysqlCfg *base.MysqlConf
	MysqlDB  *gorm.DB

	RedisCfg    *base.RedisConf
	RedisClient *redis.Client
}

func (p *EngineConf) Set(module string, val interface{}) {
	if val == nil {
		return
	}
	switch module {
	case Engine_Module_Snow:
		if v, ok := val.(*utils.Snowflake); ok {
			p.Snow = v
		}
		break
	case Engine_Module_Http:
		if v, ok := val.(*base.HttpConf); ok {
			p.HttpCfg = v
		}
		break
	case Engine_Module_Mysql:
		if v, ok := val.(*base.MysqlConf); ok {
			p.MysqlCfg = v
			p.MysqlDB = v.Connect()
		}
		break
	case Engine_Module_Redis:
		if v, ok := val.(*base.RedisConf); ok {
			p.RedisCfg = v
			p.RedisClient = v.Connect()
		}
	}
}

func (p *EngineConf) GetCurrentTime() int {
	return int(time.Now().Unix())
}
