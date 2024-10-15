package base

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	avgMysqlDialect      = flag.String("mysql_dialect", "mysql", "mysql_dialect")
	avgMysqlDatabase     = flag.String("mysql_database", "Security", "mysql_database")
	avgMysqlUsername     = flag.String("mysql_username", "root", "mysql_username")
	avgMysqlPassword     = flag.String("mysql_password", "HelloWorld", "mysql_password")
	avgMysqlCharset      = flag.String("mysql_charset", "utf8", "mysql_charset")
	avgMysqlHost         = flag.String("mysql_host", "localhost", "mysql_host")
	avgMysqlPort         = flag.Uint("mysql_port", 3306, "mysql_port")
	avgMysqlMaxIdleConns = flag.Uint("mysql_max_idle_conns", 5, "mysql_max_idle_conns")
	avgMysqlMaxOpenConns = flag.Uint("mysql_max_open_conns", 10, "mysql_max_open_conns")
)

type MysqlConf struct {
	Dialect      string
	Database     string
	Username     string
	Password     string
	Charset      string
	Host         string
	Port         int
	MaxIdleConns int
	MaxOpenConns int
}

func NewMysqlConf() *MysqlConf {
	return &MysqlConf{
		Dialect:      *avgMysqlDialect,
		Database:     *avgMysqlDatabase,
		Username:     *avgMysqlUsername,
		Password:     *avgMysqlPassword,
		Charset:      *avgMysqlCharset,
		Host:         *avgMysqlHost,
		Port:         int(*avgMysqlPort),
		MaxIdleConns: int(*avgMysqlMaxIdleConns),
		MaxOpenConns: int(*avgMysqlMaxOpenConns),
	}
}

func (p *MysqlConf) LoadFromFile(name string) error {
	file, _ := os.Open(name)
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(p)
}

func (p *MysqlConf) LoadFromBytes(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *MysqlConf) SaveToFile(name string) error {
	flag := os.O_RDWR | os.O_TRUNC
	file, _ := os.OpenFile(name, flag, 0644)
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(p)
}

func (p *MysqlConf) getUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		p.Username,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
		p.Charset)
}

func (p *MysqlConf) Connect() *gorm.DB {
	db, err := gorm.Open(p.Dialect, p.getUrl())
	if err != nil {
		panic(err)
	}
	//if config.Env == DevelopmentMode {
	//	db.LogMode(true)
	//}
	db.SingularTable(true) //默认表名不用加s
	db.DB().SetMaxIdleConns(p.MaxIdleConns)
	db.DB().SetMaxOpenConns(p.MaxOpenConns)

	return db
}
