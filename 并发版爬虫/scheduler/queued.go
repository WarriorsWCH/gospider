
package scheduler

import(
	"groot/spider/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	// 每个worker建立不同的chan，都放在chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request{

	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// worker准备好接收外面的request
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}


func (s *QueuedScheduler) Run() {
	// 生成QueuedScheduler中的chan
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 既有request排队 又有worker排队
			if len(requestQ) >0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
				// 收到的request缓存起来
			case r := <-s.requestChan:
				// 收到 排队
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				// 收到 排队
				workerQ = append(workerQ, w)
				// 如果同时又request和worker，就把request发给worker
			case activeWorker <- activeRequest:
				// 从队列中拿掉
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}










