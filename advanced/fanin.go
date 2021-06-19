package main

import (
	"fmt"
	"sync"
)

// returns a read only channel
func generator(start, end int) <-chan int {
	c := make(chan int) // ubuffered channel

	// Fire a goroutine to send values on the channel
	go func() {
		for i := start; i < end; i++ {
			c <- i // This blocks untill there is a reader for the chan
		}
		// close the channel when done; otherwise it leaks resources
		close(c)
	}()

	return c
}

/* The fan in pattern is an important pattern which combines
mulitple channels, returns a single channel from those channels
*/

func fanIn(chans ...<-chan int) chan int {
	var wg sync.WaitGroup

	c := make(chan int)

	// Closure to send values from a channel
	output := func(ch <-chan int) {
		for n := range ch {
			c <- n
		}
		wg.Done()
	}

	wg.Add(len(chans))

	// send values on c via differnt goroutines
	for _, ch := range chans {
		go output(ch)
	}

	// wait for all goroutines to finish before closing the channel
	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

func main() {
	s1 := generator(1, 10)
	s2 := generator(20, 30)
	s3 := generator(40, 50)
	s4 := generator(60, 70)

	// merge all the channels into one
	mergerd := fanIn(s1, s2, s3, s4)

	for n := range mergerd { // range loop terminates once the chan is closed, otherwise it blocks if there is no value
		fmt.Println(n)
	}
}
