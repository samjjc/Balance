package main

import (
	"container/heap"
	"fmt"

	"github.com/samjjc/basicLoadBalancer"
)

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan request.Request) {
	for {
		select {
		case req := <-work: // received a Request
			b.dispatch(req)
		case w := <-b.done: // a worker has finished
			b.completed(w)
		}
	}
}

// Send Request to worker
func (b *Balancer) dispatch(req request.Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.pending++
	heap.Push(&b.pool, w)
	fmt.Println(b.pool, "| DISPATCHED")
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
	fmt.Println(b.pool, "| COMPLETED")
}
