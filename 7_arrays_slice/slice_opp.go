package main

import "fmt"

func main5() {
    // ============================================
    // APPEND: ADDING ELEMENTS TO SLICE
    // ============================================
    
    var numbers []int  // nil slice
    fmt.Println("Initial:", numbers)
    
    // Append single element
    numbers = append(numbers, 10)
    fmt.Println("After append 10:", numbers)
    
    // Append multiple elements
    numbers = append(numbers, 20, 30, 40)
    fmt.Println("After append 20,30,40:", numbers)
    
    // Append another slice (spread operator ...)
    more := []int{50, 60, 70}
    numbers = append(numbers, more...)
    fmt.Println("After appending slice:", numbers)
    
    // ============================================
    // APPEND GROWS CAPACITY AUTOMATICALLY
    // ============================================
    
    s := make([]int, 0, 2)  // length=0, capacity=2
    fmt.Println("\n=== Capacity Growth ===")
    fmt.Printf("len=%d cap=%d: %v\n", len(s), cap(s), s)
    
    s = append(s, 1)
    fmt.Printf("len=%d cap=%d: %v\n", len(s), cap(s), s)
    
    s = append(s, 2)
    fmt.Printf("len=%d cap=%d: %v\n", len(s), cap(s), s)
    
    s = append(s, 3)  // Exceeds capacity - reallocates!
    fmt.Printf("len=%d cap=%d: %v (capacity doubled!)\n", len(s), cap(s), s)
    
    // ============================================
    // COPY: EXPLICIT SLICE COPYING
    // ============================================
    
    fmt.Println("\n=== Copy Operation ===")
    source := []int{1, 2, 3, 4, 5}
    destination := make([]int, len(source))
    
    numCopied := copy(destination, source)
    fmt.Printf("Copied %d elements\n", numCopied)
    fmt.Println("Source:", source)
    fmt.Println("Destination:", destination)
    
    // Modify destination - source unchanged
    destination[0] = 999
    fmt.Println("\nAfter modifying destination:")
    fmt.Println("Source:", source)       // [1 2 3 4 5]
    fmt.Println("Destination:", destination)  // [999 2 3 4 5]
    
    // ============================================
    // COPY WITH DIFFERENT SIZES
    // ============================================
    
    src := []int{1, 2, 3, 4, 5}
    
    // Destination smaller - only copies what fits
    small := make([]int, 3)
    copy(small, src)
    fmt.Println("\nSmall destination:", small)  // [1 2 3]
    
    // Destination larger - fills available space
    large := make([]int, 7)
    copy(large, src)
    fmt.Println("Large destination:", large)  // [1 2 3 4 5 0 0]
}
