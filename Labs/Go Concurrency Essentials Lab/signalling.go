//signalling.go Template Code
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

//Global variables shared between functions --A BAD IDEA

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")
		//wait here
		barrier <- true
		fmt.Println("StuffOne - PartB")
		wg.Done()
		return true
	}
	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")
		//wait here

		<-barrier
		fmt.Println("StuffTwo - PartB")
		wg.Done()
		return true
	}
	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait() //wait here until everyone (10 go routines) is done

}
