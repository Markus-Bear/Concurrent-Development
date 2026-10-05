//Barrier.go Template Code
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

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Mark Mukiiza (c00166672@setu.ie)
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// doStuff runs Part A then Part B for a number of rounds, using a two part barrier so no
// routine starts Part B until all have done Part A, and no routine starts the next round
// until all have done Part B
func doStuff(goNum int, arrived *atomic.Int32, max int32, wg *sync.WaitGroup, firstGate chan bool, secondGate chan bool, rounds int) bool {
	for range rounds { //loop to show the barrier is reusable
		time.Sleep(time.Second)
		fmt.Println("Part A", goNum)
		//wait here until everyone has completed part A
		if arrived.Add(1) == max { //Add returns the new value, so only one routine sees max
			for range max - 1 { //last to arrive lets the others through
				firstGate <- true
			}
		} else {
			<-firstGate
		}
		fmt.Println("PartB", goNum)
		//wait here until everyone has finished part B
		//stops a fast routine looping round and using up a part A signal
		if arrived.Add(-1) == 0 { //last to leave lets the others through
			for range max - 1 {
				secondGate <- true
			}
		} else {
			<-secondGate
		}
	}
	wg.Done()
	return true
} //end-doStuff

// main sets up the shared counter and gates, starts the routines
// and waits for all of them to finish every round
func main() {
	totalRoutines := 10
	rounds := 3
	var arrived atomic.Int32 //atomic counter replaces the mutex
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	firstGate := make(chan bool)   //gate between part A and part B
	secondGate := make(chan bool)  //gate between part B and the next round/routine
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, int32(totalRoutines), &wg, firstGate, secondGate, rounds)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
