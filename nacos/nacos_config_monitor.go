package nacos

import (
	"sync"
	"time"
)

var ConfigMonitor *Monitor

// nacos配置监控
type Monitor struct {
	mu        sync.RWMutex                     //互斥锁
	ConfigMap map[string]map[string]ConfigData //配置数据
	ConfigNum int                              //配置数量
}

// 配置数据
type ConfigData struct {
	Content    string // 配置内容
	Changed    bool   // 修改标识
	UpdateTime int64  // 修改时间
}

// 初始化nacos配置监控
func InitNacosConfigMonitor() {
	var syncOnce sync.Once
	syncOnce.Do(func() {
		ConfigMonitor = &Monitor{
			ConfigMap: make(map[string]map[string]ConfigData),
			ConfigNum: 0,
		}
	})
	return
}

// 获取nacos配置
func (monitor *Monitor) GetConfig(dataId, group string) string {
	monitor.mu.Lock()
	defer monitor.mu.Unlock()
	dataMap, ok := monitor.ConfigMap[group]
	if ok {
		_, ok = dataMap[dataId]
		if ok {
			return dataMap[dataId].Content
		}
	}
	return ""
}

// 保存nacos配置
func (monitor *Monitor) SaveConfig(dataId, group, content string) {
	monitor.mu.Lock()
	defer monitor.mu.Unlock()
	var dataMap = make(map[string]ConfigData)
	dataMap[dataId] = ConfigData{content, false, time.Now().UnixMilli()}
	monitor.ConfigNum = monitor.ConfigNum + 1
	monitor.ConfigMap[group] = dataMap
	return
}

// 修改nacos配置
func (monitor *Monitor) ChangeConfig(dataId, group, content string) {
	monitor.mu.Lock()
	defer monitor.mu.Unlock()
	dataMap, ok := monitor.ConfigMap[group]
	if ok {
		dataMap[dataId] = ConfigData{content, true, time.Now().UnixMilli()}
		monitor.ConfigMap[group] = dataMap
	}
	return
}
