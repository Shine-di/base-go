package redis

import (
	"cortex3/conf"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/qianlnk/log"
)

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Yaml.Conf.Redis.Host,
		Password: conf.Yaml.Conf.Redis.Pwd,
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic("redis无法连接 " + conf.Yaml.Conf.Redis.Pwd + err.Error())
	}
	fmt.Println("Redis 初始化成功")
}

func Redis() *redis.Client {
	if client == nil {
		InitRedis()
	}
	return client
}


func Set(key string, value string) {
	//client := Redis()
	//defer client.Close()

	Redis().Set(key, value, 0)
}

func Get(key string) (string, error) {
	//client := Redis()
	//defer client.Close()
	log.Info(key)
	log.Info(Redis())
	return Redis().Get(key).Result()
}

func Del(key string) {
	//client := Redis()
	//defer client.Close()

	Redis().Del(key)
}