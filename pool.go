package main

import (
	"container/heap"
	"fmt"
)

//Pool is a worker pool
type Pool []*Worker

//NewPool creates a new worker pool, and initializes the workers
func NewPool(size int, done chan *Worker) *Pool {
	var pool Pool
	for i := 0; i < size; i++ {
		requests := make(chan Request, 300)
		worker := Worker{requests, 0, i}
		go worker.work(done)
		pool = append(pool, &worker)
	}
	heap.Init(&pool)
	return &pool
}

func (p Pool) String() string {
	s := "Pool: "
	for _, v := range p {
		s += fmt.Sprint(" ", v.pending)
		// s += fmt.Sprint(" ", len(v.requests))
	}
	return s
}

//implement heap interface
func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pool) Push(x interface{}) {
	*p = append(*p, x.(*Worker))
}

func (p *Pool) Pop() interface{} {
	old := *p
	last := len(old) - 1
	element := old[last]
	*p = old[:last]
	return element
}
