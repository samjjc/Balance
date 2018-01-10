# Basic Load Balancer

A trivial load balancer that spreads tasks across a pool of workers, also know as a worker pool

## Worker Pool

Based on Rob Pike's `concurrency is not parallelism`, a worker pool, or thread pool, maintains multiple workers waiting for tasks to be allocated for concurrent execution by the balancer

## Benchmarking tests

run `go test -bench=.`
