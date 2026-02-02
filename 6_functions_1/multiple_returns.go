package main

import (
	"errors"
	"fmt"
)

func main() {
	// ============================================
	// MULTIPLE RETURN VALUES (Go's superpower!)
	// ============================================

	quot, rem := divMod(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quot, rem)

	// ============================================
	// ERROR HANDLING PATTERN (Most important!)
	// ============================================

	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// Try division by zero
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ============================================
	// IGNORING RETURN VALUES WITH _
	// ============================================

	// Only care about remainder
	_, remainder := divMod(20, 3)
	fmt.Println("Remainder:", remainder)

	// Only care about error (check operation succeeded)
	_, err = safeDivide(10, 2)
	if err != nil {
		fmt.Println("Operation failed")
	} else {
		fmt.Println("Operation succeeded")
	}
}

// ============================================
// RETURNING MULTIPLE VALUES
// ============================================

func divMod(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// ============================================
// VALUE + ERROR PATTERN (Standard Go idiom)
// ============================================

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil // nil means no error
}

// ============================================
// RETURNING SUCCESS STATUS
// ============================================

func parseAge(ageStr string) (int, bool) {
	// Simplified parsing
	if ageStr == "" {
		return 0, false // Parse failed
	}
	return 25, true // Parse succeeded
}

// ============================================
// RETURNING MULTIPLE DATA POINTS
// ============================================

func getStats(numbers []int) (int, int, float64) {
	if len(numbers) == 0 {
		return 0, 0, 0.0
	}

	min := numbers[0]
	max := numbers[0]
	sum := 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	avg := float64(sum) / float64(len(numbers))
	return min, max, avg
}

// ============================================
// SWAP FUNCTION (Classic example)
// ============================================

func swap(a, b int) (int, int) {
	return b, a
}

// Usage:
// x, y := 5, 10
// x, y = swap(x, y)  // Now x=10, y=5
