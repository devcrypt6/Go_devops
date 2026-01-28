package main

import "fmt"

func main7() {
    // Operator precedence (high to low):
    // 1. * / % << >> & &^
    // 2. + - | ^
    // 3. == != < <= > >=
    // 4. &&
    // 5. ||
    
    result := 10 + 5 * 2    // 20, not 30!
    fmt.Printf("10 + 5 * 2 = %d (multiplication first)\n", result)
    
    result = (10 + 5) * 2   // 30 (parentheses force addition first)
    fmt.Printf("(10 + 5) * 2 = %d\n", result)
    
    // Complex expression
    x := 2 + 3 * 4 - 10 / 2
    // Step-by-step: 2 + 12 - 5 = 9
    fmt.Printf("2 + 3 * 4 - 10 / 2 = %d\n", x)
}
