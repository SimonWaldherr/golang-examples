package main

import (
  "fmt"
  "strings"
)

var variable string
var strarray []string

func main() {
  variable = "Lorem Ipsum Dolor Sit Amet"
  fmt.Println(variable)
  strarray = strings.Split(variable, " ")
  for i := 0; i < len(strarray); i++ {
    fmt.Println(strarray[i])
  }
}
