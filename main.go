package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	sum := a + b
	product := a * b

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Product: %d\n", product)
}
