package main

import (
	"log"
	"math/rand"
	"time"
)

//Request job is executed by the worker and the result is sent on the result channel
type Request struct {
	job    func() int
	result chan int
	err    chan error
}

//InfiniteRequester sends custom requests to the work channel every 0.1 to 0.4 seconds
func InfiniteRequester(work chan<- Request, job func() int) {
	c := make(chan int, 25)
	e := make(chan error, 5)
	for {
		randDuration := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(randDuration + 100*time.Millisecond)
		select {
		case work <- Request{job, c, e}:
		case <-c:
			// here would be where you'd process the results
		case err := <-e:
			log.Printf("%v\n", err)
			work <- Request{job, c, e}
		}
	}
}

func randomFailingJob() int {
	randDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(randDuration)
	if randDuration <= time.Millisecond*400 { //random failure
		panic("3 FAST 5 ME")
	}
	return 1
}
