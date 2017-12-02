package main

import "fmt"

func main()  {

  arr := []int{23,43,54,32,45}

  sum := 0
  for _,num := range arr {
    sum += num
  }

  fmt.Println("总数是：",sum)

  fmt.Println("=====================")

  for i,num := range arr {
    if num == 54 {
      fmt.Println("54 的索引下标是：",i)
    }
  }

  maps := map[string]int {"x1": 32, "y2": 32}
  for k,v := range maps {
    fmt.Println(k," ",v)
  }

}
