package main

import (
	"fmt"
	"runtime"

	"github.com/samjjc/basicLoadBalancer"
)

func main() {
	poolSize := runtime.GOMAXPROCS(0)
	fmt.Printf("We got %d workers!!\n", poolSize)

	work := make(chan request.Request)
	done := make(chan *Worker)
	pool := NewPool(poolSize, done)
	balancer := Balancer{*pool, done}
	go balancer.balance(work)
	request.Requester(work)
}
