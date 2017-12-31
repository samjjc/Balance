package main

import (
	"math/rand"
	"time"
)

type Request struct {
	fn func() int
	c  chan int
}

func requester(work chan<- Request) {
	c := make(chan int, 5)
	for {
		randDuration := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(randDuration + 100*time.Millisecond)
		select {
		case work <- Request{workFunc, c}:
		case <-c:
		}
	}
}

func workFunc() int {
	randDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(randDuration + time.Second)
	return 1
}
