package engine

// ParserFunc解析函数类型声明
type ParserFunc func([]byte) ParseResult

// ParseResult解析响应结果类型声明
type ParseResult struct {
	Requests []Request // 请求
	Items    []Item    // 数据
}

// Item解析响应数据项类型声明
type Item struct {
	Id      string      `json:"id"`
	Url     string      `json:"url"`
	Payload interface{} `json:"payload"`
}

// Request解析响应请求项类型声明
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

// NilParser一个空的解析器
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
