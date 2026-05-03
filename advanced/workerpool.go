// Description: Worker pool pattern - distribute tasks across a fixed pool of goroutines
// Tags: worker pool, goroutine, channel, sync, WaitGroup, concurrency, pattern
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job holds the input data for one unit of work.
type Job struct {
	ID    int
	Value int
}

// Result holds the outcome produced by a worker.
type Result struct {
	JobID  int
	Output int
}

// worker processes jobs from the jobs channel and sends results to results.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate variable-length work.
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		output := job.Value * job.Value // square the input
		fmt.Printf("worker %d processed job %d (input=%d, output=%d)\n",
			id, job.ID, job.Value, output)
		results <- Result{JobID: job.ID, Output: output}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start the worker pool.
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs into the channel, then close it so workers stop ranging.
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs)

	// Wait for all workers to finish, then close the results channel.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results.
	total := 0
	for r := range results {
		total += r.Output
	}
	fmt.Printf("\nsum of squares 1..%d = %d\n", numJobs, total)
}
