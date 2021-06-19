package main

// import required modules
import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("math/rand:")
	for i := 0; i < 10; i++ {
		fmt.Println(i, rand.Intn(127))
	}

	fmt.Println("crypto/rand:")
	b := make([]byte, 3)
	for i := 0; i < 10; i++ {
		crand.Read(b)
		number := uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16)
		fmt.Println(i, number)
	}
}
