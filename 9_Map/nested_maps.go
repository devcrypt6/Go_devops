package main

import "fmt"

func main5() {
    // ============================================
    // NESTED MAPS: MAP OF MAPS
    // ============================================
    
    // User -> Property -> Value
    users := map[string]map[string]string{
        "alice": {
            "email": "alice@example.com",
            "city":  "Berlin",
            "role":  "admin",
        },
        "bob": {
            "email": "bob@example.com",
            "city":  "Munich",
            "role":  "user",
        },
    }
    
    fmt.Println("=== Nested Map ===")
    fmt.Println("Alice's email:", users["alice"]["email"])
    
    // ============================================
    // SAFE ACCESS TO NESTED MAPS
    // ============================================
    
    // UNSAFE: Can panic if outer key doesn't exist
    // email := users["charlie"]["email"]  // PANIC!
    
    // SAFE: Check each level
    if userInfo, ok := users["charlie"]; ok {
        if email, ok := userInfo["email"]; ok {
            fmt.Println("Charlie's email:", email)
        }
    } else {
        fmt.Println("Charlie not found")
    }
    
    // ============================================
    // ADDING TO NESTED MAP
    // ============================================
    
    // Must initialize inner map first!
    users["charlie"] = make(map[string]string)
    users["charlie"]["email"] = "charlie@example.com"
    users["charlie"]["city"] = "Hamburg"
    
    fmt.Println("\nAfter adding Charlie:", users["charlie"])
    
    // ============================================
    // HELPER FUNCTION FOR NESTED MAP
    // ============================================
    
    getNestedValue := func(users map[string]map[string]string, user, key string) (string, bool) {
        if userInfo, ok := users[user]; ok {
            if value, ok := userInfo[key]; ok {
                return value, true
            }
        }
        return "", false
    }
    
    if email, ok := getNestedValue(users, "alice", "email"); ok {
        fmt.Println("\nAlice's email (safe):", email)
    }
    
    // ============================================
    // GRAPH REPRESENTATION (Adjacency List)
    // ============================================
    
    fmt.Println("\n=== Graph Representation ===")
    
    // Graph: node -> list of neighbors
    graph := map[string][]string{
        "A": {"B", "C"},
        "B": {"A", "D", "E"},
        "C": {"A", "F"},
        "D": {"B"},
        "E": {"B", "F"},
        "F": {"C", "E"},
    }
    
    fmt.Println("Neighbors of B:", graph["B"])
    
    // ============================================
    // WEIGHTED GRAPH
    // ============================================
    
    type Edge struct {
        To     string
        Weight int
    }
    
    weightedGraph := map[string][]Edge{
        "A": {{To: "B", Weight: 5}, {To: "C", Weight: 3}},
        "B": {{To: "D", Weight: 2}},
        "C": {{To: "D", Weight: 7}},
    }
    
    fmt.Println("\nWeighted edges from A:")
    for _, edge := range weightedGraph["A"] {
        fmt.Printf("  -> %s (weight: %d)\n", edge.To, edge.Weight)
    }
    
    // ============================================
    // SPARSE MATRIX (2D Coordinates)
    // ============================================
    
    type Point struct {
        X, Y int
    }
    
    sparseMatrix := map[Point]int{
        {0, 0}: 1,
        {2, 3}: 5,
        {5, 5}: 9,
    }
    
    fmt.Println("\n=== Sparse Matrix ===")
    fmt.Println("Value at (2,3):", sparseMatrix[Point{2, 3}])
    fmt.Println("Value at (1,1):", sparseMatrix[Point{1, 1}])  // 0 (default)
    
    // ============================================
    // MULTI-KEY MAP (Using Struct as Key)
    // ============================================
    
    type UserSession struct {
        UserID    int
        SessionID string
    }
    
    sessions := map[UserSession]string{
        {UserID: 1, SessionID: "abc123"}: "active",
        {UserID: 1, SessionID: "def456"}: "expired",
        {UserID: 2, SessionID: "ghi789"}: "active",
    }
    
    fmt.Println("\n=== Multi-Key Map ===")
    key := UserSession{UserID: 1, SessionID: "abc123"}
    fmt.Println("Session status:", sessions[key])
}
