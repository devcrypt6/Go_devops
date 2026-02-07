package main

import "fmt"

func main() {
	// ============================================
	// MAPS: KEY-VALUE PAIRS (HASH TABLE)
	// ============================================

	// Maps are Go's built-in hash table
	// Also called: dictionary, associative array, hash map

	// ============================================
	// MAP DECLARATION - Method 1: var
	// ============================================

	var ages map[string]int // nil map (cannot add to it!)
	fmt.Println("Nil map:", ages)
	fmt.Println("Is nil:", ages == nil)
	// ages["Alice"] = 25  // PANIC! Can't assign to nil map

	// ============================================
	// MAP DECLARATION - Method 2: make()
	// ============================================

	ages = make(map[string]int) // Empty map (can add to it)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35

	fmt.Println("\nAges map:", ages)

	// ============================================
	// MAP DECLARATION - Method 3: Literal
	// ============================================

	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}

	fmt.Println("Scores:", scores)

	// ============================================
	// ACCESSING MAP VALUES
	// ============================================

	fmt.Println("\n=== Accessing Values ===")
	fmt.Println("Alice's age:", ages["Alice"]) // 25
	fmt.Println("Bob's score:", scores["Bob"]) // 87

	// Accessing non-existent key returns zero value
	fmt.Println("David's age:", ages["David"]) // 0 (not in map)

	// ============================================
	// CHECK IF KEY EXISTS (THE RIGHT WAY)
	// ============================================

	fmt.Println("\n=== Checking Existence ===")

	// Two-value assignment: value, exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Printf("Alice is %d years old\n", age)
	}

	age, exists = ages["David"]
	if exists {
		fmt.Printf("David is %d years old\n", age)
	} else {
		fmt.Println("David not found in map")
	}

	// Short form (common idiom)
	if score, ok := scores["Charlie"]; ok {
		fmt.Printf("Charlie scored %d\n", score)
	}

	// ============================================
	// MODIFYING MAP VALUES
	// ============================================

	fmt.Println("\n=== Modifying Values ===")
	ages["Alice"] = 26 // Update existing
	ages["David"] = 28 // Add new

	fmt.Println("Updated ages:", ages)

	// ============================================
	// DELETING FROM MAP
	// ============================================

	fmt.Println("\n=== Deleting Keys ===")
	delete(ages, "Bob") // Remove Bob
	fmt.Println("After deleting Bob:", ages)

	// Deleting non-existent key is safe (no error)
	delete(ages, "Eve") // Does nothing

	// ============================================
	// MAP LENGTH
	// ============================================

	fmt.Println("\nMap length:", len(ages))
}
