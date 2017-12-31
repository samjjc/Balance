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
		case res := <-c:
			if res == -1 {
				work <- Request{workFunc, c}
			}
			// here would be where you'd process the results
		}
	}
}

func workFunc() int {
	randDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(randDuration)
	if randDuration <= time.Millisecond*800 {
		panic("GENERIC ERROR")
	}
	return 1
}
