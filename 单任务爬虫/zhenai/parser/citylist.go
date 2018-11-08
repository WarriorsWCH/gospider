package parser

import (
	"regexp"
	"groot/spider/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	// [^>]* 除了^的任意字符
	re := regexp.MustCompile(cityListRe)
	// 返回值是[][]byte 返回所有
	// matches := re.FindAll(contents, -1)
	// 返回值是[][][]byte
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	limit := 10
	for _, m := range matches {
		// for _, subM := range m {
		// 	fmt.Printf("%s",subM)
		// }
		// fmt.Println()

		result.Items = append(result.Items, "City:"+string(m[2]))
		// 把城市的url添加到result.Requests数组中，
		// 解析函数为NilParser，也就是暂时不解析
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: ParseCity,
			})
		// fmt.Printf("City: %s, URL: %s\n", m[2],m[1])

		limit--
		if limit == 0 {
			break
		}
	}
	return result
}