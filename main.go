package main

import (
	"flag"
	"fmt"
	"runtime"
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
	InfiniteRequester(work, randomFailingJob)
}
