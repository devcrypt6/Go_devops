package main

import "fmt"

func main2() {
    // ============================================
    // ARRAY INITIALIZATION PATTERNS
    // ============================================
    
    // Specific indices
    arr1 := [5]int{0: 10, 2: 30, 4: 50}
    fmt.Println("Sparse array:", arr1)  // [10 0 30 0 50]
    
    // All same value (manual)
    var zeros [10]int  // All zeros by default
    fmt.Println("Zeros:", zeros)
    
    // ============================================
    // ARRAY COMPARISON
    // ============================================
    
    a1 := [3]int{1, 2, 3}
    a2 := [3]int{1, 2, 3}
    a3 := [3]int{3, 2, 1}
    
    fmt.Println("\n=== Array Comparison ===")
    fmt.Println("a1 == a2:", a1 == a2)  // true
    fmt.Println("a1 == a3:", a1 == a3)  // false
    
    // ❌ Can't compare arrays of different sizes
    // a4 := [4]int{1, 2, 3, 4}
    // fmt.Println(a1 == a4)  // COMPILE ERROR!
    
    // ============================================
    // ARRAY COPYING (Pass by Value!)
    // ============================================
    
    fmt.Println("\n=== Array Copying ===")
    original := [3]int{1, 2, 3}
    copied := original  // Creates a COPY
    
    copied[0] = 999
    fmt.Println("Original:", original)  // [1 2 3] (unchanged!)
    fmt.Println("Copied:", copied)      // [999 2 3]
    
    // ============================================
    // MULTIDIMENSIONAL ARRAYS
    // ============================================
    
    // 2D array (matrix)
    var matrix [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("\n=== 2D Array ===")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
    
    // Tic-tac-toe board
    board := [3][3]string{
        {"X", "O", "X"},
        {"O", "X", "O"},
        {"O", "X", "X"},
    }
    
    fmt.Println("\n=== Tic-Tac-Toe Board ===")
    for _, row := range board {
        for _, cell := range row {
            fmt.Printf(" %s ", cell)
        }
        fmt.Println()
    }
}

// ============================================
// ARRAY AS FUNCTION PARAMETER (Copies!)
// ============================================

func modifyArray(arr [5]int) {
    arr[0] = 999  // Modifies COPY, not original!
}

// To modify original, use pointer
func modifyArrayPointer(arr *[5]int) {
    arr[0] = 999  // Modifies original
}
