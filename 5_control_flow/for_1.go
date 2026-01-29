package main

import "fmt"

func for_1() {
    // ============================================
    // GO HAS ONLY ONE LOOP: for
    // No while, no do-while, no foreach
    // ============================================
    
    // ============================================
    // CLASSIC THREE-COMPONENT FOR LOOP
    // for init; condition; post { }
    // ============================================
    
    fmt.Println("Count from 1 to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // ============================================
    // COUNTING DOWN
    // ============================================
    
    fmt.Println("\nCountdown:")
    for i := 5; i >= 1; i-- {
        fmt.Printf("%d ", i)
    }
    fmt.Println(" Blast off!")
    
    // ============================================
    // CUSTOM INCREMENT
    // ============================================
    
    fmt.Println("\nEven numbers 0-10:")
    for i := 0; i <= 10; i += 2 {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // ============================================
    // MULTIPLE VARIABLES (Less common)
    // ============================================
    
    fmt.Println("\nFibonacci-style loop:")
    for i, j := 0, 1; i < 20; i, j = i+j, i {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // ============================================
    // NESTED LOOPS
    // ============================================
    
    fmt.Println("\nMultiplication table (3x3):")
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("%d*%d=%d  ", i, j, i*j)
        }
        fmt.Println()
    }
}
