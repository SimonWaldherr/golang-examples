package main

// This example shows how easy it is to implement a semaphore using channels.
// A semaphore can be used to limit the number of concurrent go routines executing a task.

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConcurrent = 2

	const totalTasks = 10

	semaphore := make(chan struct{}, maxConcurrent)

	wg := sync.WaitGroup{}

	for i := range make([]int, totalTasks) {
		// blocks until semaphore is released
		semaphore <- struct{}{}

		taskNumber := i
		// executes task async
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				// release semaphore
				<-semaphore
			}()

			fmt.Println("executing task: ", taskNumber)
			time.Sleep(time.Second)
		}()
	}

	wg.Wait()

	fmt.Println("done")
}
