package hugegraph

const (
	Propertykeys = "propertykeys" // 属性API
	Vertexlabels = "vertexlabels" // 顶点API
	Edgelabels   = "edgelabels"   // 边API
	Indexlabels  = "indexlabels"  // 边API
	Multiple     = "MULTIPLE"     // 允许多次连接
	Single       = "SINGLE"       // 单条连接
)

// 属性数据类型
var DateTypesAll = []string{"TEXT", "BYTE", "DATE", "INT", "LONG", "BOOLEAN", "DOUBLE", "FLOAT", "UUID", "BLOB", "OBJECT", "UNKNOWN"}

// 属性类型基数
var CardinalityAll = []string{"SINGLE", "SET", "LIST"}

// ID策略
var IdStrategyAll = []string{"DEFAULT", "AUTOMATIC", "PRIMARY_KEY", "CUSTOMIZE_STRING", "CUSTOMIZE_NUMBER", "CUSTOMIZE_UUID"}

// 边线条粗细
var ThicknessAll = []string{"THICK", "NORMAL", "FINE"}

// 顶点样式大小
var SizeAll = []string{"TINY", "SMALL", "NORMAL", "BIG", "HUGE"}

// 模型类型
var BaseTypeAll = []string{"VERTEX_LABEL", "EDGE_LABEL"}

// 索引类型
var IndexTypeAll = []string{"SECONDARY", "RANGE", "SEARCH"}
