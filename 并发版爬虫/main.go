
package main 

import (
	"groot/spider/engine"
	"groot/spider/zhenai/parser"
	"groot/spider/scheduler"
	"groot/spider/persist"
)


func  main() {

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}


















