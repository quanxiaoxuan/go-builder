package gremlin

// post接口请求参数
type Param struct {
	Gremlin  string      `json:"gremlin" `  // gremlin执行语句
	Bindings interface{} `json:"bindings" ` // 绑定参数
	Language string      `json:"language" ` // 语言
	Aliases  interface{} `json:"aliases" `  // 别名
}

// post接口返回结果
type Result struct {
	RequestId string     `json:"requestId"` // 请求ID
	Status    StatusInfo `json:"status"`    // 返回状态
	Result    ResultInfo `json:"result"`    // 返回结果
}

// 返回结果
type ResultInfo struct {
	Data interface{} `json:"data"` // 结果体
	Meta interface{} `json:"meta"` // 元数据
}

// 返回状态
type StatusInfo struct {
	Message    string      `json:"message"`    // 请求ID
	Code       int64       `json:"code"`       // 状态码
	Attributes interface{} `json:"attributes"` // 属性
}

// hugegraph查询【顶点】返回的data结果
type Vertexs []*Vertex
type Vertex struct {
	Id         string           `json:"id"`         // 主键id
	Label      string           `json:"label"`      // 类型
	Type       string           `json:"type"`       // 分类
	Properties VertexProperties `json:"properties"` // 属性
}

// 顶点属性
type VertexProperties struct {
	ObjId   string `json:"obj_id"`   // 顶点ID
	ObjName string `json:"obj_name"` // 顶点名称
	ObjType string `json:"obj_type"` // 顶点类型
	IsZdbq  int    `json:"is_zdbq"`  // 重点标签
	Sxh     int    `json:"sxh"`      // 顺序号
}

// hugegraph查询【边】返回的data结果
type Edges []*Edge
type Edge struct {
	Id         string         `json:"id"`         // ID
	Label      string         `json:"label"`      // 类型
	Type       string         `json:"type"`       // 分类
	OutV       string         `json:"outV"`       // 出顶点ID
	OutVLabel  string         `json:"outVLabel"`  // 出顶点类型
	InV        string         `json:"inV"`        // 类入顶点ID型
	InVLabel   string         `json:"inVLabel"`   // 入顶点类型
	Properties EdgeProperties `json:"properties"` // 属性
}

// 边属性
type EdgeProperties struct {
	RelationCode        string `json:"relation_code"`        // 关系编码
	RelationName        string `json:"relation_name"`        // 关系名称
	RelationDescription string `json:"relation_description"` // 关系描述
	RelationTypeLv1     string `json:"relation_type_lv1"`    // 关系类型1
	RelationTypeLv2     string `json:"relation_type_lv2"`    // 关系类型2
	PeerTimes           int32  `json:"peer_times"`           // 同行次数
	PeerDays            int32  `json:"peer_days"`            // 同行天数
	PeerSites           int32  `json:"peer_sites"`           // 同行点位数
	PeerDate            string `json:"peer_date"`            // 最后同行日期
	Sxh                 int    `json:"sxh"`                  // 顺序号
}

// hugegraph查询path()时返回的data结果
type Paths []*Path
type Path struct {
	Labels  interface{}    `json:"labels"`
	Objects PathProperties `json:"objects"`
}

// path()查询结果的【路径】中节点对象，包含顶点和边
type PathProperties []*PathProperty
type PathProperty struct {
	Id         string      `json:"id"`         // ID
	Label      string      `json:"label"`      // 类型
	Type       string      `json:"type"`       // 分类
	OutV       string      `json:"outV"`       // 出顶点ID
	OutVLabel  string      `json:"outVLabel"`  // 出顶点类型
	InV        string      `json:"inV"`        // 类入顶点ID型
	InVLabel   string      `json:"inVLabel"`   // 入顶点类型
	Properties interface{} `json:"properties"` // 对象属性
}
