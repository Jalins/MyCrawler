package engine

// 这里数据结构的理解：一个页面含有多个链接（urls）以及链接所对应的名称（items），而每一个url链接过去又是一个包含多个url以及item的页面
// 说白了就是传入一个url地址，然后使用传入的函数对这个地址进行解析，得到一系列名称以及对应的request，一个request相当于一个页面
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func ParserFunc([]byte) ParseResult {
	return ParseResult{}
}
