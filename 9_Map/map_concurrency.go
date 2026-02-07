package main

import (
    "fmt"
    "sync"
)

func main6() {
    // ============================================
    // MAPS ARE NOT THREAD-SAFE!
    // ============================================
    
    // DANGER: Concurrent writes can cause panic or corruption
    demonstrateUnsafeMap()
    
    // SOLUTION 1: Use sync.Mutex
    demonstrateMutexMap()
    
    // SOLUTION 2: Use sync.RWMutex (better for many reads)
    demonstrateRWMutexMap()
    
    // SOLUTION 3: Use sync.Map (for specific use cases)
    demonstrateSyncMap()
}

func demonstrateUnsafeMap() {
    fmt.Println("=== UNSAFE Map (Don't do this!) ===")
    
    // This could panic or corrupt data!
    unsafeMap := make(map[int]int)
    
    // DON'T RUN THIS - just showing the problem
    _ = unsafeMap
    
    fmt.Println("Skipping unsafe demonstration (would cause race condition)")
}

func demonstrateMutexMap() {
    fmt.Println("\n=== Safe Map with Mutex ===")
    
    type SafeMap struct {
        mu sync.Mutex
        m  map[string]int
    }
    
    safeMap := SafeMap{
        m: make(map[string]int),
    }
    
    // Set value (thread-safe)
    set := func(key string, value int) {
        safeMap.mu.Lock()
        defer safeMap.mu.Unlock()
        safeMap.m[key] = value
    }
    
    // Get value (thread-safe)
    get := func(key string) (int, bool) {
        safeMap.mu.Lock()
        defer safeMap.mu.Unlock()
        val, ok := safeMap.m[key]
        return val, ok
    }
    
    var wg sync.WaitGroup
    
    // Write concurrently
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", n)
            set(key, n*10)
        }(i)
    }
    
    wg.Wait()
    
    // Read result
    if val, ok := get("key5"); ok {
        fmt.Println("key5 =", val)
    }
}

func demonstrateRWMutexMap() {
    fmt.Println("\n=== Safe Map with RWMutex ===")
    
    // Better for read-heavy workloads
    type SafeMap struct {
        mu sync.RWMutex
        m  map[string]int
    }
    
    safeMap := SafeMap{
        m: make(map[string]int),
    }
    
    // Set (write lock)
    set := func(key string, value int) {
        safeMap.mu.Lock()
        defer safeMap.mu.Unlock()
        safeMap.m[key] = value
    }
    
    // Get (read lock - multiple readers allowed)
    get := func(key string) (int, bool) {
        safeMap.mu.RLock()
        defer safeMap.mu.RUnlock()
        val, ok := safeMap.m[key]
        return val, ok
    }
    
    // Initialize
    set("counter", 0)
    
    var wg sync.WaitGroup
    
    // Many concurrent reads
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            val, _ := get("counter")
            _ = val
        }()
    }
    
    wg.Wait()
    fmt.Println("Completed 100 concurrent reads safely")
}

func demonstrateSyncMap() {
    fmt.Println("\n=== sync.Map ===")
    
    // sync.Map is built-in concurrent map
    // Good for: append-only maps, or maps with stable keys
    var sm sync.Map
    
    // Store
    sm.Store("alice", 25)
    sm.Store("bob", 30)
    
    // Load
    if age, ok := sm.Load("alice"); ok {
        fmt.Println("Alice's age:", age)
    }
    
    // LoadOrStore (atomic)
    actual, loaded := sm.LoadOrStore("charlie", 35)
    if loaded {
        fmt.Println("Charlie already existed:", actual)
    } else {
        fmt.Println("Charlie added with age:", actual)
    }
    
    // Delete
    sm.Delete("bob")
    
    // Range (iterate)
    sm.Range(func(key, value interface{}) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true  // continue iteration
    })
}

// ============================================
// PRODUCTION PATTERN: Safe Map Wrapper
// ============================================

type ConcurrentMap struct {
    mu sync.RWMutex
    m  map[string]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
    return &ConcurrentMap{
        m: make(map[string]interface{}),
    }
}

func (cm *ConcurrentMap) Set(key string, value interface{}) {
    cm.mu.Lock()
    defer cm.mu.Unlock()
    cm.m[key] = value
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
    cm.mu.RLock()
    defer cm.mu.RUnlock()
    val, ok := cm.m[key]
    return val, ok
}

func (cm *ConcurrentMap) Delete(key string) {
    cm.mu.Lock()
    defer cm.mu.Unlock()
    delete(cm.m, key)
}

func (cm *ConcurrentMap) Len() int {
    cm.mu.RLock()
    defer cm.mu.RUnlock()
    return len(cm.m)
}
