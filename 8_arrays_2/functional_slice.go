package main

import (
	"fmt"
)

func main4() {
	// ============================================
	// MAP: TRANSFORM EACH ELEMENT
	// ============================================

	numbers := []int{1, 2, 3, 4, 5}
	squared := mapInt(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Original:", numbers)
	fmt.Println("Squared:", squared)

	// ============================================
	// FILTER: SELECT ELEMENTS
	// ============================================

	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filterInt(allNumbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("\nAll numbers:", allNumbers)
	fmt.Println("Even numbers:", evens)

	// ============================================
	// REDUCE: COMBINE TO SINGLE VALUE
	// ============================================

	sum := reduceInt(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("\nSum:", sum)

	product := reduceInt(numbers, 1, func(acc, n int) int {
		return acc * n
	})
	fmt.Println("Product:", product)

	// ============================================
	// CHAIN OPERATIONS
	// ============================================

	result := reduceInt(
		mapInt(
			filterInt(allNumbers, isEven),
			square,
		),
		0,
		add,
	)
	fmt.Println("\nSum of squares of evens:", result)
	// (2² + 4² + 6² + 8² + 10²) = 220

	// ============================================
	// ANY / ALL
	// ============================================

	hasEven := any(allNumbers, isEven)
	fmt.Println("\nHas even number:", hasEven)

	allPositive := all(allNumbers, func(n int) bool {
		return n > 0
	})
	fmt.Println("All positive:", allPositive)

	// ============================================
	// PARTITION
	// ============================================

	evens2, odds := partitionInt(allNumbers, isEven)
	fmt.Println("\nEvens:", evens2)
	fmt.Println("Odds:", odds)

	// ============================================
	// GROUP BY
	// ============================================

	words := []string{"apple", "banana", "apricot", "berry", "cherry"}
	grouped := groupBy(words, func(s string) string {
		return string(s[0]) // Group by first letter
	})

	fmt.Println("\nGrouped by first letter:")
	for key, vals := range grouped {
		fmt.Printf("%s: %v\n", key, vals)
	}
}

// ============================================
// MAP: TRANSFORM
// ============================================

func mapInt(slice []int, transform func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// ============================================
// FILTER: SELECT
// ============================================

func filterInt(slice []int, predicate func(int) bool) []int {
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// ============================================
// REDUCE: COMBINE
// ============================================

func reduceInt(slice []int, initial int, combine func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = combine(result, v)
	}
	return result
}

// ============================================
// ANY: CHECK IF ANY ELEMENT MATCHES
// ============================================

func any(slice []int, predicate func(int) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// ============================================
// ALL: CHECK IF ALL ELEMENTS MATCH
// ============================================

func all(slice []int, predicate func(int) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// ============================================
// PARTITION: SPLIT BY CONDITION
// ============================================

func partitionInt(slice []int, predicate func(int) bool) ([]int, []int) {
	matching := make([]int, 0)
	notMatching := make([]int, 0)

	for _, v := range slice {
		if predicate(v) {
			matching = append(matching, v)
		} else {
			notMatching = append(notMatching, v)
		}
	}

	return matching, notMatching
}

// ============================================
// GROUP BY: CREATE MAP OF SLICES
// ============================================

func groupBy(slice []string, keyFunc func(string) string) map[string][]string {
	result := make(map[string][]string)

	for _, v := range slice {
		key := keyFunc(v)
		result[key] = append(result[key], v)
	}

	return result
}

// ============================================
// HELPER FUNCTIONS
// ============================================

func isEven(n int) bool {
	return n%2 == 0
}

func square(n int) int {
	return n * n
}

func add(a, b int) int {
	return a + b
}
