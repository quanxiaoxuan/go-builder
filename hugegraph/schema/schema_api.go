package schema

import (
	"encoding/json"
	"strings"

	"github.com/quanxiaoxuan/go-builder/http"
	"github.com/quanxiaoxuan/go-builder/hugegraph"
	log "github.com/sirupsen/logrus"
)

func GetSchemaHttpUrl(config *hugegraph.Config, suffixUrl string) string {
	url := strings.Builder{}
	url.WriteString(http.PrefixHttp)
	url.WriteString(config.Host)
	url.WriteString(`:`)
	url.WriteString(config.Port)
	url.WriteString(`/graphs/`)
	url.WriteString(config.Graph)
	url.WriteString(`/schema/`)
	url.WriteString(suffixUrl)
	return url.String()
}

func SchemaPost(suffixUrl string, param interface{}) (interface{}, error) {
	url := GetSchemaHttpUrl(hugegraph.InitConfig, suffixUrl)
	log.Info("请求URL : ", url)
	responseByte, err := http.PostHttp(url, param)
	var hugeGraphResult interface{}
	_ = json.Unmarshal(responseByte, &hugeGraphResult)
	return hugeGraphResult, err
}

func SchemaGet(suffixUrl string) (interface{}, error) {
	url := GetSchemaHttpUrl(hugegraph.InitConfig, suffixUrl)
	log.Info("请求URL : ", url)
	responseByte, err := http.GetHttp(url)
	var hugeGraphResult interface{}
	_ = json.Unmarshal(responseByte, &hugeGraphResult)
	return hugeGraphResult, err
}

func SchemaPut(suffixUrl string, param interface{}) (interface{}, error) {
	url := GetSchemaHttpUrl(hugegraph.InitConfig, suffixUrl)
	log.Info("请求URL : ", url)
	responseByte, err := http.PutHttp(url, param)
	var hugeGraphResult interface{}
	_ = json.Unmarshal(responseByte, &hugeGraphResult)
	return hugeGraphResult, err
}

func SchemaDelete(suffixUrl string) (interface{}, error) {
	url := GetSchemaHttpUrl(hugegraph.InitConfig, suffixUrl)
	log.Info("请求URL : ", url)
	responseByte, err := http.DeleteHttp(url)
	var hugeGraphResult interface{}
	_ = json.Unmarshal(responseByte, &hugeGraphResult)
	return hugeGraphResult, err
}
