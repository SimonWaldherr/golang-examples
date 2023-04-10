// Description: Find lower and upper bound of a target in a sorted array
// Tags: binary, search, algorithm, slice, array, sort, sorted, sorted array, sorted slice, sorted
package main

import "fmt"

func lowerBound(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func upperBound(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	arr := []int{1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 6, 6, 7, 7, 7}
	fmt.Printf("%s%d\n", "Lower bound of 4: ", lowerBound(arr, 4))
	fmt.Printf("%s%d\n", "Upper bound of 4: ", upperBound(arr, 4))
}
