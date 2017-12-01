package main

import "fmt"
import "time"

func main()  {
  getVal(2)
  getTime()
}

func getVal(val int)  {
  switch val {
    case 1:
      println("one")
    case 2:
      println("two")
    case 3:
      println("three")
  }
}

func getTime()  {

  fmt.Println(time.Now())
  fmt.Println(time.Now().Weekday())

  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        println("it's the weekend")
    default:
        println("it's a weekday")
    }
}

//  在case语句中，使用 ，分隔不同的表达式
