package main

import "fmt"

type geometry interface {
  area() int
}

type rect struct {
  width,height int
}

func (m rect) area() int {
  return m.width * m.height
}

func getFn(g geometry)  {
  fmt.Println(g.area())
}

func main()  {

  s := rect(23,4)
  fmt.Println(s)

  getFn(s)


}
