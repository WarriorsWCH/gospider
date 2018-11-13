package parser

import (
	"regexp"
	"groot/spider/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	// [^>]* 除了^的任意字符
	re := regexp.MustCompile(cityRe)
	// 返回值是[][]byte 返回所有
	// matches := re.FindAll(contents, -1)
	// 返回值是[][][]byte
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		// for _, subM := range m {
		// 	fmt.Printf("%s",subM)
		// }
		// fmt.Println()
		name := string(m[2])
		// result.Items = append(result.Items, "User:"+string(m[2]))
		// 把城市的url添加到result.Requests数组中，
		// 解析函数为NilParser，也就是暂时不解析
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				// ParserFunc: engine.NilParser,
				ParserFunc: func (c []byte) engine.ParseResult {
					// 直接使用m[2]会有问题，因为这个函数是在for循环结束之后开始运行（排队）
					// return ParseProfile(c, string(m[2]))
					return ParseProfile(c, name)
				},
			})
		// fmt.Printf("City: %s, URL: %s\n", m[2],m[1])
	}
	return result
}