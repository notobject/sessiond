package dao

import (
	"sync"

	"github.com/cockroachdb/errors"
	"gopkg.in/redis.v4"

	"github.com/notobject/sessiond/pkg/config"
)

type dao struct{}

var instance *dao
var once sync.Once

// NewDao 创建DAO
func NewDao() *dao {
	once.Do(func() {
		instance = &dao{}
	})
	return instance
}

// GetRedisClient 获取Redis客户端
func (dao *dao) GetRedisClient() (*redis.Client, error) {
	c := config.GetConfig().Redis

	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	if pong != "PONG" {
		return nil, errors.Errorf("ping redis failed! response: %s", pong)
	}
	return client, nil
}
