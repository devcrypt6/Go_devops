package main

import (
    "fmt"
    "sort"
)

func main3() {
    // ============================================
    // IN-PLACE FILTERING (Memory Efficient)
    // ============================================
    
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Original:", numbers)
    
    // Keep only even numbers
    numbers = filterInPlace(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println("After filtering (even):", numbers)
    
    // ============================================
    // IN-PLACE DEDUPLICATION
    // ============================================
    
    data := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
    fmt.Println("\nOriginal with duplicates:", data)
    
    data = deduplicateInPlace(data)
    fmt.Println("After deduplication:", data)
    
    // ============================================
    // IN-PLACE TRANSFORMATION
    // ============================================
    
    values := []int{1, 2, 3, 4, 5}
    fmt.Println("\nOriginal values:", values)
    
    transformInPlace(values, func(n int) int {
        return n * n
    })
    fmt.Println("After squaring:", values)
    
    // ============================================
    // PARTITION SLICE (Like quicksort)
    // ============================================
    
    nums := []int{3, 7, 1, 9, 2, 8, 4, 6, 5}
    fmt.Println("\nOriginal:", nums)
    
    pivot := partition(nums, 5)  // Partition around 5
    fmt.Printf("After partition: %v (pivot at index %d)\n", nums, pivot)
}

// ============================================
// FILTER IN-PLACE (Reuse backing array)
// ============================================

func filterInPlace(slice []int, keep func(int) bool) []int {
    n := 0  // Write index
    
    for _, val := range slice {
        if keep(val) {
            slice[n] = val
            n++
        }
    }
    
    return slice[:n]  // Return only kept elements
}

// ============================================
// DEDUPLICATE SORTED SLICE IN-PLACE
// ============================================

func deduplicateInPlace(slice []int) []int {
    if len(slice) == 0 {
        return slice
    }
    
    // Sort first
    sort.Ints(slice)
    
    n := 1  // Write index (keep first element)
    
    for i := 1; i < len(slice); i++ {
        if slice[i] != slice[i-1] {
            slice[n] = slice[i]
            n++
        }
    }
    
    return slice[:n]
}

// ============================================
// TRANSFORM IN-PLACE
// ============================================

func transformInPlace(slice []int, transform func(int) int) {
    for i := range slice {
        slice[i] = transform(slice[i])
    }
}

// ============================================
// PARTITION (Quicksort-style)
// ============================================

func partition(slice []int, pivot int) int {
    left := 0
    right := len(slice) - 1
    
    for left <= right {
        // Find element >= pivot from left
        for left <= right && slice[left] < pivot {
            left++
        }
        
        // Find element < pivot from right
        for left <= right && slice[right] >= pivot {
            right--
        }
        
        // Swap if needed
        if left < right {
            slice[left], slice[right] = slice[right], slice[left]
            left++
            right--
        }
    }
    
    return left
}

// ============================================
// COMPACT SLICE (Remove specific value)
// ============================================

func compactValue(slice []int, remove int) []int {
    n := 0
    for _, val := range slice {
        if val != remove {
            slice[n] = val
            n++
        }
    }
    return slice[:n]
}

// ============================================
// ROTATE IN-PLACE
// ============================================

func rotateInPlace(slice []int, k int) {
    if len(slice) == 0 {
        return
    }
    
    k = k % len(slice)
    if k < 0 {
        k += len(slice)
    }
    
    // Reverse entire slice
    reverse(slice)
    // Reverse first k elements
    reverse(slice[:k])
    // Reverse remaining elements
    reverse(slice[k:])
}

func reverse(slice []int) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}
