package main

import (
	"log"
	"math/rand"
	"time"
)

type Request struct {
	fn func() int
	c  chan int
	e  chan error
}

func infiniteRequester(work chan<- Request) {
	c := make(chan int, 5)
	e := make(chan error, 5)
	for {
		randDuration := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(randDuration + 100*time.Millisecond)
		select {
		case work <- Request{randomWork, c, e}:
		case <-c:
			// here would be where you'd process the results
		case err := <-e:
			log.Printf("%v\n", err)
			work <- Request{randomWork, c, e}
		}
	}
}

func makeTwentyRequests(work chan<- Request) time.Duration {
	start := time.Now()
	c := make(chan int, 10)
	e := make(chan error, 5)
	for i := 0; i < 20; i++ {
		work <- Request{constantWork, c, e}
	}
	for i := 0; i < 20; i++ {
		<-c
	}
	return time.Since(start)
}

func randomWork() int {
	randDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(randDuration)
	if randDuration <= time.Millisecond*400 {
		panic("3 FAST 5 ME")
	}
	return 1
}

func constantWork() int {
	time.Sleep(200 * time.Millisecond)
	return 1
}