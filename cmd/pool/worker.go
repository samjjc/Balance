package main

import (
	"github.com/samjjc/basicLoadBalancer"
)

type Worker struct {
	requests chan request.Request
	pending  int
	index    int
}

func (w *Worker) work(done chan *Worker) {
	var req request.Request

	defer func() {
		if err := recover(); err != nil {
			req.E <- request.ErrorString{err.(string)}
			done <- w
			w.work(done)
		}
	}()

	for {
		req = <-w.requests // get Request from balancer
		req.C <- req.Fn()  // call fn and send result
		done <- w          // we've finished this request

	}
}
