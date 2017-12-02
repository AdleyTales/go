package main

import "fmt"

func cc(n int) int {
  if n == 0 {
    return 1
  }

  return n*cc(n-1)
}

func main()  {
  fmt.Println(cc(4))
  fmt.Println(cc(14))
}

// 递归： 一直调用本身  改变条件
