package main

import "fmt"

func main()  {

  var a [5]int
  fmt.Println(a) // [0 0 0 0 0]

  a[4] = 200000
  fmt.Println(a) // [0 0 0 0 200000]

  fmt.Println(len(a)) // 5

  // =============
  b := [5]int{1,2,3,4,5}
  fmt.Println(b)
  b[3] = 233
  fmt.Println(b) // [1 2 3 233 5]

}
