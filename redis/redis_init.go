package redis

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var InitCmd *Cmdable

// 初始化redis
func InitRedis(config *Config) {
	if InitCmd == nil {
		InitCmd = InitRedisCmd(config)
		_, err := InitCmd.Ping().Result()
		if err != nil {
			log.Error("=== redis连接失败 : ", err)
		} else {
			log.Info("=== redis连接成功 ===")
		}
	}
}

const (
	StandAlone = iota // 单机
	Cluster           // 集群
)

// 初始化redis命令窗
func InitRedisCmd(config *Config) *Cmdable {
	var client *Cmdable
	if config.Mode == Cluster {
		client = &Cmdable{
			config: config,
			cmd: redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:        strings.Split(config.Host, ","),
				Password:     config.Password,
				PoolSize:     config.PoolSize,
				MinIdleConns: config.MinIdleConns,
				ReadTimeout:  config.ReadTimeout * time.Second,
			}),
		}
	} else if config.Mode == StandAlone {
		client = &Cmdable{
			config: config,
			cmd: redis.NewClient(&redis.Options{
				Addr:         config.Host + ":" + strconv.Itoa(config.Port),
				Password:     config.Password,
				DB:           config.Database,
				PoolSize:     config.PoolSize,
				MinIdleConns: config.MinIdleConns,
				ReadTimeout:  config.ReadTimeout * time.Second,
			}),
		}
	} else {
		log.Warnf("未知redis模式 , %+v", config)
		client = nil
	}
	return client
}
