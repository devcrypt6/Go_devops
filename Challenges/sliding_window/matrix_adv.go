package main

import "fmt"

// Spiral order traversal of matrix
// [[1,2,3],
//
//	[4,5,6],
//	[7,8,9]] → [1,2,3,6,9,8,7,4,5]
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

// Rotate matrix 90 degrees in-place
func rotateMatrixInPlace(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

// Search in row-wise and column-wise sorted matrix
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	i, j := 0, cols-1
	for i < rows && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// Set entire row and column to zero if element is zero
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])
	rowZero := false
	colZero := false
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			colZero = true
		}
	}
	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < rows; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < cols; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if rowZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
	if colZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}

func main2() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Original:")
	printMatrix(matrix)

	spiral := spiralOrder(matrix)
	fmt.Println("\nSpiral order:", spiral)
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
