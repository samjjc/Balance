package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	compare := flag.Bool("compare", false, " compare load balancer vs single thread execution speed")
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
		distributedTime := Benchmark(work)
		synchronousTime := benchmarkSingleThread()
		fmt.Printf("Worker pool of %d workers took %s to complete\n", poolSize, distributedTime)
		fmt.Printf("Single thread took %s to complete\n", synchronousTime)
	case false:
		InfiniteRequester(work)
	}
}

func benchmarkSingleThread() time.Duration {
	singleWorker := make(chan Request, 20)

	go func() {
		for i := 0; i < 20; i++ {
			req := <-singleWorker
			req.result <- req.job()
		}
	}()
	return Benchmark(singleWorker)
}
