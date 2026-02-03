package main

import (
    "fmt"
    "strings"
)

func main7() {
    // ============================================
    // VARIADIC FUNCTIONS: VARIABLE ARGUMENTS
    // ============================================
    
    // Can pass any number of arguments [web:117]
    fmt.Println(sum1())           // 0 arguments
    fmt.Println(sum1(1))          // 1 argument
    fmt.Println(sum1(1, 2, 3))    // 3 arguments
    fmt.Println(sum1(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 10 arguments
    
    // ============================================
    // PASSING SLICE TO VARIADIC FUNCTION
    // ============================================
    
    numbers := []int{10, 20, 30, 40}
    total := sum1(numbers...) // Spread operator [web:117]
    fmt.Printf("Sum of %v = %d\n", numbers, total)
    
    // ============================================
    // VARIADIC WITH OTHER PARAMETERS
    // ============================================
    
    message := join(" ", "Hello", "from", "Go!")
    fmt.Println(message) // Hello from Go!
    
    csv := join(",", "Alice", "Bob", "Charlie")
    fmt.Println(csv) // Alice,Bob,Charlie
    
    // ============================================
    // PRINTF-STYLE FORMATTING (Real-world example)
    // ============================================
    
    logMessage("User logged in", "user_id", 123, "ip", "192.168.1.1")
}

// ============================================
// BASIC VARIADIC FUNCTION
// ============================================

func sum1(numbers ...int) int {
    total := 0
    // numbers is a slice of int [web:119]
    for _, n := range numbers {
        total += n
    }
    return total
}

// ============================================
// VARIADIC FUNCTION WITH PREFIX PARAMETER
// ============================================

// Variadic parameter MUST be last [web:118]
func join(delimiter string, parts ...string) string {
    return strings.Join(parts, delimiter)
}

// INVALID: Variadic must be last parameter
// func invalid(parts ...string, delimiter string) string {
//     return ""
// }

// ============================================
// VARIADIC WITH MULTIPLE TYPES
// ============================================

func printf(format string, values ...interface{}) {
    // interface{} accepts any type
    fmt.Printf(format+"\n", values...)
}

// ============================================
// REAL-WORLD: LOGGER WITH KEY-VALUE PAIRS
// ============================================

func logMessage(msg string, keyValuePairs ...interface{}) {
    fmt.Printf("[LOG] %s", msg)
    
    // Expect pairs: key1, value1, key2, value2...
    for i := 0; i < len(keyValuePairs); i += 2 {
        if i+1 < len(keyValuePairs) {
            fmt.Printf(" | %v=%v", keyValuePairs[i], keyValuePairs[i+1])
        }
    }
    fmt.Println()
}

// ============================================
// VARIADIC: FIND MAXIMUM
// ============================================

func max(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    maximum := numbers[0]
    for _, n := range numbers {
        if n > maximum {
            maximum = n
        }
    }
    return maximum
}

// ============================================
// VARIADIC: CONCATENATE STRINGS
// ============================================

func concat(separator string, strings ...string) string {
    result := ""
    for i, s := range strings {
        result += s
        if i < len(strings)-1 {
            result += separator
        }
    }
    return result
}
