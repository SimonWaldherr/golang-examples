// Description: Binary search algorithm implementation in Go
// Tags: binary, search, algorithm, slice, array, sort, sorted, sorted array, sorted slice, sorted
package main

import "fmt"

func binarySearch(element int, arr []int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] < element {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(arr) || arr[low] != element {
		return false
	}

	return true
}

func main() {
	arr := []int{1, 4, 5, 7, 9, 10, 35, 56, 79, 80, 100, 200, 210, 250}
	fmt.Println(binarySearch(9, arr))
}
