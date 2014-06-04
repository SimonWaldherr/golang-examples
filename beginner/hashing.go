package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	md5 := md5.New()
	io.WriteString(md5, "Lorem Ipsum dolor sit Amet")
	fmt.Printf("%x\n", md5.Sum(nil))
}
