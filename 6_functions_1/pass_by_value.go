package main

import "fmt"

func main3() {
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
    modifySlice(numbers)
    fmt.Println("After modifySlice:", numbers)  // Changed!
    
    // ============================================
    // POINTERS FOR MODIFICATION (Preview)
    // ============================================
    
    y := 10
    fmt.Println("\nBefore modifyWithPointer:", y)
    modifyWithPointer(&y)
    fmt.Println("After modifyWithPointer:", y)  // Changed to 20!
}

// ============================================
// PASS BY VALUE: COPY IS MODIFIED
// ============================================

func modifyValue(num int) {
    num = num * 2
    fmt.Println("Inside modifyValue:", num)  // 20
    // Original variable unchanged!
}

// ============================================
// SLICES: UNDERLYING ARRAY IS SHARED
// ============================================

func modifySlice(slice []int) {
    slice[0] = 999  // Modifies original!
    fmt.Println("Inside modifySlice:", slice)
}

// ============================================
// POINTERS: MODIFY ORIGINAL VARIABLE
// ============================================

func modifyWithPointer(ptr *int) {
    *ptr = *ptr * 2  // Dereference and modify
}

// ============================================
// COPYING STRUCTS (Preview for later)
// ============================================

type Person struct {
    Name string
    Age  int
}

func modifyPerson(p Person) {
    p.Age = 30  // Only modifies copy!
}

func modifyPersonPointer(p *Person) {
    p.Age = 30  // Modifies original!
}
