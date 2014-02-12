package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(time.Now().Format(time.RFC1123Z))
	fmt.Println(time.Now().Format("20060102150405"))
}
