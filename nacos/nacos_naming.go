package nacos

var register *ServerRegister

// 初始化Nacos服务注册配置
func InitNacosServerRegister(sr *ServerRegister) {
	if register == nil {
		register = sr
	}
}

// 服务注册
type ServerRegister struct {
	Ip    string `yaml:"ip"`    // IP
	Port  string `yaml:"port"`  // 端口
	Group string `yaml:"group"` // 分组
	Name  string `yaml:"name"`  // 名称
}
