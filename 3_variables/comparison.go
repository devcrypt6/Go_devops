package main

import "fmt"

func main5() {
    x, y := 10, 20
    
    // Comparison operators
    fmt.Println("Comparison Operators:")
    fmt.Printf("%d == %d: %v\n", x, y, x == y) // false
    fmt.Printf("%d != %d: %v\n", x, y, x != y) // true
    fmt.Printf("%d < %d:  %v\n", x, y, x < y)  // true
    fmt.Printf("%d <= %d: %v\n", x, y, x <= y) // true
    fmt.Printf("%d > %d:  %v\n", x, y, x > y)  // false
    fmt.Printf("%d >= %d: %v\n", x, y, x >= y) // false
    
    // Logical operators
    isStudent := true
    hasPassed := true
    
    fmt.Println("\nLogical Operators:")
    fmt.Printf("isStudent && hasPassed: %v\n", isStudent && hasPassed) // AND: true
    fmt.Printf("isStudent || hasPassed: %v\n", isStudent || hasPassed) // OR: true
    fmt.Printf("!isStudent: %v\n", !isStudent)                         // NOT: false
    
    // Combining conditions
    age := 22
    hasLicense := true
    canDrive := age >= 18 && hasLicense
    fmt.Printf("\nCan drive (age>=18 AND hasLicense): %v\n", canDrive)
}
