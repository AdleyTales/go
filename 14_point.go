package main

import "fmt"

func main() {

	i := 231
	fmt.Println(i)
	fmt.Println(&i) //0xc042058080

	j := 231
	fmt.Println(&j) //0xc042058098  地址不相同

  x := j
  fmt.Println(*(&x) == *(&i)) //true

}


// &: 获取内存地址  *: 取值
