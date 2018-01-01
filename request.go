package main

import (
	"log"
	"math/rand"
	"time"
)

type Request struct {
	Fn func() int
	C  chan int
	E  chan error
}

func requester(work chan<- Request) {
	c := make(chan int, 5)
	e := make(chan error, 5)
	for {
		randDuration := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(randDuration + 100*time.Millisecond)
		select {
		case work <- Request{workFunc, c, e}:
		case <-c:
			// here would be where you'd process the results
		case err := <-e:
			log.Printf("%v\n", err)
			work <- Request{workFunc, c, e}
		}
	}
}

func workFunc() int {
	randDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(randDuration)
	if randDuration <= time.Millisecond*400 {
		panic("3 FAST 5 ME")
	}
	return 1
}
