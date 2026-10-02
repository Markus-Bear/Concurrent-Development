# Lab 3 - Simple Barrier

Two barrier solutions for 10 goroutines. Each goroutine prints Part A,
waits until all of them have finished Part A, then prints Part B.

| File | Description |
|------|-------------|
| `barrier.go` | Barrier using a Mutex and a Semaphore |
| `rendezvous.go` | Barrier using channels, with main as coordinator |

## Running

Both files have their own `main`, so run them one at a time from this folder:

```bash
go run barrier.go
```

```bash
go run rendezvous.go
```

`barrier.go` uses `golang.org/x/sync/semaphore`. If Go reports a
missing package, run `go mod tidy` first.

Requires Go 1.22 or later.
