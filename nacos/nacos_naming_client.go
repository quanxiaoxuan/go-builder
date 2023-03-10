package nacos

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	log "github.com/sirupsen/logrus"
)

var namingClient naming_client.INamingClient = nil

// 创建nacos服务发现客户端
func CreateNamingClient() (err error) {
	if namingClient != nil {
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
	// 至少一个ServerConfig
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
	namingClient, err = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	return
}

// 注册Nacos服务实例
func RegisterInstance() {
	port, _ := strconv.Atoi(register.Port)
	_, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          register.Ip,
		Port:        uint64(port),
		GroupName:   register.Group,
		ServiceName: register.Name,
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    nil,
	})
	if err != nil {
		log.Error("=== 注册Nacos服务实例失败 : ", err)
	} else {
		log.Info("=== 注册Nacos服务实例成功 ===")
	}
}

// 随机获取一个健康的服务实例
func SelectOneHealthyInstance(serviceName, groupName string) (string, error) {
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", instance.Ip, instance.Port), nil
}

// 获取所有健康服务实例
func SelectInstances(serviceName, groupName string) ([]string, error) {
	instances, err := namingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		HealthyOnly: true,
	})
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	for _, instance := range instances {
		res = append(res, fmt.Sprintf("%s:%d", instance.Ip, instance.Port))
	}
	return res, nil
}
