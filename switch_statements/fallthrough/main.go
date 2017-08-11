package main

import "fmt"

func main() {
	switch "Iliana" {
	case "Iliana":
		fmt.Println("Wassup Iliana!")
		fallthrough
	case "Scarlett":
		fmt.Println("Wassup Scarlett!")
		fallthrough
	case "Victoria":
		fmt.Println("Wassup Victoria!")
	default:
		fmt.Println("Have you no friends?")

	}
}
