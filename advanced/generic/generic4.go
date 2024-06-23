package main

import (
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	if value, ok := intStack.Pop(); ok {
		fmt.Println("Pop from intStack:", value)
	} else {
		fmt.Println("intStack is empty")
	}
	if value, ok := intStack.Pop(); ok {
		fmt.Println("Pop from intStack:", value)
	} else {
		fmt.Println("intStack is empty")
	}

	stringStack := Stack[string]{}
	stringStack.Push("apple")
	stringStack.Push("banana")
	if value, ok := stringStack.Pop(); ok {
		fmt.Println("Pop from stringStack:", value)
	} else {
		fmt.Println("stringStack is empty")
	}
	if value, ok := stringStack.Pop(); ok {
		fmt.Println("Pop from stringStack:", value)
	} else {
		fmt.Println("stringStack is empty")
	}
}
