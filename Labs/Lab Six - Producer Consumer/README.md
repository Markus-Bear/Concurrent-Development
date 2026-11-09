# Lab Six - Producer Consumer

Producer Consumer implemented as a worker pool. One producer adds 20
tasks to a finite buffer and 4 workers consume them, squaring each value.

| File | Description |
|------|-------------|
| `producerConsumer.go` | Worker pool using a buffered channel as the finite buffer |

## Solution

The buffered channel (size 8) replaces the mutex and the `items` and
`spaces` semaphores: receives block when it is empty and sends block
when it is full. Closing `tasks` tells the workers no more work is
coming, and closing `done` stops every routine early.

## Running

```bash
go run producerConsumer.go
```

Results print in a different order each run.

Requires Go 1.22 or later.
