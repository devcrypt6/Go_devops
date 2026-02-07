package main

import "fmt"

func main4() {
    // ============================================
    // PATTERN 1: SET (Using map[T]bool)
    // ============================================
    
    // Set of unique strings
    uniqueWords := make(map[string]bool)
    words := []string{"apple", "banana", "apple", "cherry", "banana"}
    
    for _, word := range words {
        uniqueWords[word] = true
    }
    
    fmt.Println("=== Set Pattern ===")
    fmt.Println("Unique words:", len(uniqueWords))
    for word := range uniqueWords {
        fmt.Println("-", word)
    }
    
    // Check membership
    if uniqueWords["apple"] {
        fmt.Println("'apple' is in set")
    }
    
    // ============================================
    // PATTERN 2: SET (Using map[T]struct{})
    // ============================================
    
    // More memory-efficient (struct{} has zero size)
    visited := make(map[int]struct{})
    visited[1] = struct{}{}
    visited[2] = struct{}{}
    visited[3] = struct{}{}
    
    if _, exists := visited[2]; exists {
        fmt.Println("\n2 was visited")
    }
    
    // ============================================
    // PATTERN 3: GROUPING (map[K][]V)
    // ============================================
    
    fmt.Println("\n=== Grouping Pattern ===")
    
    type Student struct {
        Name  string
        Grade string
    }
    
    students := []Student{
        {"Alice", "A"},
        {"Bob", "B"},
        {"Charlie", "A"},
        {"David", "C"},
        {"Eve", "B"},
    }
    
    // Group by grade
    byGrade := make(map[string][]Student)
    for _, student := range students {
        byGrade[student.Grade] = append(byGrade[student.Grade], student)
    }
    
    for grade, studs := range byGrade {
        fmt.Printf("Grade %s: ", grade)
        for _, s := range studs {
            fmt.Printf("%s ", s.Name)
        }
        fmt.Println()
    }
    
    // ============================================
    // PATTERN 4: COUNTER/FREQUENCY
    // ============================================
    
    fmt.Println("\n=== Counter Pattern ===")
    
    text := "hello world"
    charCount := make(map[rune]int)
    
    for _, char := range text {
        charCount[char]++
    }
    
    fmt.Println("Character frequencies:")
    for char, count := range charCount {
        fmt.Printf("'%c': %d\n", char, count)
    }
    
    // ============================================
    // PATTERN 5: CACHE/MEMOIZATION
    // ============================================
    
    fmt.Println("\n=== Memoization Pattern ===")
    
    fibCache := make(map[int]int)
    
    var fibonacci func(int) int
    fibonacci = func(n int) int {
        // Check cache first
        if val, ok := fibCache[n]; ok {
            return val
        }
        
        // Base cases
        if n <= 1 {
            return n
        }
        
        // Calculate and cache
        result := fibonacci(n-1) + fibonacci(n-2)
        fibCache[n] = result
        return result
    }
    
    fmt.Println("Fibonacci(40):", fibonacci(40))
    fmt.Println("Cache size:", len(fibCache))
    
    // ============================================
    // PATTERN 6: DEFAULT VALUES
    // ============================================
    
    fmt.Println("\n=== Default Values Pattern ===")
    
    config := map[string]string{
        "host": "localhost",
        "port": "8080",
    }
    
    // Get with default
    host := config["host"]
    if host == "" {
        host = "0.0.0.0"  // Default
    }
    
    // Better: Helper function
    getOrDefault := func(m map[string]string, key, defaultVal string) string {
        if val, ok := m[key]; ok {
            return val
        }
        return defaultVal
    }
    
    host = getOrDefault(config, "host", "0.0.0.0")
    timeout := getOrDefault(config, "timeout", "30")
    
    fmt.Println("Host:", host)
    fmt.Println("Timeout:", timeout)
    
    // ============================================
    // PATTERN 7: INVERSE MAP
    // ============================================
    
    fmt.Println("\n=== Inverse Map Pattern ===")
    
    userIDToName := map[int]string{
        101: "Alice",
        102: "Bob",
        103: "Charlie",
    }
    
    // Create inverse mapping
    nameToUserID := make(map[string]int)
    for id, name := range userIDToName {
        nameToUserID[name] = id
    }
    
    fmt.Println("Alice's ID:", nameToUserID["Alice"])
    
    // ============================================
    // PATTERN 8: MERGE MAPS
    // ============================================
    
    fmt.Println("\n=== Merge Maps Pattern ===")
    
    defaults := map[string]int{
        "timeout":    30,
        "maxRetries": 3,
        "bufferSize": 1024,
    }
    
    userConfig := map[string]int{
        "timeout": 60,  // Override
    }
    
    // Merge (userConfig overrides defaults)
    merged := make(map[string]int)
    for k, v := range defaults {
        merged[k] = v
    }
    for k, v := range userConfig {
        merged[k] = v  // Overrides
    }
    
    fmt.Println("Merged config:", merged)
}
