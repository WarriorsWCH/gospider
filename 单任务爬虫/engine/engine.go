
package engine

import (
	"log"
	"groot/spider/fetcher"
)

func Run(seeds ...Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 循环一个接一个
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)

		// body 请求到的html
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error " + 
				"fetching url %s: %v", r.Url, err)
			continue
		}
		// ParserFunc(body) 解析body
		parseResult := r.ParserFunc(body)
		// 把解析后的城市url添加到请求request队列中
		requests = append(requests,
			parseResult.Requests...)
		// parseResult.Requests... 数组或者slice展开成每个元素

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}