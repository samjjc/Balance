package main

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *Worker) work(done chan *Worker) {
	var req Request

	defer func() {
		if err := recover(); err != nil {
			req.e <- ErrorString{err.(string)}
			done <- w
			w.work(done)
		}
	}()

	for {
		req = <-w.requests // get Request from balancer
		req.c <- req.fn()  // call fn and send result
		done <- w          // we've finished this request

	}
}
