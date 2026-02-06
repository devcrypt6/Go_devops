package main

import "fmt"

func main() {
	// ============================================
	// HOW SLICES GROW: CAPACITY DOUBLING
	// ============================================

	fmt.Println("=== Slice Growth Pattern ===")
	s := make([]int, 0)

	for i := 0; i < 20; i++ {
		lenBefore := len(s)
		capBefore := cap(s)

		s = append(s, i)

		if cap(s) != capBefore {
			fmt.Printf("Grew from cap=%d to cap=%d at len=%d\n",
				capBefore, cap(s), lenBefore)
		}
	}

	// Output shows capacity growth:
	// 0 → 1 → 2 → 4 → 8 → 16 → 32...
	// Approximately doubles each time!

	// ============================================
	// PRE-ALLOCATING FOR PERFORMANCE
	// ============================================

	// BAD: Multiple reallocations
	slowSlice := []int{}
	for i := 0; i < 10000; i++ {
		slowSlice = append(slowSlice, i)
		// May reallocate ~14 times!
	}

	// GOOD: Pre-allocate known size
	fastSlice := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		fastSlice = append(fastSlice, i)
		// No reallocations!
	}

	// ============================================
	// BENCHMARKING APPEND STRATEGIES
	// ============================================

	demonstratePreallocation()
}

func demonstratePreallocation() {
	fmt.Println("\n=== Pre-allocation Benefits ===")

	// Without pre-allocation
	var s1 []int
	// initialCap := cap(s1)
	reallocations := 0

	for i := 0; i < 1000; i++ {
		prevCap := cap(s1)
		s1 = append(s1, i)
		if cap(s1) != prevCap {
			reallocations++
		}
	}
	fmt.Printf("Without pre-allocation: %d reallocations\n", reallocations)

	// With pre-allocation
	s2 := make([]int, 0, 1000)
	reallocations = 0

	for i := 0; i < 1000; i++ {
		prevCap := cap(s2)
		s2 = append(s2, i)
		if cap(s2) != prevCap {
			reallocations++
		}
	}
	fmt.Printf("With pre-allocation: %d reallocations\n", reallocations)
}

// ============================================
// PRODUCTION TIP: Pre-allocate When Possible
// ============================================

func processLargeDataset(data []string) []Result {
	//  GOOD: Pre-allocate result slice
	results := make([]Result, 0, len(data))

	for _, item := range data {
		result := process(item)
		results = append(results, result)
	}

	return results
}

type Result struct {
	Value string
}

func process(s string) Result {
	return Result{Value: s}
}
