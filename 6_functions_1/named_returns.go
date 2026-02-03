package main

import (
	"errors"
	"fmt"
)

func main6() {
	// ============================================
	// NAMED RETURNS IN ACTION
	// ============================================

	sum, product := calculate(5, 3)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// Named returns with error handling
	if area, perimeter, err := rectangleMetrics(5, 3); err == nil {
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)
	}

	// Complex function with named returns
	content, size, err := readFile("config.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", size, content)
	}
}

// ============================================
// BASIC NAMED RETURNS
// ============================================

func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // "naked return" - returns named values [web:123]
}

// ============================================
// NAMED RETURNS WITH DOCUMENTATION
// ============================================

// rectangleMetrics calculates area and perimeter
// Returns (area, perimeter, error)
func rectangleMetrics(length, width float64) (area float64, perimeter float64, err error) {
	if length <= 0 || width <= 0 {
		err = errors.New("dimensions must be positive")
		return // Returns zero values for area & perimeter, plus error
	}

	area = length * width
	perimeter = 2 * (length + width)
	return
}

// ============================================
// NAMED RETURNS MAKE CODE SELF-DOCUMENTING
// ============================================

func readFile(filename string) (content string, size int, err error) {
	// Simulate file reading
	if filename == "" {
		err = errors.New("filename cannot be empty")
		return
	}

	content = "File contents here..."
	size = len(content)
	return
}

// ============================================
// BEST PRACTICES WITH NAMED RETURNS
// ============================================

// Good: Clear what each return value represents [web:123][web:126]
func getUserInfo(id int) (name string, age int, isActive bool, err error) {
	// Function body...
	name = "John Doe"
	age = 25
	isActive = true
	return
}

// Bad: Named returns in simple function (unnecessary) [web:123]
func add1(a, b int) (result int) {
	result = a + b
	return
}

// Better: Just return directly for simple functions
func addBetter(a, b int) int {
	return a + b
}

// ============================================
// NAMED RETURNS WITH DEFER (Powerful pattern!)
// ============================================

func processData() (result int, err error) {
	// Named returns work with defer - we'll see more in defer section
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	result = 42
	return
}
