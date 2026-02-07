package main

import (
    "fmt"
    "time"
)

func main7() {
    // ============================================
    // PRE-ALLOCATION FOR PERFORMANCE
    // ============================================
    
    fmt.Println("=== Pre-allocation Benefits ===")
    
    // Without pre-allocation
    start := time.Now()
    m1 := make(map[int]int)
    for i := 0; i < 100000; i++ {
        m1[i] = i * 2
    }
    elapsed1 := time.Since(start)
    
    // With pre-allocation
    start = time.Now()
    m2 := make(map[int]int, 100000)  // Pre-allocate capacity
    for i := 0; i < 100000; i++ {
        m2[i] = i * 2
    }
    elapsed2 := time.Since(start)
    
    fmt.Printf("Without pre-allocation: %v\n", elapsed1)
    fmt.Printf("With pre-allocation:    %v\n", elapsed2)
    fmt.Printf("Speedup: %.2fx\n", float64(elapsed1)/float64(elapsed2))
    
    // ============================================
    // MAP LOOKUP PERFORMANCE
    // ============================================
    
    fmt.Println("\n=== Map Lookup Performance ===")
    
    // Map lookups are O(1) average case
    largeMap := make(map[int]int, 1000000)
    for i := 0; i < 1000000; i++ {
        largeMap[i] = i
    }
    
    start = time.Now()
    for i := 0; i < 100000; i++ {
        _ = largeMap[i]
    }
    fmt.Printf("100k lookups in map of 1M: %v\n", time.Since(start))
    
    // ============================================
    // MEMORY CONSIDERATIONS
    // ============================================
    
    fmt.Println("\n=== Memory Considerations ===")
    
    // Maps don't shrink automatically
    bigMap := make(map[int]int)
    for i := 0; i < 1000000; i++ {
        bigMap[i] = i
    }
    
    fmt.Println("Map with 1M entries")
    
    // Delete all entries
    for k := range bigMap {
        delete(bigMap, k)
    }
    
    fmt.Println("After deleting all entries:")
    fmt.Println("Length:", len(bigMap))  // 0
    fmt.Println("But underlying memory is NOT freed!")
    
    // To reclaim memory, create new map
    bigMap = make(map[int]int)
    fmt.Println("New map created - memory can be GC'd")
    
    // ============================================
    // COMPARISON: MAP vs SLICE FOR LOOKUP
    // ============================================
    
    fmt.Println("\n=== Map vs Slice for Lookup ===")
    
    size := 10000
    
    // Build map
    m := make(map[int]bool, size)
    for i := 0; i < size; i++ {
        m[i] = true
    }
    
    // Build slice
    s := make([]int, size)
    for i := 0; i < size; i++ {
        s[i] = i
    }
    
    // Map lookup
    start = time.Now()
    for i := 0; i < 1000; i++ {
        _ = m[5000]  // O(1)
    }
    mapTime := time.Since(start)
    
    // Slice linear search
    start = time.Now()
    for i := 0; i < 1000; i++ {
        for _, v := range s {
            if v == 5000 {
                break
            }
        }
    }
    sliceTime := time.Since(start)
    
    fmt.Printf("Map lookup:   %v\n", mapTime)
    fmt.Printf("Slice search: %v\n", sliceTime)
    fmt.Printf("Map is %.1fx faster\n", float64(sliceTime)/float64(mapTime))
}

// ============================================
// HASH COLLISION (Conceptual)
// ============================================

// Go's map implementation handles collisions internally
// Uses chaining or open addressing
// Average case: O(1) lookup, insert, delete
// Worst case: O(n) if many collisions (very rare with good hash function)

// ============================================
// BEST PRACTICES
// ============================================

/*
1. PRE-ALLOCATE if size known:
   m := make(map[K]V, expectedSize)

2. USE APPROPRIATE KEY TYPE:
   - String keys: Good for config, lookup tables
   - Int keys: Good for indices, IDs
   - Struct keys: Good for multi-field lookups

3. AVOID LARGE STRUCTS AS VALUES:
   - Use pointers: map[K]*LargeStruct
   - Reduces copying

4. CHECK EXISTENCE:
   if val, ok := m[key]; ok { ... }

5. ITERATE WITH INTENT:
   - Delete during iteration: OK
   - Add during iteration: Unsafe

6. CONCURRENT ACCESS:
   - Always use mutexes or sync.Map
   - Never assume single-threaded

7. MEMORY MANAGEMENT:
   - If map grows large then shrinks, create new map
   - Maps don't auto-shrink

8. USE SETS EFFICIENTLY:
   - map[T]struct{} is more memory-efficient than map[T]bool
   - struct{} has zero size
*/
