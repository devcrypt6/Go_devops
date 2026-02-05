package main

import "fmt"

func main7() {
    // ============================================
    // WHEN TO USE ARRAYS
    // ============================================
    
    // 1. Fixed-size data (like RGB color)
    type Color [3]uint8  // Always 3 values: R, G, B
    red := Color{255, 0, 0}
    fmt.Println("Red color:", red)
    
    // 2. When you need value semantics (full copy)
    arr1 := [3]int{1, 2, 3}
    arr2 := arr1  // Full copy
    arr2[0] = 999
    fmt.Println("arr1:", arr1)  // [1 2 3] - unchanged
    fmt.Println("arr2:", arr2)  // [999 2 3]
    
    // 3. Small, performance-critical code (stack allocation)
    
    // ============================================
    // WHEN TO USE SLICES (Most of the time!)
    // ============================================
    
    // 1. Unknown size at compile time
    var names []string  // Will grow dynamically
    names = append(names, "Alice", "Bob", "Charlie")
    
    // 2. When passing to functions (avoid copying)
    data := []int{1, 2, 3, 4, 5}
    processSlice(data)  // Only copies pointer, len, cap
    
    // 3. Need to grow/shrink collection
    dynamic := []int{1, 2, 3}
    dynamic = append(dynamic, 4, 5, 6)  // Easy to grow!
    
    // ============================================
    // PERFORMANCE COMPARISON
    // ============================================
    
    // Large array - expensive to pass
    largeArray := [1000]int{}
    processArray(largeArray)  // Copies 1000 integers!
    
    // Large slice - cheap to pass
    largeSlice := make([]int, 1000)
    processSlice(largeSlice)  // Only copies descriptor (24 bytes)
    
    fmt.Println("Use slices for dynamic, growing collections!")
    fmt.Println("Use arrays for fixed-size, value-semantic types!")
}

func processArray(arr [1000]int) {
    // Receives COPY of entire array
    // Expensive!
}

func processSlice(s []int) {
    // Receives copy of slice descriptor
    // Cheap! Only 24 bytes (pointer + len + cap)
    // But can modify underlying array!
}
