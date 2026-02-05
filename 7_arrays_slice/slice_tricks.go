package main

import "fmt"

func main6() {
    // ============================================
    // REMOVING ELEMENTS FROM SLICE
    // ============================================
    
    // Remove by index (not built-in, manual approach)
    
    // Method 1: Remove element at index i
    numbers := []int{10, 20, 30, 40, 50}
    fmt.Println("Original:", numbers)
    
    i := 2  // Remove index 2 (value 30)
    numbers = append(numbers[:i], numbers[i+1:]...)
    fmt.Println("After removing index 2:", numbers)  // [10 20 40 50]
    
    // Method 2: Remove last element (fast)
    numbers = []int{10, 20, 30, 40, 50}
    numbers = numbers[:len(numbers)-1]
    fmt.Println("After removing last:", numbers)  // [10 20 30 40]
    
    // Method 3: Remove first element
    numbers = []int{10, 20, 30, 40, 50}
    numbers = numbers[1:]
    fmt.Println("After removing first:", numbers)  // [20 30 40 50]
    
    // ============================================
    // INSERTING ELEMENTS
    // ============================================
    
    // Insert at index i
    numbers = []int{10, 20, 40, 50}
    i = 2
    value := 30
    
    // Grow slice by 1
    numbers = append(numbers, 0)
    // Shift elements right
    copy(numbers[i+1:], numbers[i:])
    // Insert value
    numbers[i] = value
    
    fmt.Println("\nAfter inserting 30 at index 2:", numbers)  // [10 20 30 40 50]
    
    // ============================================
    // INSERT AT BEGINNING
    // ============================================
    
    numbers = []int{20, 30, 40}
    numbers = append([]int{10}, numbers...)
    fmt.Println("After prepending 10:", numbers)  // [10 20 30 40]
    
    // ============================================
    // FILTER SLICE (Remove by condition)
    // ============================================
    
    numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("\nOriginal:", numbers)
    
    // Keep only even numbers
    filtered := numbers[:0]  // Reuse backing array
    for _, n := range numbers {
        if n%2 == 0 {
            filtered = append(filtered, n)
        }
    }
    numbers = filtered
    fmt.Println("After filtering (even only):", numbers)  // [2 4 6 8 10]
    
    // ============================================
    // CLEAR SLICE
    // ============================================
    
    numbers = []int{1, 2, 3, 4, 5}
    numbers = numbers[:0]  // Clear but keep capacity
    fmt.Printf("Cleared: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))
    
    // or completely nil it
    numbers = nil
    fmt.Printf("Nil: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))
}

// ============================================
// HELPER FUNCTIONS FOR SLICE OPERATIONS
// ============================================

// Remove element at index
func remove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

// Insert element at index
func insert(slice []int, index int, value int) []int {
    slice = append(slice, 0)       // Grow by 1
    copy(slice[index+1:], slice[index:])  // Shift right
    slice[index] = value            // Insert
    return slice
}

// Filter slice by predicate
func filter(slice []int, predicate func(int) bool) []int {
    result := slice[:0]
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}
