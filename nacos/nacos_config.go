package nacos

import "github.com/nacos-group/nacos-sdk-go/vo"

const (
	System    = "system"
	Database  = "database"
	Redis     = "redis"
	Log       = "log"
	Gateway   = "gateway"
	Hugegraph = "hugegraph"
)

// 配置项
type Configs []Config
type Config struct {
	Name   string `json:"name" yaml:"name"`     // Nacos配置项
	Group  string `json:"group" yaml:"group"`   // Nacos配置分组
	DataId string `json:"dataId" yaml:"dataId"` // Nacos配置文件
}

func (param Config) Translate() vo.ConfigParam {
	return vo.ConfigParam{DataId: param.DataId, Group: param.Group}
}

func (list Configs) Get(name string) Config {
	for _, item := range list {
		if item.Name == name {
			return item
		}
	}
	return list[0]
}
