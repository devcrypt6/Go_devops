package main

import "fmt"

func main3() {
    // ============================================
    // SLICE STRUCTURE (Conceptually)
    // ============================================
    
    // A slice is a descriptor containing:
    // 1. Pointer to underlying array
    // 2. Length (number of elements)
    // 3. Capacity (max size before reallocation)
    
    // ============================================
    // CREATING SLICES FROM ARRAYS
    // ============================================
    
    arr := [5]int{10, 20, 30, 40, 50}
    
    // Slice entire array
    slice1 := arr[:]
    fmt.Println("Full slice:", slice1)
    
    // Slice from index 1 to 3 (3 not included!)
    slice2 := arr[1:4]  // [20 30 40]
    fmt.Println("Partial slice:", slice2)
    
    // Slice from start to index 3
    slice3 := arr[:3]  // [10 20 30]
    fmt.Println("From start:", slice3)
    
    // Slice from index 2 to end
    slice4 := arr[2:]  // [30 40 50]
    fmt.Println("To end:", slice4)
    
    // ============================================
    // SLICES SHARE UNDERLYING ARRAY!
    // ============================================
    
    fmt.Println("\n=== Shared Memory ===")
    original := []int{1, 2, 3, 4, 5}
    part := original[1:4]  // [2 3 4]
    
    fmt.Println("Original:", original)
    fmt.Println("Part:", part)
    
    // Modify slice
    part[0] = 999
    
    fmt.Println("After modifying part:")
    fmt.Println("Original:", original)  // [1 999 3 4 5] - CHANGED!
    fmt.Println("Part:", part)          // [999 3 4]
    
    // ============================================
    // NIL SLICES vs EMPTY SLICES
    // ============================================
    
    var nilSlice []int           // nil slice
    emptySlice := []int{}        // empty but not nil
    madeSlice := make([]int, 0)  // empty but not nil
    
    fmt.Println("\n=== Nil vs Empty ===")
    fmt.Printf("Nil slice: %v, len=%d, cap=%d, nil=%v\n", 
        nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
    fmt.Printf("Empty slice: %v, len=%d, cap=%d, nil=%v\n", 
        emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
    fmt.Printf("Made slice: %v, len=%d, cap=%d, nil=%v\n", 
        madeSlice, len(madeSlice), cap(madeSlice), madeSlice == nil)
    
    // All behave the same with append!
    nilSlice = append(nilSlice, 1)
    emptySlice = append(emptySlice, 1)
    madeSlice = append(madeSlice, 1)
}
