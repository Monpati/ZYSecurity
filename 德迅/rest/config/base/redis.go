package base

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

var (
	avgRedisHost     = flag.String("redis_host", "127.0.0.1", "redis_host")
	avgRedisPort     = flag.Uint("redis_port", 6379, "redis_port")
	avgRedisPassword = flag.String("redis_password", "", "redis_password")
	//avgRedisMaxIdle     = flag.Uint("redis_max_idle", 3, "redis_max_idle")
	//avgRedisMaxAvtive   = flag.Uint("redis_max_active", 5, "redis_max_active")
	//avgRedisTimeoutIdle = flag.Uint("redis_timeout_idle", 240, "redis_timeout_idle")
)

type RedisConf struct {
	Host     string
	Port     int
	Password string
	//MaxIdle     int
	//MaxActive   int
	//TimeoutIdle int
}

func NewRedisConf() *RedisConf {
	return &RedisConf{
		Host:     *avgRedisHost,
		Port:     int(*avgRedisPort),
		Password: *avgRedisPassword,
		//MaxIdle:     int(*avgRedisMaxIdle),
		//MaxActive:   int(*avgRedisMaxAvtive),
		//TimeoutIdle: int(*avgRedisTimeoutIdle),
	}
}

func (p *RedisConf) LoadFromFile(name string) error {
	file, _ := os.Open(name)
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(p)
}

func (p *RedisConf) LoadFromBytes(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *RedisConf) SaveToFile(name string) error {
	flag := os.O_RDWR | os.O_TRUNC
	file, _ := os.OpenFile(name, flag, 0644)
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(p)
}

func (p *RedisConf) getUrl() string {
	return fmt.Sprintf("%s:%d", p.Host, p.Port)
}

func (p *RedisConf) Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     p.getUrl(),
		Password: p.Password,
	})
}
