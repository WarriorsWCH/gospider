package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}


// 解析结果
type ParseResult struct {
	Requests []Request //城市url
	Items []interface{} //城市名字
}

// 空的解析函数
func NilParser([]byte) ParseResult {

	return ParseResult{}
}