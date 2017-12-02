package main

import "fmt"

func plus(a int,b int) int {
  return a + b
}

func vals() (int, int)  {
  return 54, 90
}

func sum(nums ...int) int {
  fmt.Println(nums)

  total := 0
  for _,n := range nums {
    total += n
  }

  return total

}

func main()  {

  res := plus(21,23)
  fmt.Println(res)

  // ############### 函数，返回多个值
  x,v := vals()
  fmt.Println(x," ",v)

  _,d := vals()
  fmt.Println(d)

  // ############# 可变参数函数
  t := sum(21,32,43,54,64,23,21,32,43)
  fmt.Println(t)

}
