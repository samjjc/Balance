package main

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}
