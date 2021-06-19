package main

// import required modules
import (
	"fmt"
)

// declare variables and define array content
var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}
var intarray = []int{1, 2, 4, 8, 16}
var mapone = map[int]string{}
var maptwo = map[string]interface{}{}

func main() {

	// do this five times
	for i := 0; i != 5; i++ {

		// print the $th value of the intarray and the strarray
		fmt.Println(intarray[i], "\t", strarray[i])

		mapone[intarray[i]] = strarray[i]
		maptwo[strarray[i]] = mapone
	}
	fmt.Println(mapone)
	fmt.Println(maptwo)
}
