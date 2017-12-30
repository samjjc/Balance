package main

import (
	"time"
)

type Request struct {
	fn func() int
	c  chan int
}

func requester(work chan<- Request) {
	c := make(chan int, 5)
	for {
		time.Sleep(time.Second / 3)
		select {
		case work <- Request{workFunc, c}:
		case <-c:
		}
	}
}

func workFunc() int {
	time.Sleep(2 * time.Second)
	return 1
}
