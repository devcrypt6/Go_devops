package main

import "fmt"

func main9() {
    name := "Gopher"
    age := 25
    height := 1.75
    isStudent := true
    pi := 3.14159265358979
    
    // ============================================
    // BASIC FORMAT VERBS
    // ============================================
    
    fmt.Println("=== Basic Formatting ===")
    fmt.Printf("%%v (default):  %v\n", name)    // Gopher
    fmt.Printf("%%s (string):   %s\n", name)    // Gopher
    fmt.Printf("%%d (decimal):  %d\n", age)     // 25
    fmt.Printf("%%f (float):    %f\n", height)  // 1.750000
    fmt.Printf("%%t (boolean):  %t\n", isStudent) // true
    fmt.Printf("%%T (type):     %T\n", name)    // string
    
    // ============================================
    // NUMBER FORMATTING
    // ============================================
    
    num := 42
    fmt.Println("\n=== Number Formatting ===")
    fmt.Printf("Decimal:     %d\n", num)      // 42
    fmt.Printf("Binary:      %b\n", num)      // 101010
    fmt.Printf("Octal:       %o\n", num)      // 52
    fmt.Printf("Hex:         %x\n", num)      // 2a
    fmt.Printf("Hex (upper): %X\n", num)      // 2A
    fmt.Printf("Unicode:     %c\n", num)      // * (ASCII 42)
    
    // ============================================
    // FLOAT PRECISION
    // ============================================
    
    fmt.Println("\n=== Float Precision ===")
    fmt.Printf("Default:       %f\n", pi)     // 3.141593
    fmt.Printf("2 decimals:    %.2f\n", pi)   // 3.14
    fmt.Printf("10 decimals:   %.10f\n", pi)  // 3.1415926536
    fmt.Printf("Scientific:    %e\n", pi)     // 3.141593e+00
    fmt.Printf("Compact:       %g\n", pi)     // 3.14159265358979
    
    // ============================================
    // WIDTH & PADDING
    // ============================================
    
    fmt.Println("\n=== Width & Padding ===")
    fmt.Printf("Right-aligned: |%10s|\n", name)   // |    Gopher|
    fmt.Printf("Left-aligned:  |%-10s|\n", name)  // |Gopher    |
    fmt.Printf("Zero-padded:   %05d\n", age)      // 00025
    
    // ============================================
    // ADVANCED: Positional Arguments
    // ============================================
    
    fmt.Println("\n=== Positional Arguments ===")
    fmt.Printf("%[2]d %[1]s\n", "apples", 5)  // 5 apples (reordered!)
    fmt.Printf("%[1]d %[1]d\n", 10)           // 10 10 (reused!)
    
    // ============================================
    // COMMON PATTERNS
    // ============================================
    
    fmt.Println("\n=== Common Patterns ===")
    price := 19.99
    fmt.Printf("Price: $%.2f\n", price)  // Price: $19.99
    
    percent := 0.753
    fmt.Printf("Completion: %.1f%%\n", percent*100) // 75.3%
    
    fmt.Printf("Table row: %-15s | %5d | %.2f\n", "Berlin", 3850000, 891.8)
}


//Format Verb Cheat Sheet:

	
// | Verb  | Meaning                    | Example          |
// | ----- | -------------------------- | ---------------- |
// | %v    | Default format             | "hello"          |
// | %+v   | With field names (structs) | {Name:Go Age:15} |
// | %#v   | Go-syntax representation   | "hello"          |
// | %T    | Type                       | string           |
// | %d    | Decimal integer            | 42               |
// | %b    | Binary                     | 101010           |
// | %f    | Float                      | 3.141593         |
// | %.2f  | Float with 2 decimals      | 3.14             |
// | %e    | Scientific notation        | 3.14e+00         |
// | %s    | String                     | hello            |
// | %q    | Quoted string              | "hello"          |
// | %10s  | Width 10, right-aligned    | hello            |
// | %-10s | Width 10, left-aligned     | hello            |

		