package main

// import required modules
import (
  "fmt";
  "math/rand";
  "time";
)

func main() {
  rand.Seed(time.Now().UnixNano());
  for i := 0; i < 10; i++ {
    fmt.Println(i, rand.Intn(127));
  }
}
