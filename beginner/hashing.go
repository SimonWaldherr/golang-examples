package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	md5 := md5.New()
	sha256 := sha256.New()
	sha512 := sha512.New()
	io.WriteString(md5, "Lorem Ipsum dolor sit Amet")
	sha256.Write([]byte("Lorem Ipsum dolor sit Amet"))
	sha512.Write([]byte("Lorem Ipsum dolor sit Amet"))
	fmt.Printf("md5:\t%x\n", md5.Sum(nil))
	fmt.Printf("sha256:\t%x\n", sha256.Sum(nil))
	fmt.Printf("sha512:\t%x\n", sha512.Sum(nil))
}
