
package engine

import(
	// "log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface{
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 输入输出
	// in := make(chan Request)
	out := make(chan ParseResult)
	// e.Scheduler.ConfigMasterWorkerChan(in)
	e.Scheduler.Run()
	// 创建两个chan 然后等待任务的到来

	//创建worker
	for i :=0; i < e.WorkerCount; i ++ {
		// createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	
	for {
		// 收
		result := <-out
		for _, item := range result.Items{

			 go func () {
			 	e.ItemChan <- item
			 }()
		}

		// 把所有的request发给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker (in chan Request, out chan ParseResult, s Scheduler){
	// in是自己的chan
	// in := make(chan Request)
	// 创建10个gorotine
	go func() {
		for{
			// tell scheduler i`m ready
			s.WorkerReady(in)
			// 把in接收到request读出来
			request := <- in
			// 使用worker解析request
			result, err := worker(request)
			if err != nil {
				continue
			}
			// out接收结果
			out <- result
		}
	}()
}












