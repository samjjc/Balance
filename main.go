package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	compare := flag.Bool("compare", false, " compare laod balancer with single thread")
	flag.Parse()

	poolSize := runtime.GOMAXPROCS(0)
	fmt.Printf("We got %d workers!!\n", poolSize)

	work := make(chan Request)
	done := make(chan *Worker)
	pool := NewPool(poolSize, done)
	balancer := Balancer{*pool, done}
	go balancer.balance(work)

	switch *compare {
	case true:
		distributedTime := makeTwentyRequests(work)
		synchronousTime := singleThreadWork()
		fmt.Println("Worker pool took", distributedTime, "to complete")
		fmt.Println("Single thread took", synchronousTime, "to complete")
	case false:
		infiniteRequester(work)
	}

}

func singleThreadWork() time.Duration {
	singleWorker := make(chan Request, 20)

	go func() {
		for i := 0; i < 20; i++ {
			req := <-singleWorker
			req.c <- req.fn()
		}
	}()
	return makeTwentyRequests(singleWorker)
}
