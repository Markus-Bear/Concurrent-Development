//mutex.go Template Code
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
)

// Global variables shared between functions --A BAD IDEA
//var wg sync.WaitGroup
//var total int64

func adds(n int, total *int64, theLock *sync.Mutex, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		*total++ //only one routine can be here at a time
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	//theLock will be passed by reference between go routines
	//better than using a global variable
	var theLock sync.Mutex
	var wg sync.WaitGroup //local and also passed by reference now.
	var total int64       //local and also passed by reference now.

	total = 0
	//the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println(i)
		go adds(1000, &total, &theLock, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
