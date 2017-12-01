package main

import "fmt"

func main()  {

  s := make([]string, 3)
  fmt.Println(s) // [    ]
  fmt.Println(len(s)) //3

  s[0] = "ssssss"
  s[1] = "cccccccccc"
  s[2] = "bbbbbbbbb"
  fmt.Println(s)
  fmt.Println(s[1])

  l := s[1:]
  fmt.Println(l)

  // 切片 在go中是一种类数组的数据类型 同python中的切片用法



}
