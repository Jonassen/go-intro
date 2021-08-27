package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}
var x int = 0

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func unsafeMemoryAccess() {
	wg := &sync.WaitGroup{}

	//ingentingHer := go increment(wg)

	wg.Add(1)
	go increment(wg)

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go increment(wg)
	}

	wg.Wait()
	fmt.Println(x)
}
