package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	poolSize := flag.Int("poolSize", runtime.GOMAXPROCS(0), " number of workers in the worker pool")
	flag.Parse()

	fmt.Printf("We got %d workers!!\n", *poolSize)

	work := make(chan Request)
	done := make(chan *Worker)
	pool := NewPool(*poolSize, done)
	balancer := Balancer{*pool, done, true}
	go balancer.balance(work)
	InfiniteRequester(work)
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
