package nacos

var access *Access

// 初始化Nacos访问配置
func InitNacosConfig(c *Access) {
	if access == nil {
		access = c
	}
}

// nacos访问配置
type Access struct {
	NacosAddress string `yaml:"nacosAddress" default:"localhost:8849"` // NACOS服务地址
	ContextPath  string `yaml:"contextPath" default:"/nacos"`          // web访问路径
	Username     string `yaml:"username" default:"nacos"`              // 用户名
	Password     string `yaml:"password" default:"nacos"`              // 密码
	NameSpace    string `yaml:"nameSpace" default:"public"`            // 命名空间
	LogDir       string `yaml:"logDir" default:".nacos/log"`           // 日志文件夹
	CacheDir     string `yaml:"cacheDir" default:".nacos/cache"`       // 缓存文件夹
}
