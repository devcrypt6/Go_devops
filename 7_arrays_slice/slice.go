package main

import "fmt"

func main() {
    // ============================================
    // SLICES: DYNAMIC, FLEXIBLE ARRAYS
    // ============================================
    
    // Slices are like arrays but:
    // - Size can change (grow/shrink)
    // - More commonly used than arrays
    // - Reference type (not copied like arrays)
    
    // ============================================
    // SLICE DECLARATION
    // ============================================
    
    // Method 1: Declare without size
    var numbers []int  // nil slice
    fmt.Println("Nil slice:", numbers)
    fmt.Println("Length:", len(numbers))
    fmt.Println("Is nil:", numbers == nil)
    
    // Method 2: Initialize with values
    fruits := []string{"Apple", "Banana", "Cherry"}
    fmt.Println("\nFruits:", fruits)
    fmt.Println("Length:", len(fruits))
    
    // Method 3: Using make()
    ages := make([]int, 5)  // length = 5, all zeros
    fmt.Println("\nMake slice:", ages)
    
    // Method 4: make with length and capacity
    scores := make([]int, 3, 10)  // length=3, capacity=10
    fmt.Println("Scores:", scores)
    fmt.Println("Length:", len(scores))
    fmt.Println("Capacity:", cap(scores))
    
    // ============================================
    // ACCESSING & MODIFYING SLICES
    // ============================================
    
    fmt.Println("\n=== Accessing Elements ===")
    fmt.Println("First fruit:", fruits[0])
    fmt.Println("Last fruit:", fruits[len(fruits)-1])
    
    fruits[1] = "Blueberry"
    fmt.Println("Modified:", fruits)
    
    // ============================================
    // ITERATING SLICES
    // ============================================
    
    fmt.Println("\n=== Iteration ===")
    for i, fruit := range fruits {
        fmt.Printf("%d: %s\n", i, fruit)
    }
    
    // ============================================
    // SLICE LENGTH vs CAPACITY
    // ============================================
    
    s := make([]int, 3, 5)
    fmt.Println("\n=== Length vs Capacity ===")
    fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
    fmt.Println("Slice:", s)
    
    // Length: number of elements currently in slice
    // Capacity: space allocated for slice (can grow up to this)
}
