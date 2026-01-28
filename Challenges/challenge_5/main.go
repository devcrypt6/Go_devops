
// advance calculator

package main

import (
    "fmt"
    "math"
)

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}	
	return a / b, nil	
}

func modulus(a, b int) (int, error) {
    if b == 0 {
		return 0, fmt.Errorf("modulus by zero")
	}	
	return a % b, nil
}

func power(base, exponent float64) float64 {
    return math.Pow(base, exponent)
}

func main() {
    fmt.Println("=== Advanced Calculator ===")
    fmt.Println("Operators: +, -, *, /, %, ^ (power)")
    fmt.Println()
    
    var num1, num2 float64
    var operator string
    
    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    
    fmt.Print("Enter operator: ")
    fmt.Scan(&operator)
    
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)
    
    // Use switch statement to handle different operators
    switch operator {
    case "+":
        fmt.Println("Result:", add(num1, num2))
    case "-":
        fmt.Println("Result:", subtract(num1, num2))
    case "*":
        fmt.Println("Result:", multiply(num1, num2))
    case "/":
        result, err := divide(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)	
		} else {
			fmt.Println("Result:", result)
		}
    case "%":
        intNum1 := int(num1)
		intNum2 := int(num2)
		result, err := modulus(intNum1, intNum2)
		if err != nil {
			fmt.Println("Error:", err)	
		} else {
			fmt.Println("Result:", result)
		}
    case "^":
        fmt.Println("Result:", power(num1, num2))
    default:
        fmt.Println("Invalid operator!")
    }
}
