package iredis

import (
	"context"
	"errors"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

var ErrGetLockFailed = errors.New("get lock failed")

const singleMode = "single"
const clusterMode = "cluster"

type IClient = redis.Cmdable

var client IClient
var lock *redislock.Client

type Config struct {
	// 连接模式，single / cluster
	Mode string
	// 连接地址，集群模式下使用逗号分隔多个地址
	Address string
	// 数据库密码
	Password string
	// 数据库，默认 0
	DB int
}

func Client() IClient {
	if client == nil {
		panic("redis client not initialized")
	}
	return client
}

func Init(config *Config) {
	if config.Mode == clusterMode {
		client = newClusterClient(config)
	} else {
		client = newSingleClient(config)
	}
	lock = newRedisLock(client)
}

// GetLock Get the lock, don't forget release Lock Release.
func GetLock(key string, duration time.Duration) (*redislock.Lock, error) {
	_lock, err := lock.Obtain(context.Background(), key, duration, nil)
	if errors.Is(err, redislock.ErrNotObtained) {
		return nil, ErrGetLockFailed
	} else if err != nil {
		return nil, err
	}
	return _lock, nil
}

func newSingleClient(config *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})

	// ping
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}

func newClusterClient(config *Config) *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(config.Address, ","),
		Password: config.Password,
	})

	// ping
	err := rdb.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	return rdb
}

func newRedisLock(rdb IClient) *redislock.Client {
	return redislock.New(rdb)
}
