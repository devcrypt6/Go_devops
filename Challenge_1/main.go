package main

import "fmt"


// Understand the Challenge 
// We need to know first the formula to convert Celsius to Fahrenheit and vice versa.
// The formula to convert Celsius to Fahrenheit is: F = (C * 9/5) + 32
// The formula to convert Fahrenheit to Celsius is: C = (F - 32) * 5/9


func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() { 
	var choice int 
	var temp float64

	fmt.Println("Temperature Conversion")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scan(&choice)

	switch choice {	
	case 1:
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scan(&temp)
		result := celsiusToFahrenheit(temp)
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temp, result) // we use %.2f to format float to 2 decimal places
	case 2:
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scan(&temp)
		result := fahrenheitToCelsius(temp)
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temp, result)
	default:
		fmt.Println("Invalid choice! Please select 1 or 2.")
	}
}