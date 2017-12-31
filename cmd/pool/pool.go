package main

import (
	"container/heap"
	"fmt"
)

type Pool []*Worker

func NewPool(size int, done chan *Worker) *Pool {
	var pool Pool
	for i := 0; i < size; i++ {
		requests := make(chan Request, 30)
		worker := Worker{requests, 0, i}
		go worker.Work(done)
		pool = append(pool, &worker)
	}
	heap.Init(&pool)
	return &pool
}

func (p Pool) String() (s string) {
	s = "Pool: "
	for _, v := range p {
		s += fmt.Sprint(" ", v.pending)
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
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
