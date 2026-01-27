package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Use if-else or switch to handle different operators
	// Handle division by zero!

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operator.")
		return
	}

	fmt.Printf("%v %s %v = %v\n", num1, operator, num2, result)

}
