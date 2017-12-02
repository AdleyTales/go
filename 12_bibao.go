package main

import "fmt"

func pp() func() int {
  i := 0
  return func () int {
    i += 1
    return i
  }
}

func main()  {
  n := pp()

  // 检测 i 的值
  fmt.Println(n()) //1
  fmt.Println(n()) //2
  fmt.Println(n()) //3
  fmt.Println(n()) //4

}
