
package main 

import (
	"groot/spider/engine"
	"groot/spider/zhenai/parser"
)


func  main() {

	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}


















