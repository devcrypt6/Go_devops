package main

import "fmt"

func main5() {
    // ============================================
    // 2D SLICES (MATRIX)
    // ============================================
    
    // Method 1: Literal initialization
    matrix1 := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("Matrix 1:")
    printMatrix(matrix1)
    
    // Method 2: Dynamic creation
    rows, cols := 3, 4
    matrix2 := make2DSlice(rows, cols)
    
    // Fill with values
    counter := 1
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            matrix2[i][j] = counter
            counter++
        }
    }
    
    fmt.Println("\nMatrix 2:")
    printMatrix(matrix2)
    
    // ============================================
    // JAGGED SLICES (Different row lengths)
    // ============================================
    
    jagged := [][]int{
        {1},
        {2, 3},
        {4, 5, 6},
        {7, 8, 9, 10},
    }
    
    fmt.Println("\nJagged slice:")
    printMatrix(jagged)
    
    // ============================================
    // 3D SLICES
    // ============================================
    
    cube := make3DSlice(2, 3, 4)
    fmt.Printf("\n3D Slice: %d×%d×%d\n", len(cube), len(cube[0]), len(cube[0][0]))
    
    // ============================================
    // MATRIX OPERATIONS
    // ============================================
    
    m := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("\nOriginal matrix:")
    printMatrix(m)
    
    transposed := transpose(m)
    fmt.Println("\nTransposed:")
    printMatrix(transposed)
    
    rotated := rotate90Clockwise(m)
    fmt.Println("\nRotated 90° clockwise:")
    printMatrix(rotated)
}

// ============================================
// CREATE 2D SLICE
// ============================================

func make2DSlice(rows, cols int) [][]int {
    // Allocate outer slice
    matrix := make([][]int, rows)
    
    // Allocate each inner slice
    for i := range matrix {
        matrix[i] = make([]int, cols)
    }
    
    return matrix
}

// ============================================
// CREATE 3D SLICE
// ============================================

func make3DSlice(x, y, z int) [][][]int {
    cube := make([][][]int, x)
    
    for i := range cube {
        cube[i] = make([][]int, y)
        for j := range cube[i] {
            cube[i][j] = make([]int, z)
        }
    }
    
    return cube
}

// ============================================
// PRINT MATRIX
// ============================================

func printMatrix(matrix [][]int) {
    for _, row := range matrix {
        for _, val := range row {
            fmt.Printf("%3d ", val)
        }
        fmt.Println()
    }
}

// ============================================
// TRANSPOSE MATRIX
// ============================================

func transpose(matrix [][]int) [][]int {
    if len(matrix) == 0 {
        return matrix
    }
    
    rows := len(matrix)
    cols := len(matrix[0])
    
    result := make([][]int, cols)
    for i := range result {
        result[i] = make([]int, rows)
    }
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            result[j][i] = matrix[i][j]
        }
    }
    
    return result
}

// ============================================
// ROTATE 90° CLOCKWISE
// ============================================

func rotate90Clockwise(matrix [][]int) [][]int {
    n := len(matrix)
    if n == 0 {
        return matrix
    }
    
    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, n)
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            result[j][n-1-i] = matrix[i][j]
        }
    }
    
    return result
}

// ============================================
// FLATTEN 2D TO 1D
// ============================================

func flatten(matrix [][]int) []int {
    totalSize := 0
    for _, row := range matrix {
        totalSize += len(row)
    }
    
    result := make([]int, 0, totalSize)
    for _, row := range matrix {
        result = append(result, row...)
    }
    
    return result
}

// ============================================
// RESHAPE 1D TO 2D
// ============================================

func reshape(slice []int, rows, cols int) ([][]int, error) {
    if len(slice) != rows*cols {
        return nil, fmt.Errorf("cannot reshape: size mismatch")
    }
    
    result := make([][]int, rows)
    for i := range result {
        result[i] = make([]int, cols)
        copy(result[i], slice[i*cols:(i+1)*cols])
    }
    
    return result, nil
}

// ============================================
// MATRIX MULTIPLICATION (Bonus)
// ============================================

func matrixMultiply(a, b [][]int) ([][]int, error) {
    if len(a) == 0 || len(b) == 0 {
        return nil, fmt.Errorf("empty matrix")
    }
    
    if len(a[0]) != len(b) {
        return nil, fmt.Errorf("incompatible dimensions")
    }
    
    rows := len(a)
    cols := len(b[0])
    inner := len(a[0])
    
    result := make([][]int, rows)
    for i := range result {
        result[i] = make([]int, cols)
    }
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            for k := 0; k < inner; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    
    return result, nil
}
