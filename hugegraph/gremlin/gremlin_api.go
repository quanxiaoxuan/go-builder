package gremlin

import (
	"encoding/json"
	"strings"

	"github.com/quanxiaoxuan/go-builder/http"
	"github.com/quanxiaoxuan/go-builder/hugegraph"
)

// POST请求参数
func InitGremlinParam(gremlin string) Param {
	var bindings interface{} // 构建绑定参数
	var aliases interface{}  // 构建图别名
	_ = json.Unmarshal([]byte(`{}`), &bindings)
	_ = json.Unmarshal([]byte(`{"graph": "hugegraph","g": "__g_hugegraph"}`), &aliases)
	return Param{Gremlin: gremlin, Bindings: bindings, Language: "gremlin-groovy", Aliases: aliases}
}

func GetGremlinHttpUrl(config *hugegraph.Config) string {
	url := strings.Builder{}
	url.WriteString(http.PrefixHttp)
	url.WriteString(config.Host)
	url.WriteString(`:`)
	url.WriteString(config.Port)
	url.WriteString(`/gremlin`)
	return url.String()
}

// 调用hugegraph的Get接口
func HttpGet(gremlin string) (result Result, err error) {
	var bytes []byte
	bytes, err = http.GetHttp(GetGremlinHttpUrl(hugegraph.InitConfig) + `?gremlin=` + gremlin)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return
	}
	return
}

// 调用hugegraph的POST接口
func HttpPost(gremlin string) (result Result, err error) {
	var bytes []byte
	bytes, err = http.PostHttp(GetGremlinHttpUrl(hugegraph.InitConfig), InitGremlinParam(gremlin))
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return
	}
	return
}

// 查询顶点
func QueryVertexsPost(gremlin string) (vertexs Vertexs, requestId string, err error) {
	var apiResult Result
	apiResult, err = HttpPost(gremlin)
	if err != nil {
		return
	}
	requestId = apiResult.RequestId
	// 将Result.Data序列化为json
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	// 将json反序列化为interface{}
	err = json.Unmarshal(bytes, &vertexs)
	if err != nil {
		return
	}
	return
}

// 查询顶点
func QueryVertexsGet(gremlin string) (vertexs Vertexs, requestId string, err error) {
	var apiResult Result
	apiResult, err = HttpGet(gremlin)
	if err != nil {
		return
	}
	requestId = apiResult.RequestId
	// 将Result.Data序列化为json
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	// 将json反序列化为interface{}
	err = json.Unmarshal(bytes, &vertexs)
	if err != nil {
		return
	}
	return
}

// 查询边
func QueryEdgesPost(gremlin string) (edges Edges, requestId string, err error) {
	var apiResult Result
	apiResult, err = HttpPost(gremlin)
	if err != nil {
		return
	}
	requestId = apiResult.RequestId
	// 将Result.Data序列化为json
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	// 将json反序列化
	err = json.Unmarshal(bytes, &edges)
	if err != nil {
		return
	}
	return
}

// 查询边
func QueryEdgesGet(gremlin string) (edges Edges, requestId string, err error) {
	var apiResult Result
	apiResult, err = HttpGet(gremlin)
	if err != nil {
		return
	}
	requestId = apiResult.RequestId
	// 将Result.Data序列化为json
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	// 将json反序列化
	err = json.Unmarshal(bytes, &edges)
	if err != nil {
		return
	}
	return
}

// 查询path()
func QueryPathsPost(gremlin string) (paths Paths, requestId string, err error) {
	var apiResult Result
	apiResult, err = HttpPost(gremlin)
	if err != nil {
		return
	}
	requestId = apiResult.RequestId
	// 将interface{}序列化为json
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	// 将json反序列化
	err = json.Unmarshal(bytes, &paths)
	return
}

// 查询path()
func QueryPathsGet(gremlin string) (paths Paths, requestId string, err error) {
	var apiResult Result
	apiResult, err = HttpGet(gremlin)
	if err != nil {
		return
	}
	requestId = apiResult.RequestId
	// 将interface{}序列化为json
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	// 将json反序列化
	err = json.Unmarshal(bytes, &paths)
	return
}

// 调用hugegraph的POST接口，返回属性值
func QueryValuesPost(gremlin string) (values []string, err error) {
	var apiResult Result
	apiResult, err = HttpPost(gremlin)
	if err != nil {
		return
	}
	var bytes []byte
	bytes, err = json.Marshal(apiResult.Result.Data)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &values)
	if err != nil {
		return
	}
	return
}
