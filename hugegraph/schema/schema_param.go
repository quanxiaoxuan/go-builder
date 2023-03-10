package schema

// 新增属性参数
type PropertyAddParam struct {
	Name        string `json:"name"`        // 属性名称
	DataType    string `json:"data_type"`   // 属性类型
	Cardinality string `json:"cardinality"` // 属性类型基数
}

// 新增顶点参数
type VertexAddParam struct {
	Name             string            `json:"name"`               // 顶点名称
	IdStrategy       string            `json:"id_strategy"`        // 主键策略
	Properties       []string          `json:"properties"`         // 属性列表
	PrimaryKeys      []string          `json:"primary_keys"`       // 主键属性列表
	NullableKeys     []string          `json:"nullable_keys"`      // 可空属性列表
	IndexLabels      []string          `json:"index_labels"`       // 索引列表
	Ttl              int               `json:"ttl"`                // TTL
	EnableLabelIndex bool              `json:"enable_label_index"` // 启用类型索引,默认为true
	UserData         map[string]string `json:"user_data"`          // 顶点风格配置
}

// 新增边参数
type EdgeAddParam struct {
	Name             string            `json:"name"`               // 边名称
	SourceLabel      string            `json:"source_label"`       // 源顶点类型
	TargetLabel      string            `json:"target_label"`       // 目标顶点类型
	Properties       []string          `json:"properties"`         // 属性列表
	NullableKeys     []string          `json:"nullable_keys"`      // 可空属性列表
	Frequency        string            `json:"frequency"`          // 允许多次连接，可以取值SINGLE和MULTIPLE
	SortKeys         []string          `json:"sort_keys"`          // 当允许关联多次时，指定区分键属性列表
	Ttl              int               `json:"ttl"`                // TTL
	EnableLabelIndex bool              `json:"enable_label_index"` // 启用类型索引,默认为true
	UserData         map[string]string `json:"user_data"`          // 边风格配置
}

// 索引新增参数
type IndexAddParam struct {
	Name      string   `json:"name"`       // 索引名称
	BaseType  string   `json:"base_type"`  // 模型类型
	BaseValue string   `json:"base_value"` // 模型名称
	IndexType string   `json:"index_type"` // 索引类型
	Fields    []string `json:"fields"`     // 属性列表
}

// hugegraph-api-append请求接口参数
type PropertiesAppendParam struct {
	Name         string            `json:"name"`          // 名称
	Properties   []string          `json:"properties"`    // 属性列表
	NullableKeys []string          `json:"nullable_keys"` // 可空属性列表
	UserData     map[string]string `json:"user_data"`     // 风格配置
}
