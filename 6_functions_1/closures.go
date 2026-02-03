package main

import "fmt"

func main8() {
    // ============================================
    // ANONYMOUS FUNCTIONS
    // ============================================
    
    // Define and call immediately
    func() {
        fmt.Println("Hello from anonymous function!")
    }() // () calls it immediately
    
    // Anonymous function with parameters
    result := func(a, b int) int {
        return a + b
    }(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    // ============================================
    // ASSIGNING ANONYMOUS FUNCTIONS TO VARIABLES
    // ============================================
    
    multiply := func(a, b int) int {
        return a * b
    }
    
    fmt.Printf("4 × 5 = %d\n", multiply(4, 5))
    
    // ============================================
    // CLOSURES: ACCESSING OUTER VARIABLES [web:113]
    // ============================================
    
    // Closure captures 'message' variable
    message := "Hello"
    greet := func(name string) {
        fmt.Printf("%s, %s!\n", message, name) // Accesses outer variable
    }
    
    greet("Alice")
    message = "Hi"
    greet("Bob") // Uses updated message!
    
    // ============================================
    // CLOSURE FACTORY PATTERN [web:116]
    // ============================================
    
    addFive := makeAdder(5)
    addTen := makeAdder(10)
    
    fmt.Println(addFive(3))  // 8 (3 + 5)
    fmt.Println(addTen(3))   // 13 (3 + 10)
    
    // ============================================
    // COUNTER CLOSURE (Encapsulation) [web:113]
    // ============================================
    
    counter := makeCounter()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
    
    // New counter - independent state!
    counter2 := makeCounter()
    fmt.Println(counter2()) // 1
    
    // ============================================
    // MULTIPLIER CLOSURE [web:116]
    // ============================================
    
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
    
    // ============================================
    // CLOSURES IN GOROUTINES (Preview for later)
    // ============================================
    
    for i := 1; i <= 3; i++ {
        // Each closure captures different i (Go 1.22+) [web:116]
        go func() {
            fmt.Printf("Goroutine %d\n", i)
        }()
    }
    
    // Wait a bit (proper sync in concurrency section)
    fmt.Scanln()
}

// ============================================
// CLOSURE FACTORY: ADDER
// ============================================

func makeAdder(x int) func(int) int {
    // Returns a closure that captures x [web:113]
    return func(y int) int {
        return x + y
    }
}

// ============================================
// CLOSURE: STATEFUL COUNTER
// ============================================

func makeCounter() func() int {
    count := 0 // Captured by closure [web:113]
    
    return func() int {
        count++
        return count
    }
}

// ============================================
// CLOSURE: MULTIPLIER FACTORY
// ============================================

func makeMultiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

// ============================================
// PRACTICAL: FILTER FUNCTION
// ============================================

func filter(numbers []int, predicate func(int) bool) []int {
    var result []int
    for _, n := range numbers {
        if predicate(n) { // Call the closure
            result = append(result, n)
        }
    }
    return result
}

// Usage example (in main):
// numbers := []int{1, 2, 3, 4, 5, 6}
// evens := filter(numbers, func(n int) bool { return n%2 == 0 })
