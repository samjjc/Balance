package main

import (
	"fmt"
	"runtime"
)

func main() {
	poolSize := runtime.GOMAXPROCS(0)
	fmt.Printf("We got %d workers!!\n", poolSize)

	work := make(chan Request)
	done := make(chan *Worker)
	pool := NewPool(poolSize, done)
	balancer := Balancer{*pool, done}
	go balancer.balance(work)
	// go func() {
	// 	for {
	// 		fmt.Println(pool)
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	requester(work)
}
