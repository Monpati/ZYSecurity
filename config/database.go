package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB
var rdb *redis.Client

func InitDB() *gorm.DB {
	var driverName, host, port, database, username, password string

	flag.StringVar(&driverName, "db_driver", "mysql", "Database driver name")
	flag.StringVar(&host, "db_host", "127.0.0.1", "Database host")
	flag.StringVar(&port, "db_port", "3306", "Database port")
	flag.StringVar(&database, "db_name", "Security", "Database name")
	flag.StringVar(&username, "db_user", "root", "Database username")
	flag.StringVar(&password, "db_pass", "HelloWorld", "Database password")

	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic(err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func GetRedis() *redis.Client {
	return rdb
}

func RedisInit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Success connect to redis")
}

func StoreCodeInRedis(client *redis.Client, code, key string) error {
	err := client.Set(key, code, 1*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetCodeFormRedis(client *redis.Client, key string) (string, error) {
	code, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

func DeleteCodeFromRedis(client *redis.Client, key string) error {
	err := client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func StoreTokenInRedis(client *redis.Client, token, key string) error {
	err := client.Set(key, token, 2*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetTokenFromRedis(client *redis.Client, key string) (string, error) {
	token, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("token not found")
	} else if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func DeleteTokenFromRedis(client *redis.Client, key string) error {
	err := client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
