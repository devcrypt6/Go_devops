package main

import (
	"fmt"
	"sort"
)

func main3() {
	// ============================================
	// ITERATING MAPS WITH RANGE
	// ============================================

	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
		"David":   28,
	}

	// Both key and value
	fmt.Println("=== All Entries ===")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Only keys
	fmt.Println("\n=== Keys Only ===")
	for name := range ages {
		fmt.Println(name)
	}

	// Only values (use _ to ignore key)
	fmt.Println("\n=== Values Only ===")
	for _, age := range ages {
		fmt.Println(age)
	}

	// ============================================
	// MAP ITERATION ORDER IS RANDOM!
	// ============================================

	fmt.Println("\n=== Random Order Demonstration ===")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d: ", i+1)
		for name := range ages {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	// Order may vary each time!

	// ============================================
	// SORTED MAP ITERATION
	// ============================================

	fmt.Println("\n=== Sorted by Key ===")

	// 1. Extract keys into slice
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	// 2. Sort keys
	sort.Strings(names)

	// 3. Iterate in sorted order
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, ages[name])
	}

	// ============================================
	// SORTED BY VALUE
	// ============================================

	fmt.Println("\n=== Sorted by Value (Age) ===")

	// Create slice of key-value pairs
	type kv struct {
		Key   string
		Value int
	}

	pairs := make([]kv, 0, len(ages))
	for k, v := range ages {
		pairs = append(pairs, kv{k, v})
	}

	// Sort by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair.Key, pair.Value)
	}

	// ============================================
	// COUNT OCCURRENCES PATTERN
	// ============================================

	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

	// Count frequency
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++ // Auto-initializes to 0, then increments
	}

	fmt.Println("\n=== Word Frequency ===")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
