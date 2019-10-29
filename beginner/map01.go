package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["a"] = 1
    m["b"] = 2

    a := m["a"]
    fmt.Println("a: ", a)

    delete(m, "b")
    b, ok := m["b"]
    fmt.Println("b: ", b)
    fmt.Println("ok? ", ok)
}
