// Author: Mark Mukiiza (c00166672@setu.ie)
// Student No: C00166672
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, barrier chan bool, release chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	barrier <- true //signal main that this goroutine has arrived
	<-release       //block until main releases everyone
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool) //unbuffered: goroutines report arrival here
	release := make(chan bool) //unbuffered: main lets goroutines go here
	threadCount := 10
	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier, release)
	}
	//wait until all goroutines have arrived
	for range threadCount {
		<-barrier
	}
	//everyone has finished Part A, so release them all
	for range threadCount {
		release <- true
	}
	//No deadlock: every send has a matching receive
	//(threadCount arrivals received, threadCount releases sent)
	wg.Wait() //wait here until everyone (10 go routines) is done
}
