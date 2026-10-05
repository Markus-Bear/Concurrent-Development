# Lab Four - Reusable Barrier

Barrier solutions that build on the Barrier Lab, including a barrier that can be
used again and again in a loop.

| File | Description |
|------|-------------|
| `barrier2.go` | Reusable two phase barrier using an atomic counter and two channels, run for 3 rounds |
| `barrierStruct.go` | Barrier as its own data type with a `wait()` method |

## Running

Both files have their own `main`, so run them one at a time from this folder:

```bash
go run barrier2.go
```

```bash
go run barrierStruct.go
```

Requires Go 1.22 or later.