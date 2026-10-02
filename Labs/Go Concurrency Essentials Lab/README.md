# Go Concurrency Essentials Lab

Basic Go concurrency tools: goroutines, WaitGroups, atomics, mutexes,
semaphores and channels.

| File | Description |
|------|-------------|
| `atomic.go` | 10 goroutines update a shared counter using atomic operations |
| `mutex.go` | 10 goroutines update a shared counter protected by a Mutex |
| `semaphore.go` | Semaphore built from a buffered channel, at most 5 tasks at a time |
| `sem-ex.go` | Worker pool limited by a weighted semaphore |
| `signalling.go` | Two goroutines synchronise using an unbuffered channel |

## Running

Each file has its own `main`, so run them one at a time from this folder:

```bash
go run atomic.go
```

```bash
go run mutex.go
```

```bash
go run semaphore.go
```

```bash
go run sem-ex.go
```

```bash
go run signalling.go
```

`sem-ex.go` uses `golang.org/x/sync/semaphore`. If Go reports a
missing package, run `go mod tidy` first.

Requires Go 1.22 or later.