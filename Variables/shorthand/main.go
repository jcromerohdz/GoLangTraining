package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", d, d)
}
