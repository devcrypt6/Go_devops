package main

import (
	"fmt"
)

// Calculate sum of slice
func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Calculate average
func average(numbers []int) float64 {
	// Return 0 if slice is empty
	if len(numbers) == 0 {
		return 0
	}
	return float64(sum(numbers)) / float64(len(numbers))
}

// Find minimum value
func min(numbers []int) (int, error) {
	// Return error if slice is empty
	if len(numbers) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	minVal := numbers[0]
	for _, num := range numbers {
		if num < minVal {
			minVal = num
		}
	}
	return minVal, nil
}

// Find maximum value
func max(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	maxVal := numbers[0]
	for _, num := range numbers {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal, nil
}

// Find median (middle value when sorted)
func median(numbers []int) float64 {
	// Hint: Sort first, then return middle value
	// If even length, return average of two middle values
	if len(numbers) == 0 {
		return 0
	}
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	// Simple bubble sort
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < sorted[i] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	mid := len(sorted) / 2
	if len(sorted)%2 == 1 {
		return float64(sorted[mid])
	}
	return float64(sorted[mid-1]+sorted[mid]) / 2
}

func main8() {
	data := []int{5, 2, 8, 1, 9, 3, 7}

	fmt.Println("Data:", data)
	fmt.Println("Sum:", sum(data))
	fmt.Printf("Average: %.2f\n", average(data))

	if minVal, err := min(data); err == nil {
		fmt.Println("Min:", minVal)
	}

	if maxVal, err := max(data); err == nil {
		fmt.Println("Max:", maxVal)
	}

	fmt.Printf("Median: %.2f\n", median(data))
}
