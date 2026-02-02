package main

import "fmt"

// ============================================
// FUNCTION SYNTAX:
// func functionName(parameters) returnType {
//     // function body
//     return value
// }
// ============================================

func main1() {
    // ============================================
    // CALLING FUNCTIONS
    // ============================================
    
    greet()                    // Function with no parameters, no return
    sayHello("Alice")         // Function with parameter, no return
    result := add(5, 3)       // Function with parameters and return value
    fmt.Println("5 + 3 =", result)
    
    // Call and use result directly
    fmt.Println("10 + 20 =", add(10, 20))
    
    // ============================================
    // FUNCTIONS CAN CALL OTHER FUNCTIONS
    // ============================================
    
    printSum(7, 8)
}

// ============================================
// FUNCTION WITH NO PARAMETERS, NO RETURN
// ============================================

func greet() {
    fmt.Println("Hello, World!")
}

// ============================================
// FUNCTION WITH PARAMETERS, NO RETURN
// ============================================

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// ============================================
// FUNCTION WITH PARAMETERS AND RETURN VALUE
// ============================================

func add(a int, b int) int {
    sum := a + b
    return sum
}

// Shorthand when parameters have same type
func multiply(a, b int) int {
    return a * b
}

// ============================================
// FUNCTION CALLING ANOTHER FUNCTION
// ============================================

func printSum(x, y int) {
    total := add(x, y)  // Call add() function
    fmt.Printf("%d + %d = %d\n", x, y, total)
}

// ============================================
// FUNCTION WITH MULTIPLE OPERATIONS
// ============================================

func calculateArea(length, width float64) float64 {
    area := length * width
    fmt.Printf("Calculating area of %.2f × %.2f\n", length, width)
    return area
}

// ============================================
// EARLY RETURN (Common pattern)
// ============================================

func divideSimple(a, b float64) float64 {
    if b == 0 {
        fmt.Println("Error: Division by zero")
        return 0  // Exit function early
    }
    
    return a / b
}
