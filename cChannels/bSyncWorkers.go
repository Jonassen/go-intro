package main

import (
	"fmt"
	"sync"
)

type job struct {
	id int
}

func startWorker(id int, jobs chan job, cancel chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case job := <-jobs:
			fmt.Printf("Worker: %v doing job: %v\n", id, job.id)
		case <-cancel:
			fmt.Printf("Worker: %v shutting down\n", id)
			wg.Done()
			return
		}
	}
}

func DoWork() {
	numWorkers := 3

	wg := &sync.WaitGroup{}

	jobChan := make(chan job)
	cancelChan := make(chan struct{})

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go startWorker(i, jobChan, cancelChan, wg)
	}

	for i := 0; i < 15; i++ {
		jobChan <- job{id: i}
	}

	for i := 0; i < numWorkers; i++ {
		cancelChan <- struct{}{}
	}
	wg.Wait()
}
