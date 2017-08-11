package main

import "fmt"


func main(){
  for i := 0; i < 66; i++{
    fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
  }
}
