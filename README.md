# Basic Load Balancer

Trivial Load balancers' performance compared (so far only a worker pool)

## Worker Pool

Based on Rob Pike's `concurrency is not parallelism`, a worker pool, or thread pool, maintains multiple workers waiting for tasks to be allocated for concurrent execution by the balancer
