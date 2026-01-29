package main

import (
    "fmt"
    "math/rand"
    "time"
)

func for_2() {
    // ============================================
    // WHILE-STYLE LOOP (Only condition)
    // ============================================
    
    fmt.Println("Count to 5 (while-style):")
    i := 1
    for i <= 5 {
        fmt.Printf("%d ", i)
        i++
    }
    fmt.Println()
    
    // ============================================
    // READING UNTIL CONDITION MET
    // ============================================
    
    fmt.Println("\nGuess the number (1-10):")
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(10) + 1
    guess := 0
    attempts := 0
    
    for guess != target {
        attempts++
        fmt.Print("Enter your guess: ")
        fmt.Scan(&guess)
        
        if guess < target {
            fmt.Println("Too low!")
        } else if guess > target {
            fmt.Println("Too high!")
        }
    }
    
    fmt.Printf("🎉 Correct! You guessed it in %d attempts!\n", attempts)
    
    // ============================================
    // PROCESSING UNTIL VALID INPUT
    // ============================================
    
    var age int
    for age < 0 || age > 120 {
        fmt.Print("\nEnter your age (0-120): ")
        fmt.Scan(&age)
        
        if age < 0 || age > 120 {
            fmt.Println(" Invalid age. Try again.")
        }
    }
    fmt.Printf(" Age %d is valid\n", age)
}
