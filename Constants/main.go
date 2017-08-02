package main

import "fmt"

const p = "death & taxes"

//Multiple declaration

const (
	PI       = 3.14
	Language = "GO"
	_        = iota             //0
	KB       = 1 << (iota * 10) // 1<<(1*10)
	MB       = 1 << (iota * 10) // 1<<(2*10)
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("PI - ", PI)
	fmt.Println("Language - ", Language)
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
