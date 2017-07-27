package main


import (
  "fmt"
  "stringutil"
  "time"
  )

func main() {
	fmt.Printf("Hello, world.\n")
  fmt.Printf(stringutil.Reverse("!oG, olleH"))
  fmt.Printf("\nThe time is ", time.Now())
  a, b := swap("Swap", "ME")
  fmt.Printf(a, b)
}

func swap(x, y string) (string, string) {
  return y, x
}
