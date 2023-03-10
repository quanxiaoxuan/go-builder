package database

// 数据源配置
type Config struct {
	Dialect  string `json:"dialect" yaml:"dialect" nacos:"database.dialect"`    // 数据库类型
	Host     string `json:"host" yaml:"host" nacos:"database.host"`             // 数据库Host
	Port     string `json:"port" yaml:"port" nacos:"database.port"`             // 数据库端口
	Database string `json:"database" yaml:"database" nacos:"database.database"` // 数据库名
	UserName string `json:"userName" yaml:"userName" nacos:"database.userName"` // 用户名
	Password string `json:"password" yaml:"password" nacos:"database.password"` // 密码
	Debug    bool   `json:"debug" yaml:"debug" nacos:"database.debug"`          // 是否开启debug
}
