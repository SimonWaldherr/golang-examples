// Description: Start and wait for concurrent tasks with sync.WaitGroup.Go (Go 1.25)
// Tags: concurrency, goroutine, sync, WaitGroup.Go, Go 1.25
package main

import (
	"fmt"
	"sync"
)

func square(value int) int {
	return value * value
}

func main() {
	inputs := []int{2, 3, 5, 8, 13}
	results := make([]int, len(inputs))

	var tasks sync.WaitGroup
	for index, value := range inputs {
		tasks.Go(func() {
			// Each loop iteration has its own index and value since Go 1.22.
			// Functions passed to WaitGroup.Go must not panic.
			results[index] = square(value)
		})
	}
	tasks.Wait()

	fmt.Println(results)
}
