package nacos

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/magiconair/properties"
	"github.com/nacos-group/nacos-sdk-go/vo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// 读取Nacos配置并监听
func ReadAndListenNacosConfig(param vo.ConfigParam) (string, error) {
	// 读取Nacos配置
	content, err := configClient.GetConfig(param)
	if err != nil {
		log.Warn("读取Nacos配置失败 ", err)
		return "", err
	}
	ConfigMonitor.SaveConfig(param.DataId, param.Group, content)
	// 监听nacos配置改动
	err = ListenConfigChange(param)
	if err != nil {
		log.Warn("监听nacos配置改动失败 ", err)
		return "", err
	}
	return content, nil
}

// 加载yml配置
func LoadYmlConfig(param vo.ConfigParam, config interface{}) error {
	valueRef := reflect.ValueOf(config)
	// 修改值必须是指针类型否则不可行
	if valueRef.Type().Kind() != reflect.Ptr {
		log.Warn("加载yml配置异常，入参conf不是指针类型")
		return errors.New("入参conf不是指针类型")
	}
	// 读取Nacos配置并监听
	content, err := ReadAndListenNacosConfig(param)
	if err != nil {
		log.Warn("读取Nacos配置或者监听失败 ", err)
		return err
	}
	return yaml.Unmarshal([]byte(content), config)
}

// 加载Json配置
func LoadJsonConfig(param vo.ConfigParam, config interface{}) error {
	valueRef := reflect.ValueOf(config)
	// 修改值必须是指针类型否则不可行
	if valueRef.Type().Kind() != reflect.Ptr {
		log.Warn("加载yml配置异常，入参conf不是指针类型")
		return errors.New("入参conf不是指针类型")
	}
	// 读取Nacos配置并监听
	content, err := ReadAndListenNacosConfig(param)
	if err != nil {
		log.Warn("读取Nacos配置或者监听失败 ", err)
		return err
	}
	return json.Unmarshal([]byte(content), &config)
}

// 加载properties配置
func LoadPropertiesConfig(param vo.ConfigParam, config interface{}) error {
	valueRef := reflect.ValueOf(config)
	// 修改值必须是指针类型否则不可行
	if valueRef.Type().Kind() != reflect.Ptr {
		log.Warn("加载yml配置异常，入参conf不是指针类型")
		return errors.New("入参conf不是指针类型")
	}
	// 读取Nacos配置并监听
	content, err := ReadAndListenNacosConfig(param)
	if err != nil {
		log.Warn("读取Nacos配置或者监听失败 ", err)
		return err
	}
	var p *properties.Properties
	p, err = properties.LoadString(content)
	if err != nil {
		return err
	}
	refreshValueByTag(p, valueRef)
	return nil
}

// 批量加载properties配置
func LoadPropertiesConfigBatch(configParams []vo.ConfigParam, config interface{}) error {
	for _, param := range configParams {
		err := LoadPropertiesConfig(param, config)
		if err != nil {
			log.Warn("加载properties配置异常 ", err)
			return err
		}
	}
	return nil
}

// 监听nacos配置改动
func ListenConfigChange(param vo.ConfigParam) error {
	param.OnChange = func(namespace, group, dataId, data string) {
		log.Warnf("nacos配置已改动 \nnamespace :%s\nGroup     :%s\nData Id   :%s\n%s", namespace, group, dataId, data)
		ConfigMonitor.ChangeConfig(dataId, group, data)
	}
	err := configClient.ListenConfig(param)
	if err != nil {
		log.Errorf("nacos配置改动监听异常 \nGroup     :%s\nData Id   :%s\nerr: %s", param.Group, param.DataId, err)
	}
	return err
}

// 通过配置标签刷新配置
func refreshValueByTag(p *properties.Properties, v reflect.Value) {
	m := p.Map()
	for i := 0; i < v.Elem().NumField(); i++ {
		//先判断有没有nacos的配置
		tag := v.Elem().Type().Field(i).Tag.Get("nacos")
		r, _ := regexp.Compile("\\${.*?}")
		gs := r.FindAllString(tag, -1)
		for _, str := range gs {
			if len(str) <= 3 {
				tag = strings.ReplaceAll(tag, str, "")
			} else {
				envStr := str[2 : len(str)-1]
				tag = strings.ReplaceAll(tag, str, strings.Split(v.FieldByName(envStr).String(), ".")[0])
			}
		}
		if tag == "" && reflect.Struct != v.Elem().Field(i).Kind() {
			continue
		}
		switch v.Elem().Field(i).Kind() {
		case reflect.String:
			temp, ok := p.Get(tag)
			if ok {
				v.Elem().Field(i).SetString(temp)
			}
		case reflect.Bool:
			_, ok := m[tag]
			if ok {
				temp := p.GetBool(tag, false)
				v.Elem().Field(i).SetBool(temp)
			}
		case reflect.Int:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Int8:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Int16:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Int32:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Int64:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Uint8:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Uint16:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Uint32:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Uint64:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Elem().Field(i).SetInt(int64(temp))
			}
		case reflect.Float32:
			_, ok := m[tag]
			if ok {
				temp := p.GetFloat64(tag, 0)
				v.Elem().Field(i).SetFloat(temp)
			}
		case reflect.Float64:
			_, ok := m[tag]
			if ok {
				temp := p.GetFloat64(tag, 0)
				v.Elem().Field(i).SetFloat(temp)
			}
		case reflect.Struct:
			refreshStructValueByTag(p, v.Elem().Field(i))
		default:
			fmt.Printf("未匹配到type %s", v.Elem().Field(i).Kind())
		}
	}
}

// 通过配置标签刷新结构体配置
func refreshStructValueByTag(p *properties.Properties, v reflect.Value) {
	m := p.Map()
	for i := 0; i < v.NumField(); i++ {
		//先判断有没有nacos的配置
		tag := v.Type().Field(i).Tag.Get("nacos")
		r, _ := regexp.Compile("\\${.*?}")
		gs := r.FindAllString(tag, -1)
		for _, str := range gs {
			if len(str) <= 3 {
				tag = strings.ReplaceAll(tag, str, "")
			} else {
				envStr := str[2 : len(str)-1]
				tag = strings.ReplaceAll(tag, str, strings.Split(v.FieldByName(envStr).String(), ".")[0])
			}
		}

		if tag == "" && reflect.Struct != v.Field(i).Kind() {
			continue
		}

		switch v.Field(i).Kind() {
		case reflect.String:
			temp, ok := p.Get(tag)
			if ok {
				v.Field(i).SetString(temp)
			}
		case reflect.Bool:
			_, ok := m[tag]
			if ok {
				temp := p.GetBool(tag, false)
				v.Field(i).SetBool(temp)
			}
		case reflect.Int:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Int8:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Int16:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Int32:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Int64:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Uint8:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Uint16:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Uint32:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Uint64:
			_, ok := m[tag]
			if ok {
				temp := p.GetInt(tag, 0)
				v.Field(i).SetInt(int64(temp))
			}
		case reflect.Float32:
			_, ok := m[tag]
			if ok {
				temp := p.GetFloat64(tag, 0)
				v.Field(i).SetFloat(temp)
			}
		case reflect.Float64:
			_, ok := m[tag]
			if ok {
				temp := p.GetFloat64(tag, 0)
				v.Field(i).SetFloat(temp)
			}
		case reflect.Struct:
			refreshStructValueByTag(p, v.Field(i))
		default:
			fmt.Printf("未匹配到type %s", v.Field(i).Kind())
		}
	}
}
