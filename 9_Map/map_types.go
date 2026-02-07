package main

import "fmt"

func main2() {
    // ============================================
    // DIFFERENT KEY TYPES
    // ============================================
    
    // String keys (most common)
    names := map[string]string{
        "en": "Hello",
        "es": "Hola",
        "fr": "Bonjour",
    }
    fmt.Println("Greetings:", names)
    
    // Int keys
    fibonacci := map[int]int{
        0: 0,
        1: 1,
        2: 1,
        3: 2,
        4: 3,
        5: 5,
    }
    fmt.Println("Fibonacci:", fibonacci)
    
    // Bool keys (rare but valid)
    settings := map[bool]string{
        true:  "enabled",
        false: "disabled",
    }
    fmt.Println("Settings:", settings)
    
    // ============================================
    // STRUCT AS KEY (Must be comparable)
    // ============================================
    
    type Point struct {
        X, Y int
    }
    
    distances := map[Point]float64{
        {0, 0}: 0.0,
        {1, 1}: 1.414,
        {3, 4}: 5.0,
    }
    
    fmt.Println("\nDistances from origin:", distances)
    fmt.Println("Distance of (3,4):", distances[Point{3, 4}])
    
    // ============================================
    // DIFFERENT VALUE TYPES
    // ============================================
    
    // Map of slices
    groups := map[string][]string{
        "fruits":     {"apple", "banana", "orange"},
        "vegetables": {"carrot", "broccoli", "spinach"},
    }
    
    fmt.Println("\nGroups:")
    for category, items := range groups {
        fmt.Printf("%s: %v\n", category, items)
    }
    
    // Map of maps (nested)
    users := map[string]map[string]string{
        "alice": {
            "email": "alice@example.com",
            "city":  "Berlin",
        },
        "bob": {
            "email": "bob@example.com",
            "city":  "Munich",
        },
    }
    
    fmt.Println("\nUsers:")
    fmt.Println("Alice's email:", users["alice"]["email"])
    
    // Map of structs
    type Student struct {
        Name  string
        Grade int
    }
    
    students := map[int]Student{
        101: {"Alice", 95},
        102: {"Bob", 87},
        103: {"Charlie", 92},
    }
    
    fmt.Println("\nStudents:")
    fmt.Println("Student 101:", students[101])
    
    // ============================================
    // INVALID KEY TYPES
    // ============================================
    
    // COMPILE ERROR: Slices cannot be keys (not comparable)
    // invalidMap := map[[]int]string{}
    
    // COMPILE ERROR: Maps cannot be keys
    // invalidMap2 := map[map[string]int]string{}
    
    // COMPILE ERROR: Functions cannot be keys
    // invalidMap3 := map[func()]string{}
    
    // VALID: Arrays CAN be keys (they're comparable)
    arrayKeys := map[[3]int]string{
        {1, 2, 3}: "one-two-three",
        {4, 5, 6}: "four-five-six",
    }
    fmt.Println("\nArray keys:", arrayKeys)
}
