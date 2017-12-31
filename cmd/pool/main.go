package main

import (
	"fmt"
)

func main() {
	fmt.Println("Coming soon!!")

	work := make(chan Request)
	done := make(chan *Worker)
	pool := NewPool(5, done)
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
