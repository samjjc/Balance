package main

import (
	"time"
)

type Request struct {
	fn func() int
	c  chan int
}

func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		time.Sleep(2 * time.Second)
		work <- Request{workFunc, c} // send request
		<-c                          // wait for answer
		//furtherProcess(result)
	}
}

func workFunc() int {
	time.Sleep(time.Second)
	return 1
}
