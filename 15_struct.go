package main

import "fmt"

type Person struct {
  name string
  age int
}

func main()  {
  fmt.Println(Person{"fly",23}) // {fly 23}
  fmt.Println(Person{name:"adley", age: 23}) //{adley 23

  p := Person{"aaaa",24}
  fmt.Println(p.age) //24

}
