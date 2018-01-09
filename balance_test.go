package main

import "testing"

func BenchmarkFastBalancer(b *testing.B) {
	normalBalancer(b, fastJob)
}

func BenchmarkFastSingleThread(b *testing.B) {
	workSingleThread(b, fastJob)
}

func BenchmarkSlowBalancer(b *testing.B) {
	normalBalancer(b, constantJob)
}

func BenchmarkSlowSingleThread(b *testing.B) {
	workSingleThread(b, constantJob)
}

func normalBalancer(b *testing.B, job func() int) {
	work := make(chan Request)
	done := make(chan *Worker)
	pool := NewPool(4, done)
	balancer := Balancer{*pool, done, false}
	b.ResetTimer()
	go balancer.balance(work)
	for n := 0; n < b.N; n++ {
		FiniteRequests(work, job, 20)
	}
}

func workSingleThread(b *testing.B, job func() int) {
	jobsRepeats := 20
	singleWorker := make(chan Request, 20)
	for n := 0; n < b.N; n++ {
		go func() {
			for i := 0; i < jobsRepeats; i++ {
				req := <-singleWorker
				req.result <- req.job()
			}
		}()
		FiniteRequests(singleWorker, job, jobsRepeats)
	}
}

func FiniteRequests(work chan<- Request, job func() int, size int) {
	c := make(chan int, 5)
	e := make(chan error, 5)
	for i := 0; i < size; i++ {
		work <- Request{job, c, e}
	}
	for i := 0; i < size; i++ {
		<-c
	}
}
