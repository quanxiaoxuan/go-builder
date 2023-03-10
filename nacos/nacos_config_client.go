package nacos

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

var configClient config_client.IConfigClient = nil

// 创建Nacos配置客户端
func CreateConfigClient() (err error) {
	if configClient != nil {
		return
	}
	clientConfig := constant.ClientConfig{
		TimeoutMs:           10 * 1000,
		BeatInterval:        3 * 1000,
		NotLoadCacheAtStart: true,
		NamespaceId:         access.NameSpace,
		LogDir:              access.LogDir,
		CacheDir:            access.CacheDir,
		Username:            access.Username,
		Password:            access.Password,
	}
	serverConfigs := make([]constant.ServerConfig, 0)
	adds := strings.Split(access.NacosAddress, ",")
	if len(adds) == 0 {
		return errors.New("NACOS服务地址 不能为空")
	}
	for _, addStr := range adds {
		add := strings.Split(addStr, ":")
		port, _ := strconv.ParseInt(add[1], 10, 64)
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			ContextPath: access.ContextPath,
			IpAddr:      add[0],
			Port:        uint64(port),
		})
	}
	// 创建服务发现客户端
	configClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	return
}
