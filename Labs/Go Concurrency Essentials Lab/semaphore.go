//semaphore.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

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

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Mark Mukiiza (c00166672@setu.ie)
// --------------------------------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release
type Semaphore struct {
	theCounter chan struct{}
}

func InitSemaphore(n int) *Semaphore {
	return &Semaphore{make(chan struct{}, n)} //buffer size n = max routines allowed in at once
}
func Acquire(sem *Semaphore) {
	sem.theCounter <- struct{}{} //take a slot, blocks when all n are taken
}

func Release(sem *Semaphore) {
	<-sem.theCounter //free a slot so a waiting routine can continue
}

func main() {
	//maxGoroutines := 5
	//semaphore := make(chan struct{}, maxGoroutines)

	semaphore := InitSemaphore(5)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//semaphore <- struct{}{}
			//defer func() { <-semaphore }()
			Acquire(semaphore)
			defer Release(semaphore)

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
