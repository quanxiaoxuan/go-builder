package redis

import (
	"time"
)

type Config struct {
	Host         string        `json:"host" yaml:"host" nacos:"redis.host"`             // 主机
	Port         int           `json:"port" yaml:"port" nacos:"redis.port"`             // 端口
	Password     string        `json:"password" yaml:"password" nacos:"redis.password"` // 密码
	Mode         int           `json:"mode" yaml:"mode" nacos:"redis.mode"`             // 模式(0-单机;1-集群)
	Database     int           `json:"database" yaml:"database" nacos:"redis.database"` // 数据库，默认0
	PoolSize     int           `json:"poolSize" yaml:"poolSize"`
	MinIdleConns int           `json:"minIdleConns" yaml:"minIdleConns"`
	ReadTimeout  time.Duration `json:"readTimeout" yaml:"readTimeout"`
}
