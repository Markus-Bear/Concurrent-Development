# Lab Five - Dining Philosophers

5 philosophers sit at a table with 5 forks, one between each pair.
Each needs both forks next to them to eat. The template could deadlock
if every philosopher picked up their left fork at the same time.

| File | Description |
|------|-------------|
| `dinPhil.go` | Dining philosophers using channels as forks, with a footman to prevent deadlock |

## Solution

A footman (buffered channel of size 4) only lets 4 philosophers sit at
the table at once. With 4 philosophers and 5 forks, at least one can
always get both forks, so deadlock is impossible.

## Running

```bash
go run dinPhil.go
```

The philosophers loop forever. Press `Ctrl+C` to stop.

Requires Go 1.22 or later.