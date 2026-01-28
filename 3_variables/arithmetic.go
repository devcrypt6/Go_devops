package main

import (
    "fmt"
    "math"
)

func main4() {
    a, b := 17, 5
    
    fmt.Println("Arithmetic Operators:")
    fmt.Printf("%d + %d = %d\n", a, b, a+b)  // Addition: 22
    fmt.Printf("%d - %d = %d\n", a, b, a-b)  // Subtraction: 12
    fmt.Printf("%d * %d = %d\n", a, b, a*b)  // Multiplication: 85
    fmt.Printf("%d / %d = %d\n", a, b, a/b)  // Integer division: 3 (not 3.4!)
    fmt.Printf("%d %% %d = %d\n", a, b, a%b) // Modulus: 2
    
    // Float division
    fmt.Printf("\nFloat division:\n")
    fmt.Printf("%.2f / %.2f = %.2f\n", float64(a), float64(b), float64(a)/float64(b))
    
    // Increment/Decrement (ONLY postfix in Go!)
    x := 10
    x++     // Valid
    // ++x  // INVALID! Go doesn't allow prefix
    fmt.Printf("\nIncrement: %d\n", x) // 11
    
    x--
    fmt.Printf("Decrement: %d\n", x) // 10
    
    // Power (not an operator, use math.Pow)
    result := math.Pow(2, 10)
    fmt.Printf("\n2^10 = %.0f\n", result) // 1024
}
