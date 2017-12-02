package main

import "fmt"

type rect struct {
  width,height int
}

func (m rect) area() int {
  return m.width * m.height
}

func main()  {

  d := rect{23,43}
  fmt.Println(d.area())

}
