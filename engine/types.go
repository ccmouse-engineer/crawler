package engine

// ParseResult代表解析响应结果结构体
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// Request代表请求参数结构体
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// NilParser代表一个空的解析器
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
