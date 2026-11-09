# Concurrent Development

Lab work for the Concurrent Development module in the fourth year of the BSc (Hons) in Software Development at SETU Carlow. The labs are written in **Go** and explore how to build programs that do many things at once, safely.

## Repository layout

```
Concurrent-Development/
├── Labs/        Lab exercises, one per folder
├── .gitignore
├── LICENSE
└── README.md
```

## Running the labs

1. Install [Go](https://go.dev/dl/).
2. Clone the repository:

   ```bash
   git clone https://github.com/Markus-Bear/Concurrent-Development.git
   cd Concurrent-Development/Labs
   ```

3. Change into a lab folder and run a program by its file name. Some lab folders hold several separate programs, each with its own `main`, so `go run .` won't work. Each lab's README lists its programs:

   ```bash
   cd "Lab Six - Producer Consumer"
   go run producerConsumer.go
   ```

4. To check a program for data races, run it with Go's race detector:

   ```bash
   go run -race producerConsumer.go
   ```

## Author

Mark Mukiiza, Software Development student at SETU Carlow. [LinkedIn](https://www.linkedin.com/in/mukiiza-mark)

## License

Released under the GPL-3.0 license. See [LICENSE](LICENSE) for details.
