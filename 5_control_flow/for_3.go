package main

import (
    "fmt"
    "time"
)

func for_3() {
    // ============================================
    // INFINITE LOOP: for { }
    // ============================================
    
    // Omit all three components = infinite loop [web:105]
    
    counter := 0
    for {
        counter++
        fmt.Printf("Iteration %d\n", counter)
        
        // Always have an exit condition!
        if counter >= 5 {
            break // Exit the loop
        }
        
        time.Sleep(500 * time.Millisecond)
    }
    
    // ============================================
    // SERVER/WORKER PATTERN (Common in production)
    // ============================================
    
    fmt.Println("\n Simulating server loop...")
    
    requestCount := 0
    for {
        requestCount++
        fmt.Printf("Processing request #%d\n", requestCount)
        
        // Simulate processing time
        time.Sleep(300 * time.Millisecond)
        
        // Shutdown condition
        if requestCount >= 3 {
            fmt.Println(" Server shutting down...")
            break
        }
    }
    
    // ============================================
    // MENU SYSTEM (Real-world pattern)
    // ============================================
    
    for {
        fmt.Println("\n=== MENU ===")
        fmt.Println("1. Option A")
        fmt.Println("2. Option B")
        fmt.Println("3. Exit")
        fmt.Print("Choose: ")
        
        var choice int
        fmt.Scan(&choice)
        
        switch choice {
        case 1:
            fmt.Println("You chose A")
        case 2:
            fmt.Println("You chose B")
        case 3:
            fmt.Println("Goodbye!")
            return // Exit entire function
        default:
            fmt.Println("Invalid choice")
        }
    }
}
