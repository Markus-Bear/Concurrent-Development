//dinPhil.go Template Code
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
// Created on 21/10/2024
// Modified by: Mark Mukiiza (c00166672@setu.ie)
// Description:
// Dining Philosophers Template Code
// Issues:
// MISSING:
// 1. Readme (Fixed)
// 2. Full licence info. (Fixed)
// 3. Comments (Fixed)
// 4. It can Deadlock! (Fixed)
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// think for a random time (0-4 seconds)
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

// eat for a random time (0-4 seconds)
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// sit at the table, then pick up left fork then right fork
func getForks(index int, forks map[int]chan bool, footman chan bool) {
	//FOOTMAN: only 4 can sit at once, so at least one can always get both forks and deadlock can't happen
	footman <- true //take a seat, blocks if 4 are already seated
	//FORKS: sending fills the fork's channel, which picks it up, and blocks if a neighbour holds it
	forks[index] <- true
	forks[(index+1)%5] <- true
}

// put down both forks, then leave the table
func putForks(index int, forks map[int]chan bool, footman chan bool) {
	//FORKS: receiving empties the fork's channel, which puts it down for a neighbour
	<-forks[index]
	<-forks[(index+1)%5]
	<-footman //free a seat
}

// each philosopher thinks, gets forks, eats and puts forks down, forever
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, footman chan bool) {
	for {
		think(index)
		getForks(index, forks, footman)
		eat(index)
		putForks(index, forks, footman)
	}
	wg.Done()
}

// main sets up the forks and the footman, then starts the philosophers,
// who run until the program is stopped
func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	//FORKS: each fork is a channel with room for one item
	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1) //buffer of 1: only one philosopher can hold a fork
	} //set up forks
	footman := make(chan bool, philCount-1) //only 4 can sit at once, so no deadlock
	for N := range philCount {
		go doPhilStuff(N, &wg, forks, footman)
	} //start philosophers
	wg.Wait() //wait here until everyone (5 go routines) is done

} //main
