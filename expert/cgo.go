package main

/*
#include <stdio.h>
#include <stdlib.h>

void print(char* s) {
	printf("%s", s);
}

long factorial(int n) {
	if (n == 0) {
		return 1;
	} else {
		return (n * factorial(n - 1));
	}
}

long gcd(long a, long b) {
	if (b == 0) {
		return a;
	} else {
		return gcd(b, a % b);
	}
}

void prime(int n) {
	int i = 0, count, c;
	for (count = 2; count <= n + 1; ) {
		for (c = 2; c <= i - 1; c++ ) {
			if (i % c == 0) {
				break;
			}
		}
		if (c == i) {
			printf("%d ", i);
			count++;
		}
		i++;
	}
	print("\n");
}

int fibonacci(int n) {
	if (n == 0) {
		return 0;
	} else if (n == 1) {
		return 1;
	} else {
		return (fibonacci(n - 1) + fibonacci(n - 2));
	}
}

*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print("C.print: ")
	cs := C.CString("Hello from cgo\n")
	C.print(cs)
	C.free(unsafe.Pointer(cs))

	fmt.Print("C.factorial: ")
	fmt.Println(C.factorial(5))

	fmt.Print("C.gcd: ")
	fmt.Println(C.gcd(15, 230))

	fmt.Print("C.prime: ")
	C.prime(6)

	fmt.Print("C.fibonacci: ")
	fmt.Println(C.fibonacci(5))
}
