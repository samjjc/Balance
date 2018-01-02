package main

import (
	"log"
	"math/rand"
	"time"
)

type Request struct {
	job    func() int
	result chan int
	err    chan error
}

func infiniteRequester(work chan<- Request) {
	c := make(chan int, 5)
	e := make(chan error, 5)
	for {
		randDuration := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(randDuration + 100*time.Millisecond)
		select {
		case work <- Request{randomJob, c, e}:
		case <-c:
			// here would be where you'd process the results
		case err := <-e:
			log.Printf("%v\n", err)
			work <- Request{randomJob, c, e}
		}
	}
}

func randomJob() int {
	randDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(randDuration)
	if randDuration <= time.Millisecond*400 {
		panic("3 FAST 5 ME")
	}
	return 1
}

// Benchmark sends 20 requests and returns how long it takes for them to complete
func Benchmark(work chan<- Request) time.Duration {
	start := time.Now()
	c := make(chan int, 5)
	e := make(chan error, 5)
	for i := 0; i < 20; i++ {
		work <- Request{constantJob, c, e}
	}
	for i := 0; i < 20; i++ {
		<-c
	}
	return time.Since(start)
}

func constantJob() int {
	time.Sleep(200 * time.Millisecond)
	return 1
}
