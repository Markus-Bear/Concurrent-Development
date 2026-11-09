//atomic.go Template Code
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
	"sync/atomic"
)

// Global variables shared between functions --A BAD IDEA
//var wg sync.WaitGroup

// addsAtomic adds 1 to the shared total n times using an atomic add
func addsAtomic(n int, total *atomic.Int64, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		//ATOMIC: the update happens as one step, so routines can't lose each other's updates
		total.Add(1) //atomic add so no lock needed.
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

// main starts 10 routines that update the total, waits for them all,
// then prints the total
func main() {
	var wg sync.WaitGroup //local and passed by reference. No longer global.
	var total atomic.Int64

	//for loop using range option
	for i := range 10 {
		//the waitgroup is used as a barrier
		// init it to number of go routines
		wg.Add(1)
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total, &wg)
	}
	wg.Wait() //wait here until everyone is done
	fmt.Println(total.Load())

}
