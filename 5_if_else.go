package main

func main()  {
  getAge(20)
  getMax(14,23)
}

func getAge(age int) {
  if(age > 18) {
    println("你是成年人了~")
  } else {
    println("不好意思，你是未成年人")
  }
}

func getMax(a ,b int) {

  if a > b {
    println(a,"大于",b)
  }else {
    println(a,"小于",b)
  }

}

// 在go语言中 if else 可以省略() 但是{} 是必须的。
