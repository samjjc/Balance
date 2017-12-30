package main

type Request struct {
	fn func() int
	c  chan int
}
