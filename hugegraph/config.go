package hugegraph

var InitConfig *Config

// 初始化配置
func InitHugegraph(config *Config) {
	if InitConfig == nil {
		InitConfig = config
	}
}

// hugegraph配置
type Config struct {
	Host  string `json:"host" yaml:"host" nacos:"hugegraph.host"`    // 主机
	Port  string `json:"port" yaml:"port" nacos:"hugegraph.port"`    // 端口
	Graph string `json:"graph" yaml:"graph" nacos:"hugegraph.graph"` // 图名称
}
