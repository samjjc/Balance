package main

import "log"

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *Worker) Work(done chan *Worker) {
	var req Request

	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
			req.c <- -1
			done <- w
			w.Work(done)
		}
	}()

	for {
		req = <-w.requests // get Request from balancer
		req.c <- req.fn()  // call fn and send result
		done <- w          // we've finished this request

	}
}
