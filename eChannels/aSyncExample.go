package main

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func unsafeMemoryAccess() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go increment(wg)
	}

	wg.Wait()
	fmt.Println(x)
}
