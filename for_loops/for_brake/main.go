package main

import "fmt"

func main(){
  // i := 0
  // for i < 10 {
  //   fmt.Println(i)
  //   i++
  // }
  j := 0
  for {
    fmt.Println(j)
    if j >= 10{
      break
    }
    j++
  }
}
