package main

import "fmt"

func main()  {

  m := make(map[string]int)

  m["k1"] = 23
  m["k2"] = 54
  m["k3"] = 4555555

  fmt.Println("map:",m) //map: map[k1:23 k2:54 k3:4555555]

  fmt.Println(len(m)) //3

  // delete 移除键值对
  delete(m,"k2")
  fmt.Println("map:",m) //  map: map[k3:4555555 k1:23]

  n := map[string]int {"n1": 2332, "n2":656, "n3":7676}
  fmt.Println(n)


}

// map 其他语言 字典 / 哈希
