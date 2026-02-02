package main

import "fmt"

func main4() {
    // ============================================
    // GO PASSES ARGUMENTS BY VALUE (COPY)
    // ============================================
    
    x := 10
    fmt.Println("Before modifyValue:", x)
    modifyValue(x)
    fmt.Println("After modifyValue:", x)  // Still 10!
    
    // ============================================
    // SLICES ARE SPECIAL (Reference type)
    // ============================================
    
    numbers := []int{1, 2, 3}
    fmt.Println("\nBefore modifySlice:", numbers)
    modifySlice2(numbers)
    fmt.Println("After modifySlice:", numbers)  // Changed!
    
    // ============================================
    // POINTERS FOR MODIFICATION (Preview)
    // ============================================
    
    y := 10
    fmt.Println("\nBefore modifyWithPointer:", y)
    modifyWithPointer2(&y)
    fmt.Println("After modifyWithPointer:", y)  // Changed to 20!
}

// ============================================
// PASS BY VALUE: COPY IS MODIFIED
// ============================================

func modifyValue2(num int) {
    num = num * 2
    fmt.Println("Inside modifyValue2:", num)  // 20
    // Original variable unchanged!
}

// ============================================
// SLICES: UNDERLYING ARRAY IS SHARED
// ============================================

func modifySlice2(slice []int) {
    slice[0] = 999  // Modifies original!
    fmt.Println("Inside modifySlice2:", slice)
}

// ============================================
// POINTERS: MODIFY ORIGINAL VARIABLE
// ============================================

func modifyWithPointer2(ptr *int) {
    *ptr = *ptr * 2  // Dereference and modify
}

// ============================================
// COPYING STRUCTS (Preview for later)
// ============================================

type Person2 struct {
    Name string
    Age  int
}

func modifyPerson2(p Person2) {
    p.Age = 30  // Only modifies copy!
}

func modifyPersonPointer2(p *Person2) {
    p.Age = 30  // Modifies original!
}
