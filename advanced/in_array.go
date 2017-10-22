package main

import (
	"fmt"
	"reflect"
)

// This function will search element inside array with any type.
// Will return boolean and index for matched element.
// True and index more than 0 if element is exist.
// needle is element to search, haystack is slice of value to be search.
func InArray(needle interface{}, haystack interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func main() {
	// Array of string
	arrStrings := []string{"foo", "bar", "baz"}
	searchString := "bar"
	stringExist, stringIndex := InArray(searchString, arrStrings)
	fmt.Printf("The '%s' is %v inside arrays %v with index: %d\n", searchString, stringExist, arrStrings, stringIndex)

	// Array of int64 (or you can use int or int32 too)
	arrInt64 := []int64{2016, 2017, 2018, 2019}
	searchInt64 := int64(2016)
	int64Exist, int64Index := InArray(searchInt64, arrInt64)
	fmt.Printf("The '%d' is %v inside arrays %v with index: %d\n", searchInt64, int64Exist, arrInt64, int64Index)

	// Example for false searching
	int64NotExist := int64(2000)
	isElementExist, foundElementIndex := InArray(int64NotExist, arrInt64)
	fmt.Printf("The '%d' is %v inside arrays %v with index: %d\n", int64NotExist, isElementExist, arrInt64, foundElementIndex)

	// False searching with different type
	// Search string inside array of int64
	searchStringInt64 := "2018"
	isStringExistInInt64Arr, stringIndexIn64ArrElement := InArray(searchStringInt64, arrInt64)
	fmt.Printf("The '%s' (%v) is %v inside arrays %v with index: %d\n", searchStringInt64, reflect.TypeOf(searchStringInt64).Name(), isStringExistInInt64Arr, arrInt64, stringIndexIn64ArrElement)

	// Or, you can also using array of interface
	arrInterface := []interface{}{"username", 123, int64(10), false}
	searchElement := false
	ok, index := InArray(searchElement, arrInterface)
	fmt.Printf("The '%v' is %v inside arrays %v with index: %d\n", searchElement, ok, arrInterface, index)
}
