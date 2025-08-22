package main

import (
	"fmt"

	"github.com/ramana264/golang/mathops"
)

func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Println("Sum:", mathops.Add(num1, num2))
	fmt.Println("Product:", mathops.Multiply(num1, num2))
}
