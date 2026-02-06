package main

import (
	"fmt"
)

func main6() {
	demonstratePitfalls()
	demonstrateBestPractices()
}

func demonstratePitfalls() {
	fmt.Println("=== COMMON PITFALLS ===\n")

	// ============================================
	// PITFALL 1: Forgetting to assign append result
	// ============================================

	s := []int{1, 2, 3}
	_ = append(s, 4)                          //  WRONG: Result discarded!
	fmt.Println("After forgotten append:", s) // Still [1 2 3]

	s = append(s, 4)                        // ✅ CORRECT
	fmt.Println("After correct append:", s) // [1 2 3 4]

	// ============================================
	// PITFALL 2: Appending in loop with wrong slice
	// ============================================

	data := []int{1, 2, 3, 4, 5}

	// WRONG: Appending to same slice you're iterating
	for i, v := range data {
		if v%2 == 0 {
			data = append(data, v*10) // Dangerous!
		}
		fmt.Printf("Iteration %d: %v\n", i, data)
	}
	// May cause infinite loop or unexpected behavior!

	// CORRECT: Append to different slice
	data = []int{1, 2, 3, 4, 5}
	var evens []int
	for _, v := range data {
		if v%2 == 0 {
			evens = append(evens, v*10)
		}
	}
	fmt.Println("Evens:", evens)

	// ============================================
	// PITFALL 3: Slice expression bounds
	// ============================================

	s2 := []int{1, 2, 3, 4, 5}

	// Valid indices
	fmt.Println("\ns2[1:3]:", s2[1:3]) // [2 3]
	fmt.Println("s2[:3]:", s2[:3])     // [1 2 3]
	fmt.Println("s2[2:]:", s2[2:])     // [3 4 5]

	// Runtime panic
	// fmt.Println(s2[1:10])  // PANIC: out of bounds

	// ============================================
	// PITFALL 4: Comparing slices with ==
	// ============================================

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	// COMPILE ERROR: Cannot use == with slices
	// if a == b { }

	// CORRECT: Use custom comparison
	if slicesEqual(a, b) {
		fmt.Println("\nSlices are equal")
	}

	// ============================================
	// PITFALL 5: nil vs empty slice confusion
	// ============================================

	var nilSlice []int
	emptySlice := []int{}

	fmt.Printf("\nnil slice: len=%d, nil=%v\n", len(nilSlice), nilSlice == nil)
	fmt.Printf("empty slice: len=%d, nil=%v\n", len(emptySlice), emptySlice == nil)

	// Both work with append, range, len
	// Prefer nil for zero-value slices
}

func demonstrateBestPractices() {
	fmt.Println("\n\n=== BEST PRACTICES ===\n")

	// ============================================
	// PRACTICE 1: Pre-allocate known sizes
	// ============================================

	// GOOD
	users := make([]User, 0, 100) // Expect ~100 users
	for i := 0; i < 100; i++ {
		users = append(users, User{ID: i})
	}

	// ============================================
	// PRACTICE 2: Use copy for independence
	// ============================================

	original := []int{1, 2, 3, 4, 5}

	//  GOOD: Independent copy
	copied := make([]int, len(original))
	copy(copied, original)

	// ============================================
	// PRACTICE 3: Clear slice properly
	// ============================================

	// For primitive types
	numbers := []int{1, 2, 3, 4, 5}
	numbers = numbers[:0] // Clear, keep capacity

	// For pointer types (help GC)
	ptrs := make([]*User, 5)
	for i := range ptrs {
		ptrs[i] = nil // Clear references first
	}
	ptrs = ptrs[:0]

	// ============================================
	// PRACTICE 4: Check nil before ranging
	// ============================================

	var data []int
	// Safe to range over nil slice
	for _, v := range data {
		fmt.Println(v) // Never executes
	}

	// ============================================
	// PRACTICE 5: Return nil for empty result
	// ============================================

	// GOOD: Consistent zero value
	result := findItems() // Returns nil if not found
	if result == nil {
		fmt.Println("No items found")
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findItems() []int {
	// Return nil instead of []int{} for empty result
	return nil
}

type User struct {
	ID int
}

// ============================================
// BEST PRACTICE: Function Signatures
// ============================================

// GOOD: Accept slice (flexible)
func processItems(items []int) {
	for _, item := range items {
		// Process
		_ = item
	}
}

// BAD: Accept array (inflexible, copies)
func processItemsArray(items [10]int) {
	// Fixed size, full copy
}

// ============================================
// BEST PRACTICE: Return Slices
// ============================================

// GOOD: Return slice
func getNumbers() []int {
	return []int{1, 2, 3}
}

// Pre-allocate if size known
func makeRange(n int) []int {
	result := make([]int, n) // Pre-allocate
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}
