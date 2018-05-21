package main

import (
	"fmt"
	timespan "github.com/senseyeio/spaniel"
	"time"
)

type dur struct {
	from time.Time
	to   time.Time
}

var times []dur
var input timespan.Spans

func init() {
	times = []dur{{from: time.Date(2018, 1, 30, 0, 0, 0, 0, time.UTC), to: time.Date(2018, 1, 30, 1, 0, 0, 0, time.UTC)},
		{from: time.Date(2018, 1, 30, 0, 30, 0, 0, time.UTC), to: time.Date(2018, 1, 30, 1, 30, 0, 0, time.UTC)},
		{from: time.Date(2018, 1, 30, 1, 31, 0, 0, time.UTC), to: time.Date(2018, 1, 30, 1, 35, 0, 0, time.UTC)},
		{from: time.Date(2018, 1, 30, 1, 33, 0, 0, time.UTC), to: time.Date(2018, 1, 30, 1, 34, 0, 0, time.UTC)},
	}
	for t := range times {
		input = append(input, timespan.New(times[t].from, times[t].to))
	}
}

func Union() {
	union := input.Union()

	fmt.Println("Union")
	for u := range union {
		fmt.Println(union[u].Start(), "->", union[u].End(), ": ", union[u].End().Sub(union[u].Start()))
	}

	// Output:
	// Union
	// 2018-01-30 00:00:00 +0000 UTC -> 2018-01-30 01:30:00 +0000 UTC :  1h30m0s
	// 2018-01-30 01:31:00 +0000 UTC -> 2018-01-30 01:35:00 +0000 UTC :  4m0s
}

func Intersection() {
	intersection := input.Intersection()

	fmt.Println("Intersection")
	for i := range intersection {
		fmt.Println(intersection[i].Start(), "->", intersection[i].End(), ": ", intersection[i].End().Sub(intersection[i].Start()))
	}

	// Output:
	// Intersection
	// 2018-01-30 00:30:00 +0000 UTC -> 2018-01-30 01:00:00 +0000 UTC :  30m0s
	// 2018-01-30 01:33:00 +0000 UTC -> 2018-01-30 01:34:00 +0000 UTC :  1m0s
}

func main() {
	Union()
	Intersection()
}
