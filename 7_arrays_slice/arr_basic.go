package main

import "fmt"

func main1() {
    // ============================================
    // ARRAYS: FIXED-SIZE, SAME-TYPE COLLECTIONS
    // ============================================
    
    // Arrays have FIXED LENGTH (cannot change!)
    // Syntax: var name [size]type
    
    // ============================================
    // ARRAY DECLARATION
    // ============================================
    
    // Method 1: Declare and initialize later
    var numbers [5]int
    fmt.Println("Empty array:", numbers)  // [0 0 0 0 0] (zero values)
    
    // Method 2: Declare with values
    var fruits [3]string = [3]string{"Apple", "Banana", "Cherry"}
    fmt.Println("Fruits:", fruits)
    
    // Method 3: Short declaration
    ages := [4]int{25, 30, 35, 40}
    fmt.Println("Ages:", ages)
    
    // Method 4: Let compiler count (...)
    colors := [...]string{"Red", "Green", "Blue", "Yellow"}
    fmt.Println("Colors:", colors)
    fmt.Println("Length:", len(colors))  // 4
    
    // ============================================
    // ACCESSING ARRAY ELEMENTS
    // ============================================
    
    fmt.Println("\n=== Accessing Elements ===")
    fmt.Println("First fruit:", fruits[0])   // Apple
    fmt.Println("Second fruit:", fruits[1])  // Banana
    fmt.Println("Last fruit:", fruits[2])    // Cherry
    
    // ❌ Index out of bounds (runtime panic!)
    // fmt.Println(fruits[3])  // PANIC: index out of range
    
    // ============================================
    // MODIFYING ARRAY ELEMENTS
    // ============================================
    
    fmt.Println("\n=== Modifying Elements ===")
    fruits[1] = "Blueberry"
    fmt.Println("Modified fruits:", fruits)
    
    // ============================================
    // ITERATING ARRAYS: Method 1 (Traditional)
    // ============================================
    
    fmt.Println("\n=== Iteration Method 1 ===")
    for i := 0; i < len(numbers); i++ {
        numbers[i] = i * 10
        fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
    }
    
    // ============================================
    // ITERATING ARRAYS: Method 2 (Range)
    // ============================================
    
    fmt.Println("\n=== Iteration Method 2 ===")
    for index, value := range colors {
        fmt.Printf("colors[%d] = %s\n", index, value)
    }
    
    // Only values
    fmt.Println("\nJust values:")
    for _, color := range colors {
        fmt.Printf("- %s\n", color)
    }
}
