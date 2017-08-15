package main

import "fmt"

func main() {
	greet("Christian")
	greet("Juan")
	fullname("Christian", "Romero")
}

func greet(name string) {
	fmt.Println(name)
}

func fullname(fname string, lname string) {
	fmt.Println(fname, lname)
}
