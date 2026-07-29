// Description: Capture recent runtime activity with runtime/trace.FlightRecorder (Go 1.25)
// Tags: runtime, trace, flight recorder, diagnostics, profiling, Go 1.25
package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	recorder := trace.NewFlightRecorder(trace.FlightRecorderConfig{
		MinAge:   2 * time.Second,
		MaxBytes: 1 << 20,
	})
	if err := recorder.Start(); err != nil {
		panic(err)
	}
	defer recorder.Stop()

	var checksum atomic.Uint64
	trace.WithRegion(context.Background(), "example-workload", func() {
		var workers sync.WaitGroup
		for worker := range 4 {
			workers.Go(func() {
				for value := range 100_000 {
					checksum.Add(uint64(value + worker))
					if value%10_000 == 0 {
						runtime.Gosched()
					}
				}
			})
		}
		workers.Wait()
	})

	output, err := os.Create("flight.trace")
	if err != nil {
		panic(err)
	}

	written, writeErr := recorder.WriteTo(output)
	closeErr := output.Close()
	if writeErr != nil {
		panic(writeErr)
	}
	if closeErr != nil {
		panic(closeErr)
	}

	fmt.Printf("checksum: %d\n", checksum.Load())
	fmt.Printf("wrote %d bytes to flight.trace\n", written)
	fmt.Println("inspect it with: go tool trace flight.trace")
}
