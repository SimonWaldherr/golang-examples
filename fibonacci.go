package main

// import required modules
import (
  "flag"
  "strconv"
  "fmt"
)

var x int64

func main() {
  flag.Parse()
  s := flag.Arg(0)
  x, err := strconv.ParseInt(s, 10, 0)

  if err != nil {
    fmt.Println(err)
    x = 10
  }

  fibonacci(x)
}

func fibonacci (n int64) {
  var a int64 = 0
  var b int64 = 1
  var i int64
  var sum int64
  for i = 0; i < n; i++ {
    fmt.Println(a)
    sum = a + b
    a = b
    b = sum
  }
}