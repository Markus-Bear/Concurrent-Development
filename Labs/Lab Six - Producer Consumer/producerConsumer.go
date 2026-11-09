//producerConsumer.go
//Copyright (C) 2026 Mark Mukiiza

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Mark Mukiiza (c00166672@setu.ie)
// Created on 9/11/2026
// Description:
// Producer Consumer implemented as a worker pool.
// One producer adds tasks to a finite buffer (buffered channel)
// and a pool of workers consumes them.
// Issues:
// None I hope
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Task is one item of work added to the buffer by the producer
type Task struct {
	ID  int
	Val int
}

// Result is what a worker sends back after consuming a task
type Result struct {
	TaskID int
	Out    int
	Err    error
}

// consume processes a task by squaring its value
func consume(t Task) int {
	time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond) //simulate work
	return t.Val * t.Val
}

// worker takes tasks from the buffer until it is closed and empty or done is closed
func worker(id int, done <-chan struct{}, tasks <-chan Task, results chan<- Result) {
	//CONSUMER: each worker repeatedly takes a task from the buffer, blocking while it is empty
	for {
		select {
		case <-done:
			return
		case t, ok := <-tasks:
			if !ok { //producer has finished and the buffer is empty
				return
			}
			out := consume(t)
			select { //never block forever sending a result nobody will read
			case <-done:
				return
			case results <- Result{TaskID: t.ID, Out: out, Err: nil}:
			}
		}
	}
}

// startWorkerPool starts n workers and returns the channel their results arrive on
func startWorkerPool(n int, done <-chan struct{}, tasks <-chan Task) (<-chan Result, *sync.WaitGroup) {
	results := make(chan Result)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		// each worker tells the WaitGroup when it has stopped
		go func(id int) {
			defer wg.Done()
			worker(id, done, tasks, results)
		}(i + 1)
	}
	go func() { //close results once every worker has stopped
		wg.Wait()
		close(results)
	}()
	return results, &wg
}

// main creates the buffer and the workers, starts the producer that fills
// the buffer, then prints each result as it arrives
func main() {
	done := make(chan struct{})
	defer close(done) //stop any routines still running when main exits
	//BUFFER: shared by the producer and consumers, the channel does the job of mutex, items and spaces
	tasks := make(chan Task, 8) //buffer size 8: producer blocks when it is full
	const workers = 4
	results, _ := startWorkerPool(workers, done, tasks)

	//PRODUCER: adds tasks to the buffer, blocking while it is full
	go func() { //the producer
		defer close(tasks) //tells the workers no more tasks are coming
		for i := 1; i <= 20; i++ {
			select {
			case <-done:
				return
			case tasks <- Task{ID: i, Val: i}:
			}
		}
	}()

	for r := range results { //ends when results is closed
		if r.Err != nil {
			fmt.Printf("task %d error: %v\n", r.TaskID, r.Err)
			continue
		}
		fmt.Printf("task %02d -> %d\n", r.TaskID, r.Out)
	}
}
