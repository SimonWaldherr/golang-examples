package main

import (
	"fmt"
	"runtime"
	"time"
)

var chan_int chan int
var stop_bool chan bool

func init() {
	chan_int = make(chan int, 2)
	stop_bool = make(chan bool, 3)
}

func maxprocs1() {
	/*
	 * Channels are useful for communications between multiple
	 * go routines, even when the routines are not runned in parralel.
	 *
	 */

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1")
		stop_bool <- true
	}()
	go func() {
		time.Sleep(1900 * time.Millisecond)
		fmt.Println("2")
		stop_bool <- true
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("3")
		stop_bool <- true
	}()

	<-stop_bool
	<-stop_bool
	<-stop_bool

	/*
	 * the previous code is compareable to the following bash code:
	 *
	 * function1 &
	 * stop1=$!
	 *
	 * function2 &
	 * stop2=$!
	 *
	 * function3 &
	 * stop3=$!
	 *
	 * wait $stop1
	 * wait $stop2
	 * wait $stop3
	 *
	 */
}

func waitforbuffer() {
	/*
	 *
	 * Starts a go routine which waits 2 seconds, prints the value
	 * from the channel buffer and waits again (over and over again).
	 * In the parrent function it tries to fill the channel with 3 integers,
	 * but after the first two values it has to wait until the buffer gets
	 * emptied. After 2 seconds, the first value gets loaded and printed,
	 * so the next value can be saved (but will never be printed, because
	 * the programm no longer have to wait for the channel and quit.
	 *
	 */

	go func(i int) {
		for {
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println(<-chan_int)
		}
	}(2)
	fmt.Println("start")
	chan_int <- 1
	fmt.Println("<-1")
	chan_int <- 2
	fmt.Println("<-2")
	chan_int <- 3
	fmt.Println("<-3")
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("starting maxprocs1()")
	maxprocs1()
	fmt.Println("starting waitforbuffer()")
	waitforbuffer()
	fmt.Println("finish")
}
