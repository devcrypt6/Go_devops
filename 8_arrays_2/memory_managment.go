package main

import (
	"fmt"
)

func main2() {
	// ============================================
	// MEMORY LEAK: SLICE HOLDING REFERENCE
	// ============================================

	// Problem: Slicing can keep entire backing array in memory!
	demonstrateMemoryLeak()

	// ============================================
	// SOLUTION: COPY TO NEW SLICE
	// ============================================

	demonstrateMemoryFix()

	// ============================================
	// APPENDING TO SUB-SLICE CAN MODIFY ORIGINAL
	// ============================================

	demonstrateAppendPitfall()

	// ============================================
	// SLICE HEADER COPYING
	// ============================================

	demonstrateSliceCopying()
}

func demonstrateMemoryLeak() {
	fmt.Println("=== Memory Leak Example ===")

	// Create large slice
	largeSlice := make([]byte, 1000000) // 1 MB
	for i := range largeSlice {
		largeSlice[i] = byte(i % 256)
	}

	// Want only first 10 bytes, but keep reference to entire array!
	small := largeSlice[:10] // BAD: Holds reference to 1 MB!

	fmt.Printf("Small slice len=%d, cap=%d\n", len(small), cap(small))
	fmt.Println("But backing array is still 1 MB in memory!")

	// largeSlice can't be garbage collected while small exists
}

func demonstrateMemoryFix() {
	fmt.Println("\n=== Memory Leak Fixed ===")

	largeSlice := make([]byte, 1000000)

	// GOOD: Copy to new slice
	small := make([]byte, 10)
	copy(small, largeSlice[:10])

	fmt.Printf("Small slice len=%d, cap=%d\n", len(small), cap(small))
	fmt.Println("Now largeSlice can be garbage collected!")

	// largeSlice no longer referenced, can be freed
}

func demonstrateAppendPitfall() {
	fmt.Println("\n=== Append Sub-Slice Pitfall ===")

	original := []int{1, 2, 3, 4, 5}
	sub := original[0:3] // [1, 2, 3]

	fmt.Println("Original:", original)
	fmt.Println("Sub-slice:", sub)
	fmt.Printf("Sub capacity: %d\n", cap(sub))

	// Append to sub-slice
	sub = append(sub, 999)

	fmt.Println("\nAfter append(sub, 999):")
	fmt.Println("Original:", original) // [1 2 3 999 5] - MODIFIED!
	fmt.Println("Sub-slice:", sub)     // [1 2 3 999]

	//  SOLUTION: Force new allocation
	sub2 := append([]int{}, original[0:3]...) // Copy first
	sub2 = append(sub2, 888)

	fmt.Println("\nWith copy:")
	fmt.Println("Original:", original) // Unchanged
	fmt.Println("Sub2:", sub2)
}

func demonstrateSliceCopying() {
	fmt.Println("\n=== Slice Assignment ===")

	s1 := []int{1, 2, 3}
	s2 := s1 // Copies slice header (pointer, len, cap)

	// Both point to same underlying array
	s2[0] = 999

	fmt.Println("s1:", s1) // [999 2 3] - AFFECTED!
	fmt.Println("s2:", s2) // [999 2 3]

	// To make independent copy:
	s3 := make([]int, len(s1))
	copy(s3, s1)
	s3[0] = 111

	fmt.Println("\nWith copy:")
	fmt.Println("s1:", s1) // [999 2 3] - unchanged
	fmt.Println("s3:", s3) // [111 2 3]
}

// ============================================
// PRODUCTION PATTERN: Safe Sub-Slicing
// ============================================

//  BAD: May cause memory leak
func getFirstNBytes(data []byte, n int) []byte {
	return data[:n] // Holds reference to entire data!
}

//  GOOD: Independent copy
func getFirstNBytesSafe(data []byte, n int) []byte {
	result := make([]byte, n)
	copy(result, data[:n])
	return result
}

// ============================================
// CLEARING SLICE ELEMENTS FOR GC
// ============================================

func clearSliceForGC(s []*HeavyObject) {
	//  BAD: Just truncating doesn't help GC
	s = s[:0]

	//  GOOD: Nil out references first
	for i := range s {
		s[i] = nil // Allow GC to collect objects
	}
	s = s[:0]
}

type HeavyObject struct {
	data [1000]byte
}
